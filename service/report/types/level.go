package types

import (
	"bytes"
	"encoding/json"
)

// Level defines the severity of the risk
type Level uint32

const (
	Low Level = iota
	Medium
	High
	Critical
	None
)

var (
	toLevel = map[Level]string{
		Low:      "Low",
		Medium:   "Medium",
		High:     "High",
		Critical: "Critical",
		None:     "None",
	}

	fromLevel = map[string]Level{
		"Low":      Low,
		"Medium":   Medium,
		"High":     High,
		"Critical": Critical,
		"None":     None,
	}
)

func (l Level) String() string {
	switch l {
	case Low:
		return "low"
	case Medium:
		return "medium"
	case High:
		return "high"
	case Critical:
		return "critical"
	default:
		return "unknown"
	}
}

func (l Level) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toLevel[l])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (l *Level) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	switch v := i.(type) {
	case uint32:
		*l = Level(v)
	case float64:
		*l = Level(uint32(v))
	case string:
		*l = fromLevel[v]
	}

	return nil
}
