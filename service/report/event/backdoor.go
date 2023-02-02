package event

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

func init() {
	RegisterAlter(Backdoor, func() *DetailInfo { return NewDetailInfo(&BackdoorDetail{}) })
}

type BackdoorDetail struct {
	FileDetail
	Description string `json:"description"`
}

func (b *BackdoorDetail) RenderTable(id string, level Level) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{simply(id), level.Color(), b.CalcSize(), b.Path, b.Description})
	return data
}

func (b *BackdoorDetail) RenderTableHeader() table.Row {
	return table.Row{"FROM", "LEVEL", "Size", "Path", "Description"}
}

func (b *BackdoorDetail) RenderTableTitle() string {
	return "BackDoor File"
}

func (b *BackdoorDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (b *BackdoorDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 12, AutoMerge: true},
		{Number: 2, WidthMax: 12},
		{Number: 3, WidthMax: 12},
		{Number: 4, WidthMax: 32},
		{Number: 5, WidthMax: 12},
		{Number: 6, WidthMax: 32},
	}
}
