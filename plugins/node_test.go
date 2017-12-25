package plugins

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bcaldwell/devctl/shell"
	"github.com/bcaldwell/devctl/shell/shell_mock"
	"github.com/stretchr/testify/assert"
)

func TestNvmInstalled(t *testing.T) {
	assert.Equal(t, nvmInstalled(), true)
	assert.Len(t, shellMock.Commands, 1)
	assert.Equal(t, shellMock.Commands[0], "sh -c source ~/.nvm/nvm.sh && command -v nvm")

	shellMock.Reset()

	shellMock.ErrorsToReturn = append(shellMock.ErrorsToReturn, errors.New("no found"))
	assert.Equal(t, nvmInstalled(), false)

	shellMock.Reset()
}

func TestUsed(t *testing.T) {
	shell.Command("ls").Run()
	fmt.Println(shellMock.Commands)
}
