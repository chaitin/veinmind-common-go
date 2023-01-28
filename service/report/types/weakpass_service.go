package types

import (
	"bytes"
	"encoding/json"
)

type WeakpassService uint32

const (
	SSH WeakpassService = iota
	Redis
	Mysql
	Tomcat
	Env
)

var (
	toWeakpassService = map[WeakpassService]string{
		SSH:    "SSH",
		Redis:  "Redis",
		Mysql:  "Mysql",
		Tomcat: "Tomcat",
		Env:    "Env",
	}

	fromWeakpassService = map[string]WeakpassService{
		"SSH":    SSH,
		"Redis":  Redis,
		"Mysql":  Mysql,
		"Tomcat": Tomcat,
		"Env":    Env,
	}
)

func (w WeakpassService) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toWeakpassService[w])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (w *WeakpassService) UnmarshalJSON(b []byte) error {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	switch i.(type) {
	case uint32:
		*w = i.(WeakpassService)
	case float64:
		*w = WeakpassService(i.(float64))
	case string:
		*w = fromWeakpassService[i.(string)]
	}

	return nil
}

func (w WeakpassService) String() string {
	switch w {
	case SSH:
		return "ssh"
	case Tomcat:
		return "tomcat"
	case Redis:
		return "redis"
	case Mysql:
		return "mysql"
	default:
		return "unknown"
	}
}
