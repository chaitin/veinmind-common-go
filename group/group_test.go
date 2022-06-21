package group

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var groupContent = `root:x:0:
daemon:x:1:
bin:x:2:
sys:x:3:
adm:x:4:syslog,bobi
tty:x:5:
disk:x:6:
lp:x:7:
mail:x:8:
news:x:9:
uucp:x:10:
`

func TestParseReader(t *testing.T) {
	r := strings.NewReader(groupContent)
	entries, err := parseReader(r)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, entries[0], Entry{
		GroupName: "root",
		Password:  "x",
		Gid:       "0",
		UserList:  "",
	})
}
