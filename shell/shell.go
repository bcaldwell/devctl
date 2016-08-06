package shell

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/benjamincaldwell/devctl/printer"
)

var MainInterface SessionInterface = new(session)

type SessionInterface interface {
	Command(name string, arg ...string) SessionInterface
	SetInput(s string) SessionInterface
	SetDir(s string) SessionInterface
	Run() error
	Output() ([]byte, error)
	PrintOutput() error
	New() SessionInterface
}

type session struct {
	cmd *exec.Cmd
}

func (c *session) New() SessionInterface {
	s := new(session)
	s.cmd = &exec.Cmd{}
	return s
}

func (c *session) Command(name string, arg ...string) SessionInterface {
	c.cmd = exec.Command(name, arg...)
	return c
}

func (c *session) SetInput(s string) SessionInterface {
	c.cmd.Stdin = strings.NewReader(s)
	return c
}

func (c *session) SetDir(s string) SessionInterface {
	c.cmd.Dir = s
	return c
}

func (c *session) Run() error {
	return c.cmd.Run()
}

func (c *session) Output() ([]byte, error) {
	return c.cmd.Output()
}

func (c *session) PrintOutput() error {
	cmdReader, _ := c.cmd.StdoutPipe()
	outScanner := bufio.NewScanner(cmdReader)
	go func() {
		for outScanner.Scan() {
			printer.InfoBar(outScanner.Text())
		}
	}()
	cmdReader, _ = c.cmd.StderrPipe()
	errScanner := bufio.NewScanner(cmdReader)
	go func() {
		for errScanner.Scan() {
			text := errScanner.Text()
			if strings.Contains(strings.ToLower(text), "warn") {
				printer.WarningBar(text)
			} else {
				printer.ErrorBar(text)
			}
		}
	}()
	err := c.cmd.Run()
	return err
}

func Session() SessionInterface {
	return MainInterface.New()
}

func New() SessionInterface {
	return Session()
}

// Command creates a new session and sets up the command with the proper arguments
func Command(name string, arg ...string) SessionInterface {
	cmd := MainInterface.New()
	return cmd.Command(name, arg...)
}
