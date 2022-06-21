package group

import (
	"bufio"
	"errors"
	"github.com/chaitin/libveinmind/go/plugin/log"
	"io"
	"strings"
)

type Entry struct {
	GroupName string
	Password  string
	Gid       string
	UserList  string
}

func parseReader(reader io.Reader) ([]Entry, error) {
	scanner := bufio.NewScanner(reader)
	entries := []Entry{}
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := parseLine(line)
		if err != nil {
			log.Error(err)
			continue
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func parseLine(line string) (Entry, error) {
	lineSplit := strings.Split(line, ":")
	if len(lineSplit) != 4 {
		return Entry{}, errors.New("group: field length not match")
	}

	return Entry{
		GroupName: lineSplit[0],
		Password:  lineSplit[1],
		Gid:       lineSplit[2],
		UserList:  lineSplit[3],
	}, nil
}
