package event

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

type SensitiveDetail struct{}

func init() {
	RegisterAlter(SensitiveFile, func() *DetailInfo { return NewDetailInfo(&SensitiveFileDetail{}) })
	RegisterAlter(SensitiveEnv, func() *DetailInfo { return NewDetailInfo(&SensitiveEnvDetail{}) })
	RegisterAlter(SensitiveHistory, func() *DetailInfo { return NewDetailInfo(&SensitiveDockerHistoryDetail{}) })
}

type SensitiveFileDetail struct {
	FileDetail
	SensitiveDetail
	RuleID                       int64   `json:"rule_id"`
	RuleName                     string  `json:"rule_name"`
	RuleDescription              string  `json:"rule_description"`
	ContextContent               string  `json:"context_content"`
	ContextContentHighlightRange []int64 `json:"context_content_highlight_range"`
}

type SensitiveEnvDetail struct {
	SensitiveDetail
	Key             string `json:"key"`
	Value           string `json:"value"`
	RuleID          int64  `json:"rule_id"`
	RuleName        string `json:"rule_name"`
	RuleDescription string `json:"rule_description"`
}

type SensitiveDockerHistoryDetail struct {
	SensitiveDetail
	Value           string `json:"value"`
	RuleID          int64  `json:"rule_id"`
	RuleName        string `json:"rule_name"`
	RuleDescription string `json:"rule_description"`
}

func (s *SensitiveFileDetail) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{simply(id), level, "File", s.Path, s.ContextContent, s.RuleName, s.RuleDescription})
	return data
}

func (s *SensitiveEnvDetail) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{simply(id), level, "Env", s.Key, s.Value, s.RuleName, s.RuleDescription})
	return data
}

func (s *SensitiveDockerHistoryDetail) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{simply(id), level, "History", "", s.Value, s.RuleName, s.RuleDescription})
	return data
}

func (s *SensitiveDetail) RenderTableHeader() table.Row {
	return table.Row{"FROM", "LEVEL", "TYPE", "Key", "Value", "RuleName", "RuleDescription"}
}

func (s *SensitiveDetail) RenderTableTitle() string {
	return "Sensitive Info"
}

func (s *SensitiveDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (s *SensitiveDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 12, AutoMerge: true},
		{Number: 2, WidthMax: 12},
		{Number: 3, WidthMax: 12},
		{Number: 4, WidthMax: 12},
		{Number: 5, WidthMax: 32},
		{Number: 6, WidthMax: 32},
		{Number: 7, WidthMax: 32},
	}
}
