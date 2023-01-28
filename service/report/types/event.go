package types

import (
	"bytes"
	"encoding/json"
)

type EventType uint32

const (
	Risk EventType = iota
	Invasion
	Info
)

var (
	toEventType = map[EventType]string{
		Risk:     "Risk",
		Invasion: "Invasion",
		Info:     "Info",
	}

	fromEventType = map[string]EventType{
		"Risk":     Risk,
		"Invasion": Invasion,
		"Info":     Info,
	}
)

func (e EventType) String() string {
	switch e {
	case Risk:
		return "risk"
	case Invasion:
		return "invasion"
	case Info:
		return "info"
	default:
		return "unknown"
	}
}

func (e EventType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toEventType[e])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (e *EventType) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	switch i.(type) {
	case uint32:
		*e = i.(EventType)
	case float64:
		*e = EventType(i.(float64))
	case string:
		*e = fromEventType[i.(string)]
	}

	return nil
}
