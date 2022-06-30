package messenger

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

type Logfunction func(m *Messenger, format string, args ...any)

type msgAction uint

const (
	FATAL msgAction = (1 << iota)
	LOG
	LOG2
	LOG3
	WARN
	INFO
	DEBUG
	DEBUG2
	DEBUG3
	PROMPT
	MESSAGE
	STACKTRACE
)

var defaultMsgmap = map[msgAction]string{
	FATAL:      "Fatal Error: %s\n",
	LOG:        "Logging: %s\n",
	LOG2:       "Logging level 2: %s\n",
	LOG3:       "Logging level 3: %s\n",
	WARN:       "Warning: %s\n",
	INFO:       "Info: %s\n",
	PROMPT:     "%s\nContinue? Y/N\n",
	MESSAGE:    "%s\n",
	DEBUG:      "Debug: %s\n",
	DEBUG2:     "Debug level 2: %s\n",
	DEBUG3:     "Debug level 3: %s\n",
	STACKTRACE: "\n*** [Stack Trace] ***\n",
}

type Messenger struct {
	logfunc     Logfunction
	logger      *log.Logger
	logout, out io.Writer
	loggingOn   bool
	debugOn     bool
	traceOn     bool
	msgmap      map[msgAction]string
}

func New(writers ...io.Writer) *Messenger {
	var lo, o io.Writer
	switch len(writers) {
	case 0:
		{
			lo = os.Stderr
			o = os.Stdout
		}
	case 1:
		{
			lo = writers[0]
			o = os.Stdout
		}
	case 2:
		{
			lo = writers[0]
			o = writers[1]
		}
	}
	m := Messenger{
		logout:    lo,
		out:       o,
		logfunc:   logit,
		loggingOn: true,
		debugOn:   false,
		traceOn:   false,
		msgmap:    defaultMsgmap,
	}
	m.logger = log.Default()
	m.logger.SetFlags(log.LstdFlags | log.Lshortfile)
	return &m
}

func (m *Messenger) SetActionPrefix(a msgAction, pf string) {
	m.msgmap[a] = pf
}

func (m *Messenger) ActionPrefix(a msgAction) string {
	if a > 0 && a <= STACKTRACE {
		return m.msgmap[a]
	} else {
		return ""
	}
}

func (m *Messenger) SetLogfunc(lf Logfunction) {
	m.logfunc = lf
}

func (m *Messenger) LogFunc() Logfunction {
	return m.logfunc
}

func (m *Messenger) ResetLoggerFlags() {
	m.logger.SetFlags(log.LstdFlags | log.Lshortfile)
}

func (m *Messenger) SetLoggerFlags(flags int) {
	m.logger.SetFlags(flags)
}

func (m *Messenger) LoggerFlags() int {
	return m.logger.Flags()
}

func (m *Messenger) Logout() io.Writer {
	return m.logout
}

func (m *Messenger) SetLogout(lo io.Writer) {
	m.logout = lo
	m.logger.SetOutput(m.logout)
}

// call with no arg to reset to stderr
func (m *Messenger) SetLogoutStr(file ...string) {
	if len(file) == 0 || len(file[0]) == 0 || file[0] == "-" {
		m.SetLogout(os.Stderr)
		return
	}
	var f *os.File
	path, err := filepath.Abs(file[0])
	if err == nil {
		f, _ = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	}
	if f == nil {
		m.SetLogout(os.Stderr)
		m.logfunc(m, "Failure to open log file %s. Will default to Stderr.", file[0])
		return
	}
	m.SetLogout(f)
}

func (m *Messenger) LoggingOff() io.Writer {
	m.loggingOn = false
	ret := m.Logout()
	m.SetLogout(io.Discard)
	return ret
}

//call with no arg to reset to stderr
func (m *Messenger) LoggingOn(writer ...io.Writer) {
	var lo = io.Writer(os.Stderr)
	if len(writer) > 0 {
		lo = writer[0]
	}
	m.loggingOn = true
	m.SetLogout(lo)
}

func (m *Messenger) Out() io.Writer {
	return m.out
}

func (m *Messenger) SetOut(o io.Writer) {
	m.out = o
}

func (m *Messenger) ResetActionPrefixes() {
	m.msgmap = defaultMsgmap
}
func (m *Messenger) Write(b []byte) (int, error) {
	return m.out.Write(b)
}
