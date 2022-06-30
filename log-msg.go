package messenger

import (
	"fmt"
	"io"
)

func (m *Messenger) LogMsg(w io.Writer, level int, msg string, args ...any) {

	if !m.loggingOn {
		return
	}
	err := fmt.Errorf(msg, args...)

	switch level {
	case 1:
		m.Catch(LOG, err, w)
	case 2:
		m.Catch(LOG2, err, w)
	case 3:
		m.Catch(LOG3, err, w)
	default:
		w.Write([]byte(err.Error()))
	}
}
