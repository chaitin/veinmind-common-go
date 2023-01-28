package entity

import "os"

type FileDetail struct {
	Path  string      `json:"path"`
	Perm  os.FileMode `json:"perm"`
	Size  int64       `json:"size"`
	Gname string      `json:"gname"`
	Gid   int64       `json:"gid"`
	Uid   int64       `json:"uid"`
	Uname string      `json:"uname"`
	Ctim  int64       `json:"ctim"`
	Mtim  int64       `json:"mtim"`
	Atim  int64       `json:"atim"`
}

type FilterFileDetail struct {
	FileDetail
	Type   os.FileMode `json:"type"`
	ELF    bool        `json:"elf"`
	Md5    string      `json:"md5"`
	Sha256 string      `json:"sha256"`
}
