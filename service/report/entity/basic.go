package entity

import "github.com/chaitin/veinmind-common-go/service/report/types"

type ImageBasicDetail struct {
	References  []string `json:"references"`
	CreatedTime int64    `json:"created_time"`
	Env         []string `json:"env"`
	Entrypoint  []string `json:"entrypoint"`
	Cmd         []string `json:"cmd"`
	WorkingDir  string   `json:"working_dir"`
	Author      string   `json:"author"`
}

type ContainerBasicDetail struct {
	Name            string            `json:"name"`
	CreatedTime     int64             `json:"created_time"`
	State           string            `json:"state"`
	Runtime         types.RuntimeType `json:"runtime"`
	RuntimeUniqDesc string            `json:"runtime_uniq_desc,omitempty"`
	Hostname        string            `json:"hostname"`
	ImageID         string            `json:"imageID"`
	Privileged      bool              `json:"privileged,omitempty"`
	RootProcess     RootProcessDetail `json:"process"`
	Mounts          []MountDetail     `json:"mounts,omitempty"`
	Processes       []ProcessDetail   `json:"processes,omitempty"`
}
