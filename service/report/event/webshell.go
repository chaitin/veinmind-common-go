package event

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

type WebshellDetail struct {
	FileDetail
	Type   string `json:"type"`
	Reason string `json:"reason"`
	Engine string `json:"engine"`
}

func init() {
	RegisterAlter(Webshell, func() *DetailInfo { return NewDetailInfo(&WebshellDetail{}) })
}

func (w *WebshellDetail) RenderTable(id string, level Level) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{simply(id), level.Color(), w.Type, w.Engine, w.CalcSize(), w.Path, w.Reason})
	return data
}

func (w *WebshellDetail) RenderTableHeader() table.Row {
	return table.Row{"FROM", "level", "Type", "Engine", "SIZE", "PATH", "Reason"}
}

func (w *WebshellDetail) RenderTableTitle() string {
	return "WEBSHELL"
}

func (w *WebshellDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (w *WebshellDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 12, AutoMerge: true},
		{Number: 2, WidthMax: 12},
		{Number: 3, WidthMax: 12},
		{Number: 4, WidthMax: 12},
		{Number: 5, WidthMax: 32},
		{Number: 6, WidthMax: 32},
	}
}
