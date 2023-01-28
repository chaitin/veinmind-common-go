package entity

type MountDetail struct {
	Destination string   `json:"destination"`
	Type        string   `json:"type,omitempty" platform:"linux,solaris"`
	Source      string   `json:"source,omitempty"`
	Options     []string `json:"options,omitempty"`
	VolumeName  string   `json:"volume_name,omitempty"`
	Permission  string   `json:"permission,omitempty"`
}
