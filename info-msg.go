// info-msg.go
package messenger

import (
	"fmt"
	"io"
)

func (m *Messenger) InfoMsg(w io.Writer, a msgAction, msg string, args ...any) {
	err := fmt.Errorf(msg, args...)

	if a&MESSAGE == MESSAGE {
		m.Catch(MESSAGE, err, w)
		return
	}
	if a&INFO == INFO {
		m.Catch(INFO, err, w)
		return
	}
	w.Write([]byte(err.Error()))
}
