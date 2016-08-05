package shellMock

import (
	"strings"

	"github.com/benjamincaldwell/devctl/shell"
)

var Commands []string

func init() {
	shell.MainInterface = new(sessionMock)
}

type sessionMock struct {
	cmd   string
	input string
}

func (c *sessionMock) New() shell.SessionInterface {
	return new(sessionMock)
}

func (c *sessionMock) Command(name string, arg ...string) shell.SessionInterface {
	args := append([]string{name}, arg...)
	c.cmd = strings.Join(args, " ")
	return c
}

func (c *sessionMock) SetInput(s string) shell.SessionInterface {
	c.input = s
	return c
}

func (c *sessionMock) Output() ([]byte, error) {
	return make([]byte, 0), nil
}

func (c *sessionMock) PrintOutput() error {
	return nil
}

func (c *sessionMock) Run() error {
	Commands = append(Commands, c.cmd)
	return nil
}
