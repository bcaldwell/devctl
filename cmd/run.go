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
	"sort"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/plugins"
	"github.com/benjamincaldwell/devctl/postCommand"
	"github.com/benjamincaldwell/devctl/printer"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/renstrom/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Aliases: []string{"r"},
	Use:     "run",
	Short:   "run predefined scripts from devctl.yml file or language defaults",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: run,
}

func init() {
	devctlCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func run(cmd *cobra.Command, args []string) {
	var scriptName string
	if len(args) == 0 {
		scriptName = "run"
	} else {
		scriptName = args[0]
	}

	config := new(parser.ConfigurationStruct)
	config.ParseFileDefault()

	pluginsUsed := plugins.Used(config)

	scripts := generateScriptMap(config, pluginsUsed)

	if val, ok := findScript(scriptName, scripts); ok {
		runScript(val, config, pluginsUsed)
	} else {
		printer.Fail("%s script not found", scriptName)
	}
}

func runScript(script utilities.RunCommand, config *parser.ConfigurationStruct, pluginsUsed []plugins.Plugin) {
	for _, i := range pluginsUsed {
		i.PreScript(config)
	}

	postCommand.RunCommand(script.Command)

	for _, i := range pluginsUsed {
		i.PostScript(config)
	}
}

func findScript(scriptName string, scripts map[string]utilities.RunCommand) (utilities.RunCommand, bool) {
	if val, ok := scripts[scriptName]; ok {
		return val, true
	}
	// fuzzy search
	keys := make([]string, len(scripts))
	i := 0
	for k := range scripts {
		keys[i] = k
		i++
	}
	fuzzyFind := fuzzy.RankFind(scriptName, keys)
	sort.Sort(fuzzyFind)

	if len(fuzzyFind) > 0 {
		return scripts[fuzzyFind[0].Target], true
	}
	return utilities.RunCommand{}, false
}

func generateScriptMap(config *parser.ConfigurationStruct, pluginsUsed []plugins.Plugin) map[string]utilities.RunCommand {
	scripts := make(map[string]utilities.RunCommand)

	for _, i := range pluginsUsed {
		scripts = mapmerge(scripts, i.Scripts(config))
	}

	// merge most important last
	scripts = mapmerge(scripts, config.Scripts)

	return scripts
}

func mapmerge(base map[string]utilities.RunCommand, mapsToMerge ...map[string]utilities.RunCommand) map[string]utilities.RunCommand {
	for _, m := range mapsToMerge {
		for key, value := range m {
			base[key] = value
		}
	}
	return base
}
