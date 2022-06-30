// dbg-msg.go
package messenger

import (
	"fmt"
	"io"
)

func (m *Messenger) DebugMsg(w io.Writer, level int, msg string, args ...any) {
	err := fmt.Errorf(msg, args...)

	switch level {
	case 1:
		m.Catch(DEBUG, err, w)
	case 2:
		m.Catch(DEBUG2, err, w)
	case 3:
		m.Catch(DEBUG3, err, w)
	default:
		w.Write([]byte(err.Error()))
	}
}
