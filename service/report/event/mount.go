package event

import "github.com/jedib0t/go-pretty/v6/table"

func init() {
	RegisterAlter(UnsafeMount, func() *DetailInfo { return NewDetailInfo(&UnSafeMountDetail{}) })
}

type MountEvent struct {
	Source      string `json:"source" yaml:"source"`
	Destination string `json:"destination" yaml:"destination"`
	Type        string `json:"type" yaml:"type"`
}

type UnSafeMountDetail struct {
	Mount MountEvent
}

func (a *UnSafeMountDetail) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	data = append(data, table.Row{a.Mount.Type, a.Mount.Source, a.Mount.Destination})
	return data
}

func (a *UnSafeMountDetail) RenderTableHeader() table.Row {
	return table.Row{"TYPE", "RESOURCE", "DESTINATION"}
}

func (a *UnSafeMountDetail) RenderTableTitle() string {
	return "Unsafe MOUNT RISK"
}

func (a *UnSafeMountDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (a *UnSafeMountDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 36},
		{Number: 2, WidthMax: 36},
		{Number: 3, WidthMax: 36},
	}
}
