package types

import (
	"bytes"
	"encoding/json"
)

// DetectType defines the cloud-native object type to scan
type DetectType uint32

const (
	Image DetectType = iota
	Container
	IaC
	Cluster
)

var (
	toDetectType = map[DetectType]string{
		Image:     "Image",
		Container: "Container",
		IaC:       "IaC",
		Cluster:   "cluster",
	}

	fromDetectType = map[string]DetectType{
		"Image":     Image,
		"Container": Container,
		"IaC":       IaC,
		"cluster":   Cluster,
	}
)

func (d DetectType) String() string {
	switch d {
	case Image:
		return "image"
	case Container:
		return "container"
	case IaC:
		return "iac"
	default:
		return "unknown"
	}
}

func (d DetectType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toDetectType[d])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (d *DetectType) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	switch i.(type) {
	case uint32:
		*d = i.(DetectType)
	case float64:
		*d = DetectType(i.(float64))
	case string:
		*d = fromDetectType[i.(string)]
	}

	return nil
}
