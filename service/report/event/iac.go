package event

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type IaCDetail struct {
	RuleInfo IaCRule `json:"rule_info"`
	FileInfo IaCData `json:"file_info"`
}

func init() {
	RegisterAlter(IaCRisk, func() *DetailInfo { return NewDetailInfo(&IaCDetail{}) })
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

func (i *IaCDetail) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{simply(id), level, i.RuleInfo.Name, i.RuleInfo.Description,
		i.RuleInfo.Solution, i.FileInfo.FilePath, i.FileInfo.Original})
	return data
}

func (i *IaCDetail) RenderTableHeader() table.Row {
	return table.Row{"FROM", "LEVEL", "NAME", "Desc", "Solution", "PATH", "CONTENT"}
}

func (i *IaCDetail) RenderTableTitle() string {
	return "IAC Risk"
}

func (i *IaCDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (i *IaCDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 36},
		{Number: 2, WidthMax: 12},
		{Number: 3, WidthMax: 12, AutoMerge: true},
		{Number: 4, WidthMax: 12, Align: text.AlignLeft},
		{Number: 5, WidthMax: 16},
		{Number: 6, WidthMax: 16},
		{Number: 7, WidthMax: 16},
	}
}
