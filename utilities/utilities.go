package utilities

import (
	"os"

	"github.com/benjamincaldwell/devctl/printer"
	"github.com/spf13/cobra"
)

// ErrorWithHelp Show show error message, help menu and exit
func ErrorWithHelp(cmd *cobra.Command, message string) {
	printer.Fail(message)
	cmd.Help()
	os.Exit(1)
}
