package entity

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
