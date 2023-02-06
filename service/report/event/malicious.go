package event

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

type MaliciousFileDetail struct {
	FileDetail
	Engine        string `json:"engine"`
	MaliciousType string `json:"malicious_type"`
	MaliciousName string `json:"malicious_name"`
}

func init() {
	RegisterAlter(MaliciousFile, func() *DetailInfo { return NewDetailInfo(&MaliciousFileDetail{}) })
}

func (m *MaliciousFileDetail) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{simply(id), level, m.MaliciousType, m.MaliciousName, m.CalcSize(), m.Path})
	return data
}

func (m *MaliciousFileDetail) RenderTableHeader() table.Row {
	return table.Row{"FROM", "LEVEL", "Type", "Name", "Size", "Path"}
}

func (m *MaliciousFileDetail) RenderTableTitle() string {
	return "Malicious File"
}

func (m *MaliciousFileDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (m *MaliciousFileDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 12, AutoMerge: true},
		{Number: 2, WidthMax: 12},
		{Number: 3, WidthMax: 12},
		{Number: 4, WidthMax: 32},
		{Number: 5, WidthMax: 12},
		{Number: 6, WidthMax: 32},
	}
}
