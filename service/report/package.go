// Package report provides report service for veinmind-runner
// and veinmind-plugin
package report

import (
	"os"
	"time"
)

type Level uint32

const (
	Low Level = iota
	Medium
	High
	Critical
	None
)

type DetectType uint32

const (
	Image DetectType = iota
	Container
	IaC
)

type EventType uint32

const (
	Risk EventType = iota
	Invasion
	Info
)

type AlertType uint32

const (
	Vulnerability AlertType = iota
	MaliciousFile
	Backdoor
	Sensitive
	AbnormalHistory
	Weakpass
	Asset
	Basic
	IaCRisk
)

type WeakpassService uint32

const (
	SSH WeakpassService = iota
	Redis
	Mysql
	Tomcat
	Env
)

type RuntimeType uint32

const (
	Docker RuntimeType = iota
	Containerd
	Remote
	Tarball
	Kubernetes
)

type AlertDetail struct {
	FilterFileDetail             *FilterFileDetail             `json:"filter_file_detail,omitempty"`
	MaliciousFileDetail          *MaliciousFileDetail          `json:"malicious_file_detail,omitempty"`
	WeakpassDetail               *WeakpassDetail               `json:"weakpass_detail,omitempty"`
	BackdoorDetail               *BackdoorDetail               `json:"backdoor_detail,omitempty"`
	SensitiveFileDetail          *SensitveFileDetail           `json:"sensitive_file_detail,omitempty"`
	SensitiveEnvDetail           *SensitiveEnvDetail           `json:"sensitive_env_detail,omitempty"`
	SensitiveDockerHistoryDetail *SensitiveDockerHistoryDetail `json:"sensitive_docker_history_detail,omitempty"`
	HistoryDetail                *HistoryDetail                `json:"history_detail,omitempty"`
	AssetDetail                  *AssetDetail                  `json:"asset_detail,omitempty"`
	ImageBasicDetail             *ImageBasicDetail             `json:"image_basic_detail,omitempty"`
	WebshellDetail               *WebshellDetail               `json:"webshell_detail,omitempty"`
	ContainerBasicDetail         *ContainerBasicDetail         `json:"container_basic_detail,omitempty"`
	IaCDetail                    *IaCDetail                    `json:"iac_detail,omitempty"`
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

type MountDetail struct {
	Destination string   `json:"destination"`
	Type        string   `json:"type,omitempty" platform:"linux,solaris"`
	Source      string   `json:"source,omitempty"`
	Options     []string `json:"options,omitempty"`
	VolumeName  string   `json:"volume_name,omitempty"`
	Permission  string   `json:"permission,omitempty"`
}

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

type MaliciousFileDetail struct {
	FileDetail
	Engine        string `json:"engine"`
	MaliciousType string `json:"malicious_type"`
	MaliciousName string `json:"malicious_name"`
}

type WeakpassDetail struct {
	Username string          `json:"username"`
	Password string          `json:"password"`
	Service  WeakpassService `json:"service"`
	Path     string          `json:"path"`
}

type BackdoorDetail struct {
	FileDetail
	Description string `json:"description"`
}

type SensitveFileDetail struct {
	FileDetail
	RuleID                       int64   `json:"rule_id"`
	RuleName                     string  `json:"rule_name"`
	RuleDescription              string  `json:"rule_description"`
	ContextContent               string  `json:"context_content"`
	ContextContentHighlightRange []int64 `json:"context_content_highlight_range"`
}

type SensitiveEnvDetail struct {
	Key             string `json:"key"`
	Value           string `json:"value"`
	RuleID          int64  `json:"rule_id"`
	RuleName        string `json:"rule_name"`
	RuleDescription string `json:"rule_description"`
}

type SensitiveDockerHistoryDetail struct {
	Value           string `json:"value"`
	RuleID          int64  `json:"rule_id"`
	RuleName        string `json:"rule_name"`
	RuleDescription string `json:"rule_description"`
}

type HistoryDetail struct {
	Instruction string `json:"instruction"`
	Content     string `json:"content"`
	Description string `json:"description"`
}

type AssetDetail struct {
	OS           AssetOSDetail             `json:"os"`
	PackageInfos []AssetPackageDetails     `json:"package_infos"`
	Applications []AssetApplicationDetails `json:"applications"`
}

type AssetOSDetail struct {
	Family string `json:"family"`
	Name   string `json:"name"`
	Eosl   bool   `json:"EOSL,omitempty"`
}

type AssetPackageDetails struct {
	FilePath string               `json:"file_path"`
	Packages []AssetPackageDetail `json:"packages"`
}

type AssetApplicationDetails struct {
	Type     string               `json:"type"`
	FilePath string               `json:"file_path,omitempty"`
	Packages []AssetPackageDetail `json:"packages"`
}

type AssetPackageDetail struct {
	Name            string `json:"name"`
	Version         string `json:"version"`
	Release         string `json:"release"`
	Epoch           int    `json:"epoch"`
	Arch            string `json:"arch"`
	SrcName         string `json:"srcName"`
	SrcVersion      string `json:"srcVersion"`
	SrcRelease      string `json:"srcRelease"`
	SrcEpoch        int    `json:"srcEpoch"`
	Modularitylabel string `json:"modularitylabel"`
	Indirect        bool   `json:"indirect"`
	License         string `json:"license"`
	Layer           string `json:"layer"`
}

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
	Runtime         RuntimeType       `json:"runtime"`
	RuntimeUniqDesc string            `json:"runtime_uniq_desc,omitempty"`
	Hostname        string            `json:"hostname"`
	ImageID         string            `json:"imageID"`
	Privileged      bool              `json:"privileged,omitempty"`
	RootProcess     RootProcessDetail `json:"process"`
	Mounts          []MountDetail     `json:"mounts,omitempty"`
	Processes       []ProcessDetail   `json:"processes,omitempty"`
}

type WebshellDetail struct {
	FileDetail
	Type   string `json:"type"`
	Reason string `json:"reason"`
	Engine string `json:"engine"`
}

type GeneralDetail []byte

type IaCDetail struct {
	RuleInfo IaCRule `json:"rule_info"`
	FileInfo IaCData `json:"file_info"`
}

type IaCData struct {
	StartLine int64
	EndLine   int64
	FilePath  string
	Original  string
}

type IaCRule struct {
	Id          string
	Name        string
	Description string
	Reference   string
	Severity    string
	Solution    string
	Type        string
}

type ReportEvent struct {
	ID             string          `json:"id"`
	RuntimeType    RuntimeType     `json:"runtime_type"`
	RuntimeRoot    string          `json:"runtime_root"`
	Time           time.Time       `json:"time"`
	Level          Level           `json:"level"`
	DetectType     DetectType      `json:"detect_type"`
	EventType      EventType       `json:"event_type"`
	AlertType      AlertType       `json:"alert_type"`
	AlertDetails   []AlertDetail   `json:"alert_details"`
	GeneralDetails []GeneralDetail `json:"general_details"`
}
