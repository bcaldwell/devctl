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
	"strings"

	"github.com/benjamincaldwell/devctl/post_command"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/spf13/cobra"
)

var force bool

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Aliases: []string{"g"},
	Use:     "git",
	Short:   "a set of git alaises",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		postCommand.RunCommand("git " + strings.Join(args, " "))
	},
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "removes local branches which where deleted remotely",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if force {
			postCommand.RunCommand("sh -c \"git fetch -p && git branch -vv | awk '/: gone]/{print $1}' | xargs git branch -D\"")
		} else {
			postCommand.RunCommand("sh -c \"git fetch -p && git branch -vv | awk '/: gone]/{print $1}' | xargs git branch -d\"")
		}
	},
}

var yoloCmd = &cobra.Command{
	Use:   "yolo",
	Short: "removes local branches which where deleted remotely",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := shell.Command("git", "add", "-A").PrintOutput()
		if !utilities.HandleError(err) {
			err = shell.Command("git", "commit", "--amend", "--no-edit").PrintOutput()
			if !utilities.HandleError(err) {
				if currentBranch, err := gitOutput("rev-parse", "--abbrev-ref", "HEAD"); err == nil {
					// for some reason its not working in go...
					postCommand.RunCommand("git push -f origin " + string(currentBranch))
					// err = shell.Command("git", "push", "-f", "-v", "origin", string(currentBranch)+":"+string(currentBranch)).PrintOutput()
					// utilities.HandleError(err)
				}
			}
		}
	},
}

//

func init() {
	devctlCmd.AddCommand(gitCmd)
	gitCmd.AddCommand(cleanCmd)
	gitCmd.AddCommand(yoloCmd)
	cleanCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "force")
}
