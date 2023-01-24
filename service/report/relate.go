package report

import (
	"errors"

	api "github.com/chaitin/libveinmind/go"
	"github.com/chaitin/libveinmind/go/containerd"
	"github.com/chaitin/libveinmind/go/docker"
	"github.com/chaitin/libveinmind/go/iac"
	"github.com/chaitin/libveinmind/go/remote"
)

// RelatedObject return relate object for ReportEvent
// Include multi of cloud native objects, such as api.Image/api.Container/iac.IAC
func (r *ReportEvent) RelatedObject() (interface{}, error) {
	var (
		runtime api.Runtime
		err     error
	)
	// TODO(d_infinite): support runtime init arguments
	switch r.RuntimeType {
	case Docker:
		runtime, err = docker.New()
		if err != nil {
			return nil, err
		}
	case Containerd:
		runtime, err = containerd.New()
		if err != nil {
			return nil, err
		}
	case Remote:
		if r.RuntimeRoot == "" {
			return nil, errors.New("report: remote runtime root can't be set as empty")
		}
		runtime, err = remote.New(r.RuntimeRoot)
		if err != nil {
			return nil, err
		}
	// TODO(d_infinite): support other runtime
	default:
		return nil, errors.New("report: not support runtime type")
	}

	switch r.DetectType {
	case Image:
		return runtime.OpenImageByID(r.ID)
	case Container:
		return runtime.OpenContainerByID(r.ID)
	case IaC:
		t, _ := iac.DiscoverType(r.ID)
		return iac.IAC{
			Path: r.ID,
			Type: t,
		}, nil
	default:
		return nil, errors.New("report: not support detect type")
	}
}
