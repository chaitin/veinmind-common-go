package entity

import (
	"encoding/json"
	"errors"

	api "github.com/chaitin/libveinmind/go"
	"github.com/chaitin/libveinmind/go/containerd"
	"github.com/chaitin/libveinmind/go/docker"
	"github.com/chaitin/libveinmind/go/iac"
	"github.com/chaitin/libveinmind/go/kubernetes"
	"github.com/chaitin/libveinmind/go/remote"
	"github.com/chaitin/libveinmind/go/tarball"
	"github.com/chaitin/veinmind-common-go/service/report/types"
)

// Object represents various cloud-native objects
type Object struct {
	Raw interface{} `json:"-"`

	// container runtime
	ID          string            `json:"id,omitempty"`
	Type        types.DetectType  `json:"type,omitempty"`
	RuntimeType types.RuntimeType `json:"runtime_type,omitempty"`
	RuntimeRoot string            `json:"runtime_root,omitempty"`

	// cluster config
	ClusterType       types.ClusterType `json:"cluster_type"`
	ClusterConfigPath string            `json:"cluster_config_path"`
	ClusterConfigByte []byte            `json:"cluster_config_byte"`
}

func (o *Object) MarshalJSON() ([]byte, error) {
	switch v := o.Raw.(type) {
	case api.Image:
		o.Type = types.Image

		switch cast := v.(type) {
		case *docker.Image:
			o.RuntimeType = types.Docker
			o.ID = cast.ID()
		case *containerd.Image:
			o.RuntimeType = types.Containerd
			o.ID = cast.ID()
		case *remote.Image:
			o.RuntimeType = types.Remote
			o.RuntimeRoot = cast.Runtime().Root()
			o.ID = cast.ID()
		case *tarball.Image:
			o.RuntimeType = types.Tarball
			o.RuntimeRoot = cast.Runtime().Root()
			o.ID = cast.ID()
		}
	case api.Container:
		o.Type = types.Container

		switch cast := v.(type) {
		case *docker.Container:
			o.RuntimeType = types.Docker
			o.ID = cast.ID()
		case *containerd.Container:
			o.RuntimeType = types.Containerd
			o.ID = cast.ID()
		}
	case api.Cluster:
		o.Type = types.Cluster

		switch cast := v.(type) {
		case *kubernetes.Kubernetes:
			o.ClusterType = types.ClusterKubernetes
			o.ClusterConfigPath = cast.ConfigPath()
			o.ClusterConfigByte = cast.ConfigBytes()
		}
	case iac.IAC:
		o.Type = types.IaC
		o.ID = v.Path
	}

	type Alias Object
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(o),
	})
}

func (o *Object) UnmarshalJSON(data []byte) error {
	type Alias Object
	alias := &struct {
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	err := json.Unmarshal(data, alias)
	if err != nil {
		return err
	}

	var (
		runtime api.Runtime
	)
	// TODO(d_infinite): support runtime init arguments
	switch o.RuntimeType {
	case types.Docker:
		runtime, err = docker.New()
		if err != nil {
			return err
		}
	case types.Containerd:
		runtime, err = containerd.New()
		if err != nil {
			return err
		}
	case types.Remote:
		if o.RuntimeRoot == "" {
			return errors.New("report: remote runtime root can't be set as empty")
		}
		runtime, err = remote.New(o.RuntimeRoot)
		if err != nil {
			return err
		}
	case types.Tarball:
		if o.RuntimeRoot == "" {
			return errors.New("report: remote runtime root can't be set as empty")
		}
		runtime, err = tarball.New(tarball.WithRoot(o.RuntimeRoot))
		if err != nil {
			return err
		}
	// TODO(d_infinite): support other runtime
	default:
		return errors.New("report: not support runtime type")
	}

	switch o.Type {
	case types.Image:
		i, err := runtime.OpenImageByID(o.ID)
		if err != nil {
			return err
		}
		o.Raw = i
	case types.Container:
		c, err := runtime.OpenContainerByID(o.ID)
		if err != nil {
			return err
		}
		o.Raw = c
	case types.IaC:
		t, _ := iac.DiscoverType(o.ID)
		o.Raw = iac.IAC{
			Path: o.ID,
			Type: t,
		}
	case types.Cluster:
		switch o.ClusterType {
		case types.ClusterKubernetes:
			k, err := kubernetes.New(kubernetes.WithKubeConfigPath(o.ClusterConfigPath), kubernetes.WithKubeConfigBytes(o.ClusterConfigByte))
			if err != nil {
				return err
			}
			o.Raw = k
		}
	default:
		return errors.New("report: not support detect type")
	}

	return nil
}
