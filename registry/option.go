package registry

import (
	"errors"
	"io"

	commonAuth "github.com/chaitin/veinmind-common-go/pkg/auth"
	_ "github.com/chaitin/veinmind-common-go/runtime"
)

type Option func(client *Client) (*Client, error)

// WithAuth parse auth struct
func WithAuth(auth *commonAuth.Auth) Option {
	return func(client *Client) (*Client, error) {
		if auth.Registry != "" {
			if client.auth == nil {
				client.auth = make(map[string]commonAuth.Auth)
			}

			client.auth[auth.Registry] = *auth
		}

		return client, nil
	}
}

// WithAuthFromPath parse auth path to config entity
func WithAuthFromPath(path string) Option {
	return func(c *Client) (*Client, error) {
		if path == "" {
			return nil, errors.New("auth config path can't be empty")
		}

		authConfig, err := commonAuth.ParseAuthConfig(path)
		if err != nil {
			return nil, err
		}

		if c.auth == nil {
			c.auth = make(map[string]commonAuth.Auth)
		}
		for _, a := range authConfig.Auths {
			c.auth[a.Registry] = a
		}

		return c, nil
	}
}

func WithAuthFromReader(reader io.Reader) Option {
	return func(c *Client) (*Client, error) {
		authConfig, err := commonAuth.ParseAuthConfigFromReader(reader)
		if err != nil {
			return nil, err
		}

		if c.auth == nil {
			c.auth = make(map[string]commonAuth.Auth)
		}
		for _, a := range authConfig.Auths {
			c.auth[a.Registry] = a
		}

		return c, nil
	}
}

type Options struct {
	insecure bool
}

func DefaultOptions() Options {
	return Options{
		insecure: false,
	}
}

func (o Options) WithInsecure() Options {
	o.insecure = true
	return o
}
