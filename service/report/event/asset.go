package event

import (
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type AssetDetail struct {
	OS           AssetOSDetail         `json:"os"`
	PackageInfos []AssetPackageDetails `json:"package_infos"`
}

func init() {
	RegisterAlter(Asset, func() *DetailInfo { return NewDetailInfo(&AssetDetail{}) })
}

type AssetOSDetail struct {
	Family string `json:"family"`
	Name   string `json:"name"`
	Eosl   bool   `json:"EOSL,omitempty"`
}

type AssetPackageDetails struct {
	Type     string               `json:"type"`
	FilePath string               `json:"file_path"`
	Packages []AssetPackageDetail `json:"packages"`
}

type AssetPackageDetail struct {
	Name            string `json:"name"`
	Version         string `json:"version"`
	Release         string `json:"release"`
	Epoch           int    `json:"epoch"`
	Arch            string `json:"arch"`
	SrcName         string `json:"srcName"`
	SrcVersion      string `json:"srcVersion"`
	SrcRelease      string `json:"srcRelease"`
	SrcEpoch        int    `json:"srcEpoch"`
	Modularitylabel string `json:"modularitylabel"`
	Indirect        bool   `json:"indirect"`
	License         string `json:"license"`
	Layer           string `json:"layer"`
}

func (a *AssetDetail) RenderTable(id string, level string) []table.Row {
	data := make([]table.Row, 0)
	for _, pkgs := range a.PackageInfos {
		for _, pkg := range pkgs.Packages {
			data = append(data, table.Row{simply(id),
				strings.Join([]string{a.OS.Family, a.OS.Name}, ":"),
				pkgs.Type, pkg.Name, pkg.Version, pkgs.FilePath,
			})
		}
	}
	return data
}

func (a *AssetDetail) RenderTableHeader() table.Row {
	return table.Row{"FROM", "OS", "TYPE", "NAME", "VERSION", "PATH"}
}

func (a *AssetDetail) RenderTableTitle() string {
	return "APPLICATION-INFO"
}

func (a *AssetDetail) RenderRowConfig() table.RowConfig {
	return table.RowConfig{}
}

func (a *AssetDetail) RenderColumnConfig() []table.ColumnConfig {
	return []table.ColumnConfig{
		{Number: 1, WidthMax: 12, AutoMerge: true},
		{Number: 2, WidthMax: 18, AutoMerge: true},
		{Number: 3, WidthMax: 12},
		{Number: 4, WidthMax: 32, Align: text.AlignJustify, AutoMerge: true},
		{Number: 5, WidthMax: 16, Align: text.AlignJustify, AutoMerge: true},
		{Number: 6, WidthMax: 32, Align: text.AlignJustify},
	}
}
