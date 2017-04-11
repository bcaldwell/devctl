package cmd

import (
	"os"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/plugins"
	"github.com/benjamincaldwell/devctl/postCommand"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/go-printer"
	"github.com/spf13/cobra"
)

var cfgFile string

var cmdWhitelist = [2]string{"update", "setup"}

var Version string
var BuildDate string

var Verbose bool
var DryRun bool

// devctlCmd represents the base command when called without any subcommands
var devctlCmd = &cobra.Command{
	Use:           "devctl",
	Short:         "devctl is enables developers to manage their development environments across different projects.",
	Long:          `devctl is enables developers to manage their development environments across different projects. To get started clone a respositories with devctl clone user/repo`,
	SilenceErrors: true,
	// PersistentPreRun: initConfig,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run:               devctl,
	PersistentPostRun: persistentPostRun,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the devctlCmd.
func Execute() {
	if err := devctlCmd.Execute(); err != nil {
		config := new(parser.ConfigurationStruct)
		config.ParseFileDefault()

		pluginsUsed := plugins.Used(config)
		scripts := generateScriptMap(config, pluginsUsed)
		if len(os.Args) >= 2 {
			scriptName := os.Args[1]
			if val, ok := findScript(scriptName, scripts); ok {
				runScript(val, config, pluginsUsed)
			} else {
				devctlCmd.Help()
			}
		}
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	devctlCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (located at $HOME/.devctlconfig)")
	devctlCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	devctlCmd.PersistentFlags().BoolVar(&DryRun, "dryrun", false, "dry run")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// devctlCmd.Flags().BoolP("help", "h", false, "Help message for command")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	cfgFile := os.Getenv("HOME") + "/.devctlconfig"

	shell.DryRun = DryRun
	postCommand.DryRun = DryRun

	printer.Verbose = Verbose

	if err := parser.ParseDevctlConfig(cfgFile); err != nil {
		printer.Warning("Warning: devctl config was not found or could not be parsed. %s", err)

		cmdUsed := os.Args[1]
		shouldExit := true
		for _, cmd := range cmdWhitelist {
			if cmd == cmdUsed {
				shouldExit = false
			}
		}
		if shouldExit {
			os.Exit(1)
		}
	}
}

func devctl(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func persistentPostRun(cmd *cobra.Command, args []string) {
	postCommand.Write()
}
