package entity

type WebshellDetail struct {
	FileDetail
	Type   string `json:"type"`
	Reason string `json:"reason"`
	Engine string `json:"engine"`
}
