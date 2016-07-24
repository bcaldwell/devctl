package utilities

import (
	"os"

	"github.com/benjamincaldwell/devctl/printer"
	"github.com/spf13/cobra"
)

type RunCommand struct {
	Desc    string
	Command string
}

// ErrorWithHelp Show show error message, help menu and exit
func ErrorWithHelp(cmd *cobra.Command, message string) {
	printer.Fail(message)
	cmd.Help()
	os.Exit(1)
}

func ErrorCheck(err error, message string) {
	if err == nil {
		printer.Success("%s successful", message)
	} else {
		printer.Fail("%s failed", message)
	}
}
