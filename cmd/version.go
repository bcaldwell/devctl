package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showDate bool

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
		if showDate {
			fmt.Println("Built: " + BuildDate)
		}
	},
}

func init() {
	devctlCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVarP(&showDate, "date", "d", false, "Show build date")
}
