package report

// Object represents various cloud-native objects
type Object struct {
	Raw interface{} `json:"-"`

	// container runtime
	ID          string      `json:"id,omitempty"`
	Type        DetectType  `json:"type,omitempty"`
	RuntimeType RuntimeType `json:"runtime_type,omitempty"`
	RuntimeRoot string      `json:"runtime_root,omitempty"`

	// cluster config
	ClusterType       ClusterType `json:"cluster_type"`
	ClusterConfigPath string      `json:"cluster_config_path"`
	ClusterConfigByte []byte      `json:"cluster_config_byte"`
}
