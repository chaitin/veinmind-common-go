package group

import api "github.com/chaitin/libveinmind/go"

const groupPath = "/etc/group"

func ParseImageGroup(image api.Image) ([]Entry, error) {
	f, err := image.Open(groupPath)
	if err != nil {
		return nil, err
	}

	return parseReader(f)
}
