package messenger

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (m *Messenger) termPrompt(s string) (string, error) {
	m.output(s)
	var reader = bufio.NewReader(os.Stdin)
	ans, err := reader.ReadString('\n')
	answer := strings.ToLower(string(ans[0]))
	return answer, err
}

func (m *Messenger) prompt(fs string, args ...any) {
	s := fmt.Sprintf(fs, args...)
	answer, err := m.termPrompt(s)

	if answer == "n" || answer == "\r" || answer == "\n" || err != nil {
		m.output("Exiting...\n")
		os.Exit(int(FATAL))
	}
}
