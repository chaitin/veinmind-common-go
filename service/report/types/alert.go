package types

import (
	"bytes"
	"encoding/json"
)

type AlertType uint32

const (
	Vulnerability AlertType = iota
	MaliciousFile
	Backdoor
	Sensitive
	AbnormalHistory
	Weakpass
	Asset
	Basic
	IaCRisk
)

var (
	toAlertType = map[AlertType]string{
		Vulnerability:   "Vulnerability",
		MaliciousFile:   "MaliciousFile",
		Backdoor:        "Backdoor",
		Sensitive:       "Sensitive",
		AbnormalHistory: "AbnormalHistory",
		Weakpass:        "Weakpass",
		Asset:           "Asset",
		Basic:           "Basic",
		IaCRisk:         "IaC",
	}

	fromAlertType = map[string]AlertType{
		"Vulnerability":   Vulnerability,
		"MaliciousFile":   MaliciousFile,
		"Backdoor":        Backdoor,
		"Sensitive":       Sensitive,
		"AbnormalHistory": AbnormalHistory,
		"Weakpass":        Weakpass,
		"Asset":           Asset,
		"Basic":           Basic,
		"IaC":             IaCRisk,
	}
)

func (a AlertType) String() string {
	switch a {
	case Vulnerability:
		return "vulnerability"
	case MaliciousFile:
		return "malicious"
	case Backdoor:
		return "backdoor"
	case Sensitive:
		return "sensitive"
	case AbnormalHistory:
		return "history"
	case Weakpass:
		return "weakpass"
	case IaCRisk:
		return "iac"
	case Basic:
		return "basic"
	case Asset:
		return "asset"
	default:
		return "unknown"
	}
}

func (a AlertType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toAlertType[a])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (a *AlertType) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	switch i.(type) {
	case uint32:
		*a = i.(AlertType)
	case float64:
		*a = AlertType(i.(float64))
	case string:
		*a = fromAlertType[i.(string)]
	}

	return nil
}
