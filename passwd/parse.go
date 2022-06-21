package passwd

import (
	"bufio"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// An Entry contains all the fields for a specific user
type Entry struct {
	Username string
	Pass     string
	Uid      string
	Gid      string
	Gecos    string
	Home     string
	Shell    string
}

func parseReader(r io.Reader) ([]Entry, error) {
	lines := bufio.NewReader(r)
	entries := make([]Entry, 0)
	for {
		line, _, err := lines.ReadLine()
		if err != nil {
			break
		}
		entry, err := parseLine(string(copyBytes(line)))
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}
	return entries, nil
}

func parseLine(line string) (Entry, error) {
	fs := strings.Split(line, ":")
	if len(fs) != 7 {
		return Entry{}, errors.New("Unexpected number of fields in /etc/passwd")
	}

	return Entry{fs[0], fs[1], fs[2], fs[3], fs[4], fs[5], fs[6]}, nil
}

func copyBytes(x []byte) []byte {
	y := make([]byte, len(x))
	copy(y, x)
	return y
}
