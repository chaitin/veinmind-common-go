package runtime

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"github.com/chaitin/libveinmind/go/plugin/log"
	commonAuth "github.com/chaitin/veinmind-common-go/pkg/auth"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/distribution/reference"
	dockertypes "github.com/docker/docker/api/types"
	dockercli "github.com/docker/docker/client"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const dockerConfigPath = "/root/.docker/config.json"

type DockerClient struct {
	ctx     context.Context
	auth    map[string]commonAuth.Auth
	options []remote.Option
}

func parseDockerAuthConfig(path string) (map[string]commonAuth.Auth, error) {
	dockerConfig := configfile.ConfigFile{}
	authConfigMap := make(map[string]commonAuth.Auth)

	if _, err := os.Stat(path); err != nil {
		dockerConfigByte, err := ioutil.ReadFile(path)

		err = json.Unmarshal(dockerConfigByte, &dockerConfig)
		if err != nil {
			return nil, err
		} else {
			for server, config := range dockerConfig.AuthConfigs {
				u, err := url.Parse(server)
				registryName := ""
				if err != nil {
					registryName = server
				} else {
					registryName = u.Host
				}

				registry, err := name.NewRegistry(registryName)
				if err != nil {
					log.Error(err)
					continue
				}

				if config.Auth != "" {
					authDecode, err := base64.StdEncoding.DecodeString(config.Auth)
					if err == nil {
						authSplit := strings.Split(string(authDecode), ":")
						if len(authSplit) == 2 {
							auth := commonAuth.Auth{
								Username: authSplit[0],
								Password: authSplit[1],
							}
							authConfigMap[registry.String()] = auth
						} else {
							log.Error("docker config auth block length wrong")
							continue
						}
					} else {
						log.Error(err)
						continue
					}
				}
			}
			return authConfigMap, nil
		}
	} else {
		return nil, err
	}
}

func NewDockerClient(opts ...Option) (Client, error) {
	c := &DockerClient{}
	c.ctx = context.Background()
	c.auth = make(map[string]commonAuth.Auth)

	// Get Auth Token From Config File
	auth, err := parseDockerAuthConfig(dockerConfigPath)
	if err != nil {
		log.Error(err)
	} else {
		c.auth = auth
	}

	// Double check
	if c.auth == nil {
		c.auth = make(map[string]commonAuth.Auth)
	}

	// Options handle
	for _, opt := range opts {
		cNew, err := opt(c)
		if err != nil {
			log.Error(err)
			continue
		}
		c = cNew.(*DockerClient)
	}

	var clientOpts []remote.Option
	clientOpts = append(clientOpts, remote.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}))
	c.options = clientOpts

	return c, nil
}

func (client *DockerClient) Auth(config commonAuth.AuthConfig) error {
	for _, auth := range config.Auths {
		client.auth[auth.Registry] = auth
	}

	return nil
}

func (client *DockerClient) Pull(repo string) (string, error) {
	c, err := dockercli.NewClientWithOpts(dockercli.FromEnv, dockercli.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

	named, err := reference.ParseDockerRef(repo)
	if err != nil {
		return "", err
	}

	domain := reference.Domain(named)
	var auth commonAuth.Auth
	if v, ok := client.auth[domain]; ok {
		auth = v
	}

	// Generate Auth Token
	token, err := command.EncodeAuthToBase64(dockertypes.AuthConfig{
		Username: auth.Username,
		Password: auth.Password})

	var closer io.ReadCloser
	if token == "" {
		closer, err = c.ImagePull(client.ctx, repo, dockertypes.ImagePullOptions{})
		if err != nil {
			return "", err
		}
	} else {
		closer, err = c.ImagePull(client.ctx, repo, dockertypes.ImagePullOptions{
			RegistryAuth: token,
		})
		if err != nil {
			return "", err
		}
	}

	_, err = ioutil.ReadAll(closer)
	if err != nil {
		return "", err
	}

	return named.String(), nil
}

func (client *DockerClient) Remove(id string) error {
	c, err := dockercli.NewClientWithOpts(dockercli.FromEnv, dockercli.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	_, err = c.ImageRemove(client.ctx, id, dockertypes.ImageRemoveOptions{
		Force:         true,
		PruneChildren: true,
	})
	return err
}
