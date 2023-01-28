package entity

type SensitiveFileDetail struct {
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
