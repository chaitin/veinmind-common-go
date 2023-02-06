package event

import (
	"strings"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func init() {
	RegisterAlter(BasicImage, func() *DetailInfo { return NewDetailInfo(&ImageBasicDetail{}) })
	RegisterAlter(BasicContainer, func() *DetailInfo { return NewDetailInfo(&ContainerBasicDetail{}) })
}

type ImageBasicDetail struct {
	References  []string `json:"references"`
	CreatedTime int64    `json:"created_time"`
	Env         []string `json:"env"`
	Entrypoint  []string `json:"entrypoint"`
	Cmd         []string `json:"cmd"`
	WorkingDir  string   `json:"working_dir"`
	Author      string   `json:"author"`
}

type ContainerBasicDetail struct {
	Name            string            `json:"name"`
	CreatedTime     int64             `json:"created_time"`
	State           string            `json:"state"`
	Runtime         RuntimeType       `json:"runtime"`
	RuntimeUniqDesc string            `json:"runtime_uniq_desc,omitempty"`
	Hostname        string            `json:"hostname"`
	ImageID         string            `json:"imageID"`
	Privileged      bool              `json:"privileged,omitempty"`
	RootProcess     RootProcessDetail `json:"process"`
	Mounts          []MountDetail     `json:"mounts,omitempty"`
	Processes       []ProcessDetail   `json:"processes,omitempty"`
}

type ClusterBasicDetail struct {
	// todo
}

func (i *ImageBasicDetail) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	point := i.Cmd
	if len(point) == 0 {
		point = i.Entrypoint
	}
	for _, r := range i.References {
		parsed, err := name.ParseReference(r)
		if err != nil {
			data = append(data, table.Row{r, "latest", simply(id), i.CreatedTime, point})
			continue
		}
		if strings.HasPrefix(parsed.Identifier(), "sha256:") {
			continue
		}
		data = append(data, table.Row{parsed.Context().String(), parsed.Identifier(), simply(id), i.CreatedTime, point})
	}
	return data
}

func (i *ImageBasicDetail) RenderTableHeader() table.Row {
	return table.Row{"REPOSITORY", "TAG", "IMAGE ID", "CREATED", "CMD/ENTRYPOINT"}
}

func (i *ImageBasicDetail) RenderTableTitle() string {
	return "IMAGE-INFO"
}

func (i *ImageBasicDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (i *ImageBasicDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 36},
		{Number: 2, WidthMax: 12},
		{Number: 3, WidthMax: 12, AutoMerge: true},
		{Number: 4, WidthMax: 12, Align: text.AlignLeft},
		{Number: 5, WidthMax: 16},
	}
}

func (i *ContainerBasicDetail) RenderTable(id string, level string) []table.Row {
	return []table.Row{
		{simply(id), simply(i.ImageID), i.Privileged, i.CreatedTime, i.State, i.Name},
	}
}

func (i *ContainerBasicDetail) RenderTableHeader() table.Row {
	return table.Row{"CONTAINER_ID", "IMAGE", "Is Privileged", "CREATED", "STATUS", "NAMES"}
}

func (i *ContainerBasicDetail) RenderTableTitle() string {
	return "CONTAINER-INFO"
}

func (i *ContainerBasicDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (i *ContainerBasicDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 12},
		{Number: 2, WidthMax: 12},
		{Number: 3, WidthMax: 16},
		{Number: 4, WidthMax: 12, Align: text.AlignLeft},
		{Number: 5, WidthMax: 16},
	}
}
