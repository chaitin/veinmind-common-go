// Package report provides report service for veinmind-runner
// and veinmind-plugin
package report

import (
	"time"

	"github.com/chaitin/veinmind-common-go/service/report/entity"
	"github.com/chaitin/veinmind-common-go/service/report/types"
)

type Event struct {
	ID             string                 `json:"id"`
	Object         entity.Object          `json:"object"`
	Time           time.Time              `json:"time"`
	Level          types.Level            `json:"level"`
	DetectType     types.DetectType       `json:"detect_type"`
	EventType      types.EventType        `json:"event_type"`
	AlertType      types.AlertType        `json:"alert_type"`
	AlertDetails   []entity.AlertDetail   `json:"alert_details"`
	GeneralDetails []entity.GeneralDetail `json:"general_details"`
}
