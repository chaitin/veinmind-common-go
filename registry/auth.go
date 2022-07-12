package registry

import (
	"errors"
	commonAuth "github.com/chaitin/veinmind-common-go/pkg/auth"
	"github.com/docker/distribution/reference"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func (c *Client) fetchAuthOption(repo string) (remote.Option, error) {
	var auth *commonAuth.Auth
	auth, err := c.fetchAuth(repo)
	if err != nil {
		return nil, err
	}

	return remote.WithAuth(&authn.Basic{
		Username: auth.Username,
		Password: auth.Password,
	}), nil
}

func (c *Client) fetchAuth(repo string) (*commonAuth.Auth, error) {
	ref, err := reference.Parse(repo)
	if err != nil {
		return nil, err
	}

	var domain string
	if named, ok := ref.(reference.Named); ok {
		domain = reference.Domain(named)

		if domain == "" {
			domain = named.Name()
		}
	}

	if v, ok := c.auth[domain]; ok {
		return &v, nil
	} else {
		return nil, errors.New("auth: registry domain not match")
	}
}
