package event

import "github.com/fatih/color"

type Level uint32

type EventType uint32

// DetectType defines the cloud-native object type to scan
type DetectType uint32

// AlertType Defines details info type
type AlertType string

type RuntimeType uint32

type ClusterType string

type WeakpassService uint32

const (
	Low Level = iota
	Medium
	High
	Critical
	None

	Risk EventType = iota
	Invasion
	Info

	Image DetectType = iota
	Container
	IaC
	Cluster

	Docker RuntimeType = iota
	Containerd
	Remote
	Tarball
	Kubernetes

	ClusterKubernetes ClusterType = "kubernetes"
	ClusterOpenshift  ClusterType = "openshift"

	Asset            AlertType = "Asset"
	BasicImage       AlertType = "BasicImage"
	BasicContainer   AlertType = "BasicContainer"
	BasicCluster     AlertType = "BasicCluster"
	Vulnerability    AlertType = "Vulnerability"
	MaliciousFile    AlertType = "MaliciousFile"
	Backdoor         AlertType = "Backdoor"
	SensitiveFile    AlertType = "SensitiveFile"
	SensitiveEnv     AlertType = "SensitiveEnv"
	SensitiveHistory AlertType = "SensitiveHistory"
	AbnormalHistory  AlertType = "AbnormalHistory"
	Webshell         AlertType = "Webshell"
	Weakpass         AlertType = "Weakpass"
	IaCRisk          AlertType = "IaC"

	SSH WeakpassService = iota
	Redis
	Mysql
	Tomcat
	Env
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

	colorFns = map[Level]func(a ...interface{}) string{
		Low:      color.New(color.FgBlue).SprintFunc(),
		Medium:   color.New(color.FgYellow).SprintFunc(),
		High:     color.New(color.FgRed).SprintFunc(),
		Critical: color.New(color.FgRed).SprintFunc(),
		None:     color.New(color.FgWhite).SprintFunc(),
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

	toDetectType = map[DetectType]string{
		Image:     "Image",
		Container: "Container",
		IaC:       "IaC",
		Cluster:   "cluster",
	}

	fromDetectType = map[string]DetectType{
		"Image":     Image,
		"Container": Container,
		"IaC":       IaC,
		"cluster":   Cluster,
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
)
