package event

import "github.com/jedib0t/go-pretty/v6/table"

type EscapeDetail struct {
	Target string `json:"target"`
	Reason string `json:"reason"`
	Detail string `json:"detail"`
}

func init() {
	RegisterAlter(Escape, func() *DetailInfo { return NewDetailInfo(&EscapeDetail{}) })
}

func (e *EscapeDetail) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{simply(id), level, e.Target, e.Reason, e.Detail})
	return data
}

func (e *EscapeDetail) RenderTableHeader() table.Row {
	return table.Row{"FROM", "LEVEL", "Target", "Reason", "Detail"}
}

func (e *EscapeDetail) RenderTableTitle() string {
	return "Escape Risk"
}

func (e *EscapeDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (e *EscapeDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 12, AutoMerge: true},
		{Number: 2, WidthMax: 12},
		{Number: 3, WidthMax: 32},
		{Number: 4, WidthMax: 32},
		{Number: 5, WidthMax: 32},
	}
}
