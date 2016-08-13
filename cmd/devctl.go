// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/benjamincaldwell/devctl/post_command"
	"github.com/benjamincaldwell/devctl/printer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// devctlCmd represents the base command when called without any subcommands
var devctlCmd = &cobra.Command{
	Use:   "devctl",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
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
		scripts := generateScriptMap()
		if len(os.Args) >= 2 {
			scriptName := os.Args[1]
			if val, ok := findScript(scriptName, scripts); ok {
				postCommand.RunCommand(val.Command)
			} else {
				devctlCmd.Help()
			}
			postCommand.Write()
		}
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	devctlCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.devctlconfig)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	devctlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		fmt.Print(cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		cfgFile := os.Getenv("HOME") + "/.devctlconfig"
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName(".devctlconfig") // name of config file (without extension)
	viper.AddConfigPath("$HOME")         // adding home directory as first search path
	viper.AutomaticEnv()                 // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// 	// fmt.Println("Using config file:", viper.ConfigFileUsed())
		// } else {
		// 	// fmt.Println("Error reading config " + cfgFile + " :")
		// 	// fmt.Print(err)
	} else {
		printer.Warning("Warning: devctl config was not found")
	}
}

func devctl(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func persistentPostRun(cmd *cobra.Command, args []string) {
	postCommand.Write()
}
