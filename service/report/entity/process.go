package entity

type ProcessDetail struct {
	Cmdline    string   `json:"cmdline,omitempty"`
	Cwd        string   `json:"cwd,omitempty"`
	Environ    []string `json:"environ,omitempty"`
	Exe        string   `json:"exe,omitempty"`
	Gids       []int32  `json:"gids,omitempty"`
	Groupnames []string `json:"groupnames,omitempty"`
	Uids       []int32  `json:"uids,omitempty"`
	Usernames  []string `json:"usernames,omitempty"`
	Ppid       int32    `json:"ppid,omitempty"`
	Pid        int32    `json:"pid,omitempty"`
	HostPid    int32    `json:"host_pid,omitempty"`
	Name       string   `json:"name,omitempty"`
	Status     string   `json:"status,omitempty"`
	CreateTime int64    `json:"createTime,omitempty"`
}

type RootProcessDetail struct {
	Terminal     bool               `json:"terminal,omitempty"`
	UID          uint32             `json:"uid" platform:"linux,solaris"`
	Username     string             `json:"username" platform:"linux,solaris"`
	GID          uint32             `json:"gid" platform:"linux,solaris"`
	Groupname    string             `json:"groupname" platform:"linux,solaris"`
	Args         []string           `json:"args"`
	Env          []string           `json:"env,omitempty"`
	Cwd          string             `json:"cwd"`
	Capabilities CapabilitiesDetail `json:"capabilities,omitempty" platform:"linux"`
}

type CapabilitiesDetail struct {
	Bounding    []string `json:"bounding,omitempty" platform:"linux"`
	Effective   []string `json:"effective,omitempty" platform:"linux"`
	Inheritable []string `json:"inheritable,omitempty" platform:"linux"`
	Permitted   []string `json:"permitted,omitempty" platform:"linux"`
	Ambient     []string `json:"ambient,omitempty" platform:"linux"`
}
