package runtime

import (
	"context"
	"github.com/chaitin/veinmind-common-go/pkg/auth"
	"github.com/pkg/errors"
)

type Option func(c Client) (Client, error)

// WithAuthFromPath parse auth path to config entity
func WithAuthFromPath(path string) Option {
	return func(c Client) (Client, error) {
		if path == "" {
			return nil, errors.New("auth config path can't be empty")
		}

		authConfig, err := auth.ParseAuthConfig(path)
		if err != nil {
			return nil, err
		}

		err = c.Auth(context.Background(), *authConfig)
		if err != nil {
			return nil, err
		}

		return c, nil
	}
}

// WithAuth init client authInfo with an entity directly
func WithAuth(authConfig auth.AuthConfig) Option {
	return func(c Client) (Client, error) {
		err := c.Auth(context.Background(), authConfig)
		if err != nil {
			return nil, err
		}

		return c, nil
	}
}

// WithHost set host address for runtime client
func WithHost(host string) Option {
	return func(c Client) (Client, error) {
		if dockerClient, ok := c.(*DockerClient); ok {
			dockerClient.host = host
			return dockerClient, nil
		} else {
			return nil, errors.New("runtime: can't apply docker options to other runtime")
		}
	}
}
