package entity

type AssetDetail struct {
	OS           AssetOSDetail             `json:"os"`
	PackageInfos []AssetPackageDetails     `json:"package_infos"`
	Applications []AssetApplicationDetails `json:"applications"`
}

type AssetOSDetail struct {
	Family string `json:"family"`
	Name   string `json:"name"`
	Eosl   bool   `json:"EOSL,omitempty"`
}

type AssetPackageDetails struct {
	FilePath string               `json:"file_path"`
	Packages []AssetPackageDetail `json:"packages"`
}

type AssetApplicationDetails struct {
	Type     string               `json:"type"`
	FilePath string               `json:"file_path,omitempty"`
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
