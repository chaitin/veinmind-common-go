package registry

import (
	"encoding/json"
	"github.com/chaitin/libveinmind/go/plugin/log"
	commonAuth "github.com/chaitin/veinmind-common-go/pkg/auth"
	"github.com/chaitin/veinmind-common-go/pkg/request"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"strings"
)

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

func (client *Client) GetRepoManifest(repo string, options ...remote.Option) (*schema2.Manifest, error) {
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

	manifest := &schema2.Manifest{}
	err = json.Unmarshal(desc.Manifest, manifest)
	if err != nil {
		return nil, err
	}

	return manifest, nil
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
