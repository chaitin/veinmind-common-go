package passwd

import (
	api "github.com/chaitin/libveinmind/go"
)

const (
	passwdPath = "/etc/passwd"
)

func ParseImagePasswd(image api.Image) ([]Entry, error) {
	f, err := image.Open(passwdPath)
	if err != nil {
		return nil, err
	} else {
		defer f.Close()
	}

	return parseReader(f)
}
