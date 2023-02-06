package event

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

type WeakpassDetail struct {
	Username string          `json:"username"`
	Password string          `json:"password"`
	Service  WeakpassService `json:"service"`
	Path     string          `json:"path"`
}

func init() {
	RegisterAlter(Weakpass, func() *DetailInfo { return NewDetailInfo(&WeakpassDetail{}) })
}

func (w *WeakpassDetail) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{simply(id), level, w.Username, w.Password, w.Service, w.Path})
	return data
}

func (w *WeakpassDetail) RenderTableHeader() table.Row {
	return table.Row{"FROM", "LEVEL", "User", "Pass", "Service", "Path"}
}

func (w *WeakpassDetail) RenderTableTitle() string {
	return "WEAK PASS"
}

func (w *WeakpassDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (w *WeakpassDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 12, AutoMerge: true},
		{Number: 2, WidthMax: 12},
		{Number: 3, WidthMax: 12},
		{Number: 4, WidthMax: 12},
		{Number: 5, WidthMax: 32},
		{Number: 6, WidthMax: 32},
	}
}
