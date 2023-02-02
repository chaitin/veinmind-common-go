package event

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"

	api "github.com/chaitin/libveinmind/go"
	"github.com/chaitin/libveinmind/go/containerd"
	"github.com/chaitin/libveinmind/go/docker"
	"github.com/chaitin/libveinmind/go/iac"
	"github.com/chaitin/libveinmind/go/kubernetes"
	"github.com/chaitin/libveinmind/go/remote"
	"github.com/chaitin/libveinmind/go/tarball"
)

func (a AlertType) String() string {
	return (string)(a)
}

// MarshalJSON && UnmarshalJSON Level
func (l Level) String() string {
	if v, ok := toLevel[l]; ok {
		return v
	}
	return "unknown"
}

func (l Level) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toLevel[l])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (l *Level) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	switch v := i.(type) {
	case uint32:
		*l = Level(v)
	case float64:
		*l = Level(uint32(v))
	case string:
		*l = fromLevel[v]
	}

	return nil
}

func (l Level) Color() string {
	if f, ok := colorFns[l]; ok {
		return f(strings.ToUpper(l.String()))
	}
	return strings.ToUpper(l.String())
}

// MarshalJSON && UnmarshalJSON EventType
func (e EventType) String() string {
	if v, ok := toEventType[e]; ok {
		return v
	}
	return "unknown"
}

func (e EventType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toEventType[e])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (e *EventType) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	switch i.(type) {
	case uint32:
		*e = i.(EventType)
	case float64:
		*e = EventType(i.(float64))
	case string:
		*e = fromEventType[i.(string)]
	}

	return nil
}

// MarshalJSON && UnmarshalJSON DetectType
func (d DetectType) String() string {
	if v, ok := toDetectType[d]; ok {
		return v
	}
	return "unknown"
}

func (d DetectType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toDetectType[d])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (d *DetectType) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	switch i.(type) {
	case uint32:
		*d = i.(DetectType)
	case float64:
		*d = DetectType(i.(float64))
	case string:
		*d = fromDetectType[i.(string)]
	}

	return nil
}

// MarshalJSON && UnmarshalJSON RuntimeType
func (t RuntimeType) String() string {
	if v, ok := toContainerRuntimeType[t]; ok {
		return v
	}
	return "unknown"
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

// MarshalJSON && UnmarshalJSON WeakpassService
func (w WeakpassService) String() string {
	if v, ok := toWeakpassService[w]; ok {
		return v
	}
	return "unknown"
}

func (w WeakpassService) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toWeakpassService[w])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (w *WeakpassService) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	switch i.(type) {
	case uint32:
		*w = i.(WeakpassService)
	case float64:
		*w = WeakpassService(i.(float64))
	case string:
		*w = fromWeakpassService[i.(string)]
	}

	return nil
}

// MarshalJSON && UnmarshalJSON Object
func (o *Object) MarshalJSON() ([]byte, error) {
	switch v := o.Raw.(type) {
	case api.Image:
		o.Type = Image

		switch cast := v.(type) {
		case *docker.Image:
			o.RuntimeType = Docker
		case *containerd.Image:
			o.RuntimeType = Containerd
		case *remote.Image:
			o.RuntimeType = Remote
			o.RuntimeRoot = cast.Runtime().Root()
		case *tarball.Image:
			o.RuntimeType = Tarball
			o.RuntimeRoot = cast.Runtime().Root()
		}
	case api.Container:
		o.Type = Container
		switch v.(type) {
		case *docker.Container:
			o.RuntimeType = Docker
		case *containerd.Container:
			o.RuntimeType = Containerd
		}
	case api.Cluster:
		o.Type = Cluster

		switch cast := v.(type) {
		case *kubernetes.Kubernetes:
			o.ClusterType = ClusterKubernetes
			o.ClusterConfigPath = cast.ConfigPath()
			o.ClusterConfigByte = cast.ConfigBytes()
		}
	case iac.IAC:
		o.Type = IaC
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
	case Docker:
		runtime, err = docker.New()
		if err != nil {
			return err
		}
	case Containerd:
		runtime, err = containerd.New()
		if err != nil {
			return err
		}
	case Remote:
		if o.RuntimeRoot == "" {
			return errors.New("report: remote runtime root can't be set as empty")
		}
		runtime, err = remote.New(o.RuntimeRoot)
		if err != nil {
			return err
		}
	case Tarball:
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
	case Image:
		i, err := runtime.OpenImageByID(o.ID)
		if err != nil {
			return err
		}
		o.Raw = i
	case Container:
		c, err := runtime.OpenContainerByID(o.ID)
		if err != nil {
			return err
		}
		o.Raw = c
	case IaC:
		t, _ := iac.DiscoverType(o.ID)
		o.Raw = iac.IAC{
			Path: o.ID,
			Type: t,
		}
	case Cluster:
		switch o.ClusterType {
		case ClusterKubernetes:
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

// UnmarshalJSON Event
func (e *Event) UnmarshalJSON(b []byte) error {
	var basic = &BasicInfo{}
	err := json.Unmarshal(b, &basic)
	if err != nil {
		return err
	}
	objFunc := alertDetailMap[basic.AlertType]
	details := objFunc()
	err = json.Unmarshal(b, &details)
	if err != nil {
		return err
	}
	e.BasicInfo = basic
	e.DetailInfo = details
	return nil
}
