package event

import (
	"errors"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Event struct {
	*BasicInfo
	*DetailInfo
}

type BasicInfo struct {
	ID         string     `json:"id"`
	Time       time.Time  `json:"time"`
	Source     string     `json:"source"`     // 事件来源
	Level      Level      `json:"level"`      // 事件级别
	Object     Object     `json:"object"`     // 事件原对象
	EventType  EventType  `json:"event_type"` // 事件类型
	DetectType DetectType `json:"detect_type"`
	AlertType  AlertType  `json:"alert_type"`
}

type DetailInfo struct {
	AlertDetail AlertDetail `json:"alert_detail"`
}

func NewDetailInfo(a AlertDetail) *DetailInfo {
	return &DetailInfo{a}
}

// AlertDetail is interfaces of Entity that can be displayed
// if a custom detail satisfied AlertDetail, you can use
// RegisterAlter to add it into alertDetailMap so that it can be rendered
type AlertDetail interface {
	RenderTable(id string, level string) []table.Row
	RenderTableHeader() table.Row
	RenderTableTitle() string
	RenderRowConfig() table.RowConfig
	RenderColumnConfig() []table.ColumnConfig
}

var alertDetailMap = map[AlertType]func() *DetailInfo{}

func RegisterAlter(alertType AlertType, generateFunc func() *DetailInfo) error {
	if _, ok := alertDetailMap[alertType]; ok {
		return errors.New("already exits alert detail type")
	}
	alertDetailMap[alertType] = generateFunc
	return nil
}
