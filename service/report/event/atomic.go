package event

// atomic stored lower level event struct without render func
// and some common func

import (
	"fmt"
	"os"
	"strings"
)

type FileDetail struct {
	Path  string      `json:"path"`
	Perm  os.FileMode `json:"perm"`
	Size  int64       `json:"size"`
	Gname string      `json:"gname"`
	Gid   int64       `json:"gid"`
	Uid   int64       `json:"uid"`
	Uname string      `json:"uname"`
	Ctim  int64       `json:"ctim"`
	Mtim  int64       `json:"mtim"`
	Atim  int64       `json:"atim"`
}

type FilterFileDetail struct {
	FileDetail
	Type   os.FileMode `json:"type"`
	ELF    bool        `json:"elf"`
	Md5    string      `json:"md5"`
	Sha256 string      `json:"sha256"`
}

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

type MountDetail struct {
	Destination string   `json:"destination"`
	Type        string   `json:"type,omitempty" platform:"linux,solaris"`
	Source      string   `json:"source,omitempty"`
	Options     []string `json:"options,omitempty"`
	VolumeName  string   `json:"volume_name,omitempty"`
	Permission  string   `json:"permission,omitempty"`
}

func (f *FileDetail) CalcSize() string {
	size := f.Size
	if size < 1024 {
		return fmt.Sprintf("%.2f B", float64(size)/float64(1))
	} else if size < (1024 * 1024) {
		return fmt.Sprintf("%.2f KB", float64(size)/float64(1024))
	} else if size < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f MB", float64(size)/float64(1024*1024))
	} else if size < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f GB", float64(size)/float64(1024*1024*1024))
	} else if size < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f TB", float64(size)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2f PB", float64(size)/float64(1024*1024*1024*1024*1024))
	}
}

func simply(sha string) string {
	r := strings.Replace(sha, "sha256:", "", -1)
	if len(r) > 12 {
		return r[0:12]
	}
	return r
}
