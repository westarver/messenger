// catch.go
package messenger

import (
	"io"
	"os"
)

//────────────────────┤ Catch ├────────────────────

func (m *Messenger) Catch(a msgAction, err error, writers ...io.Writer) error {
	var oldl, oldo io.Writer

	if a == 0 || err == nil {
		return err
	}

	if len(writers) > 0 {
		if writers[0] != nil {
			oldl = m.logout
			m.logout = writers[0]
		}
		if len(writers) > 1 {
			oldo = m.out
			m.out = writers[1]
		}

		m.logger.SetOutput(m.logout)
		defer func(io.Writer, io.Writer) {
			if oldl != nil {
				m.logout = oldl
				if m.logger != nil {
					m.logger.SetOutput(m.logout)
				}
			}
			if oldo != nil {
				m.out = oldo
			}
		}(oldl, oldo)
	}

	var stacktrace, s string
	em := err.Error()

	for a != 0 {
		if a&STACKTRACE == STACKTRACE {
			a ^= STACKTRACE
			stacktrace = string(m.Stacktrace())
			continue
		}
		if a&LOG == LOG {
			a ^= LOG
			if len(stacktrace) != 0 {
				s = em + "\n" + m.msgmap[STACKTRACE] + stacktrace
			} else {
				s = em
			}
			if m.loggingOn {
				m.logfunc(m, m.msgmap[LOG], s)
			}
			continue
		}
		if a&LOG2 == LOG2 {
			a ^= LOG2
			if stacktrace != "" {
				s = em + "\n" + m.msgmap[STACKTRACE] + stacktrace
			} else {
				s = em
			}
			if m.loggingOn {
				m.logfunc(m, m.msgmap[LOG2], s)
			}
			continue
		}
		if a&LOG3 == LOG3 {
			a ^= LOG3
			if stacktrace != "" {
				s = em + "\n" + m.msgmap[STACKTRACE] + stacktrace
			} else {
				s = em
			}
			if m.loggingOn {
				m.logfunc(m, m.msgmap[LOG3], s)
			}
			continue
		}

		if a&MESSAGE == MESSAGE {
			a ^= MESSAGE
			m.output(m.msgmap[MESSAGE], em)
			continue
		}
		if a&INFO == INFO {
			a ^= INFO
			m.output(m.msgmap[INFO], em)
			continue
		}
		if a&DEBUG == DEBUG {
			a ^= DEBUG
			if m.debugOn {
				m.output(m.msgmap[DEBUG], em)
			}
			continue
		}
		if a&DEBUG2 == DEBUG2 {
			a ^= DEBUG2
			if m.debugOn {
				m.output(m.msgmap[DEBUG2], em)
			}
			continue
		}
		if a&DEBUG3 == DEBUG3 {
			a ^= DEBUG3
			if m.debugOn {
				m.output(m.msgmap[DEBUG3], em)
			}
			continue
		}
		if a&WARN == WARN {
			a ^= WARN
			m.output(m.msgmap[WARN], em)
			continue
		}
		if a&PROMPT == PROMPT {
			a ^= PROMPT
			m.prompt(m.msgmap[PROMPT], em)
			continue
		}
		if a&FATAL == FATAL {
			m.output(m.msgmap[FATAL], em)
			m.output("Exiting...\n")
			os.Exit(int(FATAL))
		}
	}
	return err
}
