package plugins

import (
	"errors"
	"fmt"
	"testing"

	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/shell/shell_mock"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/stretchr/testify/assert"
)

func TestNvmInstalled(t *testing.T) {
	assert.Equal(t, utilities.CheckIfInstalled("nvm", "~/.nvm/nvm.sh"), true)
	assert.Len(t, shellMock.Commands, 2)
	assert.Equal(t, shellMock.Commands[0], "sh -c command -v nvm")
	assert.Equal(t, shellMock.Commands[1], "sh -c source ~/.nvm/nvm.sh && command -v nvm")

	shellMock.Reset()

	shellMock.ErrorsToReturn = append(shellMock.ErrorsToReturn, errors.New("no found"))
	assert.Equal(t, utilities.CheckIfInstalled("nvm", "~/.nvm/nvm.sh"), false)

	shellMock.Reset()
}

func TestUsed(t *testing.T) {
	shell.Command("ls").Run()
	fmt.Println(shellMock.Commands)
}
