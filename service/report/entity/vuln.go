package entity

import "time"

type VulnDetail struct {
	ID         string       `json:"id"`
	Published  time.Time    `json:"published"`
	Aliases    []string     `json:"aliases"`
	Summary    string       `json:"summary"`
	Details    string       `json:"details"`
	References []References `json:"references"`
	Sources    []Source     `json:"sources"`
}

type References struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Source struct {
	OS       AssetOSDetail      `json:"os"`
	Type     string             `json:"type"`
	FilePath string             `json:"file_path"`
	Packages AssetPackageDetail `json:"packages"`
}
