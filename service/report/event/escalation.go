package event

import "github.com/jedib0t/go-pretty/v6/table"

type EscalationDetail struct {
	BinName     string `json:"binName"`
	Description string `json:"description"`
	FilePath    string `json:"filePath"`
	Mod         string `json:"mod"`
	Exp         string `json:"exp"`
}

func init() {
	RegisterAlter(Escalation, func() *DetailInfo {
		return NewDetailInfo(&EscalationDetail{})
	})
}

func (e *EscalationDetail) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{simply(id), level, e.BinName, e.FilePath, e.Description, e.Mod, e.Exp})
	return data
}

func (e *EscalationDetail) RenderTableHeader() table.Row {
	return table.Row{"FROM", "LEVEL", "BinName", "FilePath", "Description", "Mod", "Exp"}
}

func (e *EscalationDetail) RenderTableTitle() string {
	return "Privilege Escalation Risk"
}

func (e *EscalationDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}
func (e *EscalationDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 12, AutoMerge: true},
		{Number: 2, WidthMax: 12, AutoMerge: true},
		{Number: 3, WidthMax: 32, AutoMerge: true},
		{Number: 4, WidthMax: 32, AutoMerge: true},
		{Number: 5, WidthMax: 32, AutoMerge: true},
		{Number: 6, WidthMax: 32, AutoMerge: true},
		{Number: 7, WidthMax: 32, AutoMerge: true},
	}
}
