package types

import (
	"bytes"
	"encoding/json"
)

type RuntimeType uint32

const (
	Docker RuntimeType = iota
	Containerd
	Remote
	Tarball
	Kubernetes
)

var (
	toContainerRuntimeType = map[RuntimeType]string{
		Docker:     "docker",
		Containerd: "containerd",
		Remote:     "remote",
		Tarball:    "tarball",
		Kubernetes: "kubernetes",
	}

	fromContainerRuntimeType = map[string]RuntimeType{
		"docker":     Docker,
		"containerd": Containerd,
		"remote":     Remote,
		"tarball":    Tarball,
		"kubernetes": Kubernetes,
	}
)

func (t RuntimeType) String() string {
	switch t {
	case Docker:
		return "docker"
	case Containerd:
		return "containerd"
	case Remote:
		return "remote"
	case Tarball:
		return "tarball"
	case Kubernetes:
		return "kubernetes"
	default:
		return "unknown"
	}
}

func (t RuntimeType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toContainerRuntimeType[t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *RuntimeType) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	switch i.(type) {
	case uint32:
		*t = i.(RuntimeType)
	case float64:
		*t = RuntimeType(i.(float64))
	case string:
		*t = fromContainerRuntimeType[i.(string)]
	}

	return nil
}
