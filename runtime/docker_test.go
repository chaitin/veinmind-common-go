package runtime

import (
	"fmt"
	"testing"
)

func TestDockerClient_Pull(t *testing.T) {
	d, err := NewDockerClient()
	if err != nil {
		t.Error(err)
	}

	repo, err := d.Pull("ubuntu:latest")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(repo)
}
