package report

func (l Level) String() string {
	switch l {
	case Low:
		return "low"
	case Medium:
		return "medium"
	case High:
		return "high"
	case Critical:
		return "critical"
	default:
		return "unknown"
	}
}

func (t DetectType) String() string {
	switch t {
	case Image:
		return "image"
	case Container:
		return "container"
	case IaC:
		return "iac"
	default:
		return "unknown"
	}
}

func (t EventType) String() string {
	switch t {
	case Risk:
		return "risk"
	case Invasion:
		return "invasion"
	default:
		return "unknown"
	}
}

func (t AlertType) String() string {
	switch t {
	case Vulnerability:
		return "vulnerability"
	case MaliciousFile:
		return "malicious"
	case Backdoor:
		return "backdoor"
	case Sensitive:
		return "sensitive"
	case AbnormalHistory:
		return "history"
	case Weakpass:
		return "weakpass"
	case IaCRisk:
		return "iac"
	case Basic:
		return "basic"
	case Asset:
		return "asset"
	default:
		return "unknown"
	}
}

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

func (s WeakpassService) String() string {
	switch s {
	case SSH:
		return "ssh"
	case Tomcat:
		return "tomcat"
	case Redis:
		return "redis"
	case Mysql:
		return "mysql"
	default:
		return "unknown"
	}
}
