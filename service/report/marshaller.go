package report

import (
	"bytes"
	"encoding/json"

	api "github.com/chaitin/libveinmind/go"
	"github.com/chaitin/libveinmind/go/containerd"
	"github.com/chaitin/libveinmind/go/docker"
	"github.com/chaitin/libveinmind/go/iac"
	"github.com/chaitin/libveinmind/go/kubernetes"
	"github.com/chaitin/libveinmind/go/remote"
	"github.com/chaitin/libveinmind/go/tarball"
	"github.com/pkg/errors"
)

var (
	toLevel = map[Level]string{
		Low:      "Low",
		Medium:   "Medium",
		High:     "High",
		Critical: "Critical",
		None:     "None",
	}

	fromLevel = map[string]Level{
		"Low":      Low,
		"Medium":   Medium,
		"High":     High,
		"Critical": Critical,
		"None":     None,
	}

	toDetectType = map[DetectType]string{
		Image:     "Image",
		Container: "Container",
		IaC:       "IaC",
	}

	fromDetectType = map[string]DetectType{
		"Image":     Image,
		"Container": Container,
		"IaC":       IaC,
	}

	toEventType = map[EventType]string{
		Risk:     "Risk",
		Invasion: "Invasion",
		Info:     "Info",
	}

	fromEventType = map[string]EventType{
		"Risk":     Risk,
		"Invasion": Invasion,
		"Info":     Info,
	}

	toAlertType = map[AlertType]string{
		Vulnerability:   "Vulnerability",
		MaliciousFile:   "MaliciousFile",
		Backdoor:        "Backdoor",
		Sensitive:       "Sensitive",
		AbnormalHistory: "AbnormalHistory",
		Weakpass:        "Weakpass",
		Asset:           "Asset",
		Basic:           "Basic",
		IaCRisk:         "IaC",
	}

	fromAlertType = map[string]AlertType{
		"Vulnerability":   Vulnerability,
		"MaliciousFile":   MaliciousFile,
		"Backdoor":        Backdoor,
		"Sensitive":       Sensitive,
		"AbnormalHistory": AbnormalHistory,
		"Weakpass":        Weakpass,
		"Asset":           Asset,
		"Basic":           Basic,
		"IaC":             IaCRisk,
	}

	toWeakpassService = map[WeakpassService]string{
		SSH:    "SSH",
		Redis:  "Redis",
		Mysql:  "Mysql",
		Tomcat: "Tomcat",
		Env:    "Env",
	}

	fromWeakpassService = map[string]WeakpassService{
		"SSH":    SSH,
		"Redis":  Redis,
		"Mysql":  Mysql,
		"Tomcat": Tomcat,
		"Env":    Env,
	}

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

func (a AlertType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toAlertType[a])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (a *AlertType) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	switch i.(type) {
	case uint32:
		*a = i.(AlertType)
	case float64:
		*a = AlertType(i.(float64))
	case string:
		*a = fromAlertType[i.(string)]
	}

	return nil
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

func (t *GeneralDetail) MarshalJSON() ([]byte, error) {
	return *t, nil
}

func (t *GeneralDetail) UnmarshalJSON(b []byte) error {
	*t = b
	return nil
}

func (o *Object) MarshalJSON() ([]byte, error) {
	switch v := o.Raw.(type) {
	case api.Image:
		o.Type = Image

		switch cast := v.(type) {
		case *docker.Image:
			o.RuntimeType = Docker
			o.ID = cast.ID()
		case *containerd.Image:
			o.RuntimeType = Containerd
			o.ID = cast.ID()
		case *remote.Image:
			o.RuntimeType = Remote
			o.RuntimeRoot = cast.Runtime().Root()
			o.ID = cast.ID()
		case *tarball.Image:
			o.RuntimeType = Tarball
			o.RuntimeRoot = cast.Runtime().Root()
			o.ID = cast.ID()
		}
	case api.Container:
		o.Type = Container

		switch cast := v.(type) {
		case *docker.Container:
			o.RuntimeType = Docker
			o.ID = cast.ID()
		case *containerd.Container:
			o.RuntimeType = Containerd
			o.ID = cast.ID()
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
