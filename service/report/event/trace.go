package event

import "github.com/jedib0t/go-pretty/v6/table"

func init() {
	RegisterAlter(TraceRisk, func() *DetailInfo { return NewDetailInfo(&TraceEvent{}) })
}

type TraceEvent struct {
	Name        string `json:"name" yaml:"name"`
	From        string `json:"from" yaml:"from"`
	Path        string `json:"path" yaml:"path"`
	Level       Level  `json:"level" yaml:"level"`
	Description string `json:"description" yaml:"description"`
	Detail      string `json:"detail" yaml:"detail"`
}

func (a *TraceEvent) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{simply(id), level, a.Name, a.Path, a.Description, a.Detail})
	return data
}

func (a *TraceEvent) RenderTableHeader() table.Row {
	return table.Row{"FROM", "LEVEL", "NAME", "PATH", "DES", "DETAIL"}
}

func (a *TraceEvent) RenderTableTitle() string {
	return "Security Trace Risk"
}

func (a *TraceEvent) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (a *TraceEvent) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 12, AutoMerge: true},
		{Number: 2, WidthMax: 12},
		{Number: 3, WidthMax: 36},
		{Number: 4, WidthMax: 36},
		{Number: 5, WidthMax: 36},
		{Number: 6, WidthMax: 36},
	}
}
