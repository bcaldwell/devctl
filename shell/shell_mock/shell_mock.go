package shellMock

import (
	"strings"

	"github.com/benjamincaldwell/devctl/shell"
)

// list of commands that go run
var Commands []string

// top element is popped from this array and returned as the error value
var ErrorsToReturn []error

// top element is popped and returned from output function
var OutputToRetrun []string

func Reset() {
	Commands = Commands[:0]
	ErrorsToReturn = ErrorsToReturn[:0]
	OutputToRetrun = OutputToRetrun[:0]
}

func init() {
	shell.MainInterface = new(sessionMock)
}

type sessionMock struct {
	cmd   string
	dir   string
	input string
	env   map[string]string
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

func (c *sessionMock) SetDir(s string) shell.SessionInterface {
	c.dir = s
	return c
}

func (c *sessionMock) SetEnv(key, value string) shell.SessionInterface {
	c.env[key] = value
	return c
}

func (c *sessionMock) Output() ([]byte, error) {
	return make([]byte, 0), errorReturnValue()
}

func (c *sessionMock) PrintOutput() error {
	return errorReturnValue()
}

func (c *sessionMock) Run() error {
	Commands = append(Commands, c.cmd)

	return errorReturnValue()
}

func errorReturnValue() error {
	var returnErr error
	if len(ErrorsToReturn) > 0 {
		returnErr = ErrorsToReturn[0]
		ErrorsToReturn = ErrorsToReturn[1:]
	}
	return returnErr
}
