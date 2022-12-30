package registry

import (
	"crypto/tls"
	"encoding/json"
	"errors"

	"net"
	"net/http"
	"strings"
	"time"

	"github.com/chaitin/libveinmind/go/plugin/log"
	commonAuth "github.com/chaitin/veinmind-common-go/pkg/auth"
	"github.com/chaitin/veinmind-common-go/pkg/request"
	"github.com/docker/distribution/manifest/manifestlist"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

func init() {
	remote.DefaultTransport = &http.Transport{
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
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
}

type Client struct {
	auth map[string]commonAuth.Auth
}

func NewClient(opts ...Option) (*Client, error) {
	c := new(Client)

	for _, opt := range opts {
		cTemp, err := opt(c)
		if err != nil {
			log.Error(err)
			continue
		}
		c = cTemp
	}

	return c, nil
}

func (client *Client) GetRepos(registry string, opts ...Option) ([]string, error) {
	var (
		repos     []string
		optionsV2 []remote.Option
		optionsV1 []request.Option
	)

	// init options
	optionsV1 = make([]request.Option, 0)
	optionsV2 = make([]remote.Option, 0)

	for _, opt := range opts {
		cTemp, err := opt(client)
		if err != nil {
			log.Error(err)
			continue
		}
		*client = *cTemp
	}

	// fetch auth
	authOption, err := client.fetchAuthOption(registry)
	if err == nil && authOption != nil {
		optionsV2 = append(optionsV2, authOption)
	}

	authT, err := client.fetchAuth(registry)
	if err == nil && authT != nil {
		optionsV1 = append(optionsV1, request.WithAuth(&authn.Basic{
			Username: authT.Username,
			Password: authT.Password,
		}))
	}

	r, err := name.NewRegistry(registry)
	if err != nil {
		return nil, err
	}

	// select version to auth (some registry catalog api doesn't support v2 auth, e.g. harbor)
	var (
		v1 bool
		v2 bool
	)

	_, err = remote.CatalogPage(r, "", 1, optionsV2...)
	if err != nil {
		v2 = false
	} else {
		v2 = true
	}

	_, err = request.CatalogPageV1(r, "", 1, optionsV1...)
	if err != nil {
		v1 = false
	} else {
		v1 = true
	}

	last := ""
	for {
		reposTemp := []string{}
		switch {
		case v2:
			reposTemp, err = remote.CatalogPage(r, last, 10000, optionsV2...)
			if err != nil {
				log.Error(err)
				continue
			}
			break
		case v1:
			reposTemp, err = request.CatalogPageV1(r, last, 10000, optionsV1...)
			if err != nil {
				log.Error(err)
				continue
			}
			break
		}

		if len(reposTemp) > 0 {
			repos = append(repos, reposTemp...)
		} else {
			break
		}

		last = reposTemp[len(reposTemp)-1]
	}

	// handle registry address
	for i, repo := range repos {
		if strings.HasPrefix(repo, r.RegistryStr()) {
			continue
		}
		repoT := strings.Join([]string{r.RegistryStr(), repo}, "/")
		repos[i] = repoT
	}

	return repos, err
}

func (client *Client) GetRepoTags(repo string, options ...remote.Option) ([]string, error) {
	repoR, err := name.NewRepository(repo)
	if err != nil {
		return nil, err
	}

	authOption, err := client.fetchAuthOption(repo)
	if err == nil {
		options = append(options, authOption)
	}

	return remote.List(repoR, options...)
}

func (client *Client) GetRepoManifests(repo string, options ...remote.Option) ([]*schema2.Manifest, error) {
	ref, err := name.ParseReference(repo)
	if err != nil {
		return nil, err
	}

	authOption, err := client.fetchAuthOption(repo)
	if err == nil {
		options = append(options, authOption)
	}

	desc, err := remote.Get(ref, options...)
	if err != nil {
		return nil, err
	}

	var (
		descData  map[string]interface{}
		manifests []*schema2.Manifest
	)
	err = json.Unmarshal(desc.Manifest, &descData)
	if err != nil {
		return nil, err
	}

	if mediaTypeInterface, ok := descData["mediaType"]; ok {
		if mediaType, ok := mediaTypeInterface.(string); ok {
			switch types.MediaType(mediaType) {
			case types.DockerManifestList:
				manifestList := &manifestlist.ManifestList{}
				err = json.Unmarshal(desc.Manifest, manifestList)
				if err != nil {
					return nil, err
				}

				for _, manifestFromList := range manifestList.Manifests {
					manifest := &schema2.Manifest{}
					manifest.Versioned = manifestList.Versioned
					manifest.Config = manifestFromList.Descriptor
					manifests = append(manifests, manifest)
				}
			case types.DockerManifestSchema2:
				manifest := &schema2.Manifest{}
				err = json.Unmarshal(desc.Manifest, manifest)
				if err != nil {
					return nil, err
				}
				manifests = append(manifests, manifest)
			}
		}
	}

	if len(manifests) == 0 {
		return nil, errors.New("registry: can't fetch manifest from repo: " + repo)
	}

	return manifests, nil
}
func (client *Client) Login(registry string, username string, password string) error {
	if client.auth == nil {
		client.auth = make(map[string]commonAuth.Auth)
	}

	client.auth[registry] = commonAuth.Auth{
		Registry: registry,
		Username: username,
		Password: password,
	}

	return nil
}
