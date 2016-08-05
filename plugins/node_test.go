package plugins

import (
	"fmt"
	"testing"

	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/shell/shell_mock"
)

func TestUsed(t *testing.T) {
	shell.Command("ls").Run()
	fmt.Println(shellMock.Commands)
}
