package utilities

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// ErrorWithHelp Show show error message, help menu and exit
func ErrorWithHelp(cmd *cobra.Command, message string) {
	color.Red(message)
	cmd.Help()
	os.Exit(1)
}
