package registry

import (
	"encoding/json"
	"github.com/distribution/distribution/manifest/schema2"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type Client struct {
}

func NewClient() (*Client, error) {
	c := new(Client)
	return c, nil
}

func (client *Client) GetRepoTags(repo string, options ...remote.Option) ([]string, error) {
	repoR, err := name.NewRepository(repo)
	if err != nil {
		return nil, err
	}
	return remote.List(repoR, options...)
}

func (client *Client) GetRepoManifest(repo string, options ...remote.Option) (*schema2.Manifest, error) {
	ref, err := name.ParseReference(repo)
	if err != nil {
		return nil, err
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
