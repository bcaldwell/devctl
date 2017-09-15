package cmd

import (
	"os"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/plugins"
	printer "github.com/benjamincaldwell/go-printer"
	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: up,
}

func init() {
	devctlCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func up(cmd *cobra.Command, args []string) {
	config := new(parser.ProjectConfigStruct)
	config.ParseFileDefault()

	// create .devctl folder
	os.Mkdir(".devctl", 0700)

	pluginsUsed := plugins.Used(config)

	pluginCount := len(pluginsUsed)

	for i, plugin := range pluginsUsed {
		// get tasks
		tasks, err := plugin.UpTasks(config)

		// run pre checks
		check, err := plugins.RunChecks(tasks)
		if err != nil {
			printer.ErrorBar("%s", err)
		}

		// log completed check
		if check {
			printer.Success("%d/%d %s (Already completed)", i+1, pluginCount, plugin)
			continue
		}

		// Run task if necessary
		printer.InfoLineTextTop("%d/%d %s", i+1, pluginCount, plugin)
		err = plugins.RunTasks(tasks)
		printer.InfoLineBottom()

		// Run post checks
		if err == plugins.CheckFailedAfterTaskErr {
			os.Exit(1)
		}
	}
}
