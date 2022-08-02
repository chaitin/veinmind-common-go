package runtime

import (
	"context"
	"github.com/chaitin/veinmind-common-go/pkg/auth"
)

type Client interface {
	Pull(ctx context.Context, repo string) (string, error)
	Remove(ctx context.Context, id string) error
	Auth(ctx context.Context, config auth.AuthConfig) error
}
