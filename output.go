// output.go
package messenger

import (
	"fmt"
	"os"
)

func (m *Messenger) output(fs string, args ...string) {
	var s string
	if len(args) == 0 {
		s = fs
	} else {
		s = fmt.Sprintf(fs, args[0])
	}
	if m.out == nil {
		m.out = os.Stdout
	}
	m.out.Write([]byte(s))
}
