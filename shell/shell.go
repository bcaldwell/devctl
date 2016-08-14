package shell

import (
	"bufio"
	"io"
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
	cmd   *exec.Cmd
	dir   string
	stdin io.Reader
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
	c.stdin = strings.NewReader(s)
	return c
}

func (c *session) SetDir(s string) SessionInterface {
	c.dir = s
	return c
}

func (c *session) applySettings() {
	if c.dir != "" {
		c.cmd.Dir = c.dir
	}
	if c.stdin != nil {
		c.cmd.Stdin = c.stdin
	}
}

func (c *session) Run() error {
	c.applySettings()
	return c.cmd.Run()
}

func (c *session) Output() ([]byte, error) {
	c.applySettings()
	return c.cmd.Output()
}

func (c *session) PrintOutput() error {
	c.applySettings()
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
			lowerText := strings.ToLower(text)
			if strings.Contains(lowerText, "warn") {
				printer.WarningBar(text)
			} else if strings.Contains(lowerText, "info") {
				printer.InfoBar(text)
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
