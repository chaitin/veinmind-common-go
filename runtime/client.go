package runtime

import "github.com/chaitin/veinmind-common-go/pkg/auth"

type Client interface {
	Pull(repo string) (string, error)
	Remove(id string) error
	Auth(config auth.AuthConfig) error
}
