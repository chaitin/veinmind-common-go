package runtime

import (
	"context"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/images"
	"github.com/containerd/containerd/namespaces"
	"github.com/docker/distribution/reference"
	"strings"
)

const (
	ns             = "veinmind-agent"
	containerdSock = "/run/containerd/containerd.sock"
)

type ContainerdClient struct {
	client *containerd.Client
}

func NewContainerdClient() (Client, error) {
	c := &ContainerdClient{}
	client, err := containerd.New(containerdSock, containerd.WithDefaultNamespace(ns))
	if err != nil {
		return nil, err
	}

	c.client = client

	return c, nil
}

func (c *ContainerdClient) Auth(config AuthConfig) error {
	return nil
}

func (c *ContainerdClient) Pull(repo string) (string, error) {
	if named, err := reference.ParseDockerRef(repo); err == nil {
		repo = named.String()
	}

	image, err := c.client.Pull(context.Background(), repo, containerd.WithPullUnpack)
	if err != nil {
		return "", err
	}

	imageID := strings.Join([]string{ns, string(image.Target().Digest)}, "/")
	return imageID, nil
}

func (c *ContainerdClient) Remove(repo string) error {
	if named, err := reference.ParseDockerRef(repo); err == nil {
		repo = named.String()
	}

	var (
		ctx        = namespaces.WithNamespace(context.Background(), ns)
		imageStore = c.client.ImageService()
	)

	var opts []images.DeleteOpt
	opts = append(opts, images.SynchronousDelete())
	if err := imageStore.Delete(ctx, repo, opts...); err != nil {
		return err
	}

	return nil
}
