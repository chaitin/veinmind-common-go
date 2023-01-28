package entity

import "github.com/chaitin/veinmind-common-go/service/report/types"

type WeakpassDetail struct {
	Username string                `json:"username"`
	Password string                `json:"password"`
	Service  types.WeakpassService `json:"service"`
	Path     string                `json:"path"`
}
