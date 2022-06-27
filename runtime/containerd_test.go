package runtime

import (
	"fmt"
	"testing"
)

func TestContainerdClient_Pull(t *testing.T) {
	c, err := NewContainerdClient()
	if err != nil {
		t.Error(err)
	}

	repo, err := c.Pull("ubuntu:latest")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(repo)
}
