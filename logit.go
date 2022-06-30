package messenger

import (
	"fmt"
	"log"
	"os"
)

func logit(m *Messenger, fs string, args ...any) {
	s := fmt.Sprintf(fs, args...)
	if m.logger == nil {
		m.logger = log.Default()
	}
	if m.logger.Writer() == nil {
		m.logger.SetOutput(os.Stderr)
	}
	m.logger.Output(3, s)
}
