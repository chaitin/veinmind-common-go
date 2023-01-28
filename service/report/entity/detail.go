package entity

type AlertDetail struct {
	FilterFileDetail             *FilterFileDetail             `json:"filter_file_detail,omitempty"`
	MaliciousFileDetail          *MaliciousFileDetail          `json:"malicious_file_detail,omitempty"`
	WeakpassDetail               *WeakpassDetail               `json:"weakpass_detail,omitempty"`
	BackdoorDetail               *BackdoorDetail               `json:"backdoor_detail,omitempty"`
	SensitiveFileDetail          *SensitiveFileDetail          `json:"sensitive_file_detail,omitempty"`
	SensitiveEnvDetail           *SensitiveEnvDetail           `json:"sensitive_env_detail,omitempty"`
	SensitiveDockerHistoryDetail *SensitiveDockerHistoryDetail `json:"sensitive_docker_history_detail,omitempty"`
	HistoryDetail                *HistoryDetail                `json:"history_detail,omitempty"`
	AssetDetail                  *AssetDetail                  `json:"asset_detail,omitempty"`
	WebshellDetail               *WebshellDetail               `json:"webshell_detail,omitempty"`
	ImageBasicDetail             *ImageBasicDetail             `json:"image_basic_detail,omitempty"`
	ContainerBasicDetail         *ContainerBasicDetail         `json:"container_basic_detail,omitempty"`
	IaCDetail                    *IaCDetail                    `json:"iac_detail,omitempty"`
}
