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

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/codeskyblue/go-sh"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone the github repository to the source_dir",
	Long:  ``,
	Run:   clone,
}

var gitlab bool

func init() {
	devctlCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().BoolVarP(&gitlab, "gitlab", "l", false, "Clone from gitlab url")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

type cloneConfig struct {
	repo      string
	user      string
	url       string
	sourceDir string
	command   *cobra.Command
}

func clone(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		utilities.ErrorWithHelp(cmd, "\nMinimum of one argument is required\n ")
	}

	cfg := new(cloneConfig)
	cfg.sourceDir = parser.GetString("source_dir")
	cfg.command = cmd

	cfg.github()
	if gitlab {
		cfg.gitlab()
	}
	cfg.parseArgs(args)

	// validate(cfg)

	cfg.clone()
}

func (cfg *cloneConfig) parseArgs(args []string) {
	if isFullURL(args[0]) {
		cfg.url = args[0]
		cfg.user = ""
		cfg.repo = ""
		return
	}

	args = strings.Split(args[0], "/")

	if len(args) == 1 {
		cfg.repo = args[0]
	} else if len(args) == 2 {
		cfg.user = args[0]
		cfg.repo = args[1]
	}
}

func isFullURL(s string) bool {
	if len(strings.Split(s, ":")) == 2 {
		return true
	} else if strings.HasPrefix(s, "https://") || strings.HasPrefix(s, "http://") {
		return true
	}
	return false
}

func (cfg *cloneConfig) github() {
	cfg.url = "github.com"

	user := parser.GetString("github_user")
	if user != "" {
		cfg.user = user
	}
}

func (cfg *cloneConfig) gitlab() {
	cfg.url = parser.GetString("gitlab_url")
	if cfg.url == "" {
		utilities.ErrorWithHelp(cfg.command, "\nGitlab url not provided in devctlconfig\n ")
	}

	cfg.user = parser.GetString("gitlab_user")
	if cfg.user == "" {
		color.Yellow("Gitlab user not specified, falling back to github user")
	}

}

func (cfg *cloneConfig) clone() {
	var cloneUrl string
	if cfg.user == "" && cfg.repo == "" {
		cloneUrl = cfg.url
	} else {
		cloneUrl = fmt.Sprintf("git@%s:%s/%s", cfg.url, cfg.user, cfg.repo)
	}

	session := sh.NewSession()
	session.SetDir(cfg.sourceDir)
	session.ShowCMD = true
	err := session.Command("git", "clone", cloneUrl).Run()
	if err != nil {
		fmt.Print(err)
	} else {
		color.Green("Clone was successful")
	}
}
