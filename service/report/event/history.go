package event

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

type HistoryDetail struct {
	Instruction string `json:"instruction"`
	Content     string `json:"content"`
	Description string `json:"description"`
}

func init() {
	RegisterAlter(AbnormalHistory, func() *DetailInfo { return NewDetailInfo(&HistoryDetail{}) })
}

func (h *HistoryDetail) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{simply(id), level, h.Instruction, h.Content, h.Description})
	return data
}

func (h *HistoryDetail) RenderTableHeader() table.Row {
	return table.Row{"FROM", "LEVEL", "Instruction", "Content", "Description"}
}

func (h *HistoryDetail) RenderTableTitle() string {
	return "History"
}

func (h *HistoryDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (h *HistoryDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 12, AutoMerge: true},
		{Number: 2, WidthMax: 12},
		{Number: 3, WidthMax: 32},
		{Number: 4, WidthMax: 32},
		{Number: 5, WidthMax: 32},
	}
}
