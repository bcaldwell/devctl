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

	"github.com/benjamincaldwell/devctl/post_command"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("open called")
	},
}

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Open github repo in default browser",
	Run: func(cmd *cobra.Command, args []string) {
		postCommand.RunCommand("open `git remote -v | grep fetch | awk '{print $2}' | sed 's/git@/http:\\/\\//' | sed 's/com:/com\\//'`| head -n1")
	},
}

func init() {
	devctlCmd.AddCommand(openCmd)
	openCmd.AddCommand(repoCmd)
}
