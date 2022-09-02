package passwd

import (
	api "github.com/chaitin/libveinmind/go"
)

const (
	passwdPath = "/etc/passwd"
)

func ParseFilesystemPasswd(fileSystem api.FileSystem) ([]Entry, error) {
	f, err := fileSystem.Open(passwdPath)
	if err != nil {
		return nil, err
	} else {
		defer f.Close()
	}

	return parseReader(f)
}
