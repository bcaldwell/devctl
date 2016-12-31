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

	"strings"

	"github.com/benjamincaldwell/devctl/post_command"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/spf13/cobra"
)

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Aliases: []string{"g"},
	Use:     "git",
	Short:   "a set of git alaises",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git called")
		postCommand.RunCommand("git " + strings.Join(args, " "))
	},
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "removes local branches which where deleted remotely",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := shell.Command("git", "branch", "--merged", "master").Output()
		if utilities.HandleError(err) {
			output, err = shell.Command("grep", "-v", "master").SetInput(string(output)).Output()
			if utilities.HandleError(err) {
				err = shell.Command("xargs", "-n", "1", "git", "branch", "-d").SetInput(string(output)).PrintOutput()
				utilities.HandleError(err)
			}
		}
	},
}

func init() {
	devctlCmd.AddCommand(gitCmd)
	gitCmd.AddCommand(cleanCmd)
}
