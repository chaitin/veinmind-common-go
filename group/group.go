package group

import api "github.com/chaitin/libveinmind/go"

const groupPath = "/etc/group"

func ParseFilesystemGroup(fileSystem api.FileSystem) ([]Entry, error) {
	f, err := fileSystem.Open(groupPath)
	if err != nil {
		return nil, err
	} else {
		defer f.Close()
	}

	return parseReader(f)
}
