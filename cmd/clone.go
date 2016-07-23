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
	"path"
	"strings"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/printer"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/codeskyblue/go-sh"
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
var tag string

func init() {
	devctlCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().BoolVarP(&gitlab, "gitlab", "l", false, "Clone from gitlab url")
	cloneCmd.Flags().StringVarP(&tag, "tag", "t", "", "Clone from gitlab url")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

type cloneConfig struct {
	Repo      string
	User      string
	Host      string
	Url       string
	SourceDir string
	command   *cobra.Command
}

func clone(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		utilities.ErrorWithHelp(cmd, "\nMinimum of one argument is required\n ")
	}

	cfg := new(cloneConfig)
	cfg.command = cmd

	cfg.github()
	if gitlab {
		cfg.gitlab()
	}
	cfg.parseArgs(args)

	cfg.setSourceDir()
	// validate(cfg)

	cfg.clone()
}

func (cfg *cloneConfig) parseArgs(args []string) {
	if isFullURL(args[0]) {
		cfg.parseFullURL(args[0])
		return
	}

	args = strings.Split(args[0], "/")

	if len(args) == 1 {
		cfg.Repo = args[0]
	} else if len(args) == 2 {
		cfg.User = args[0]
		cfg.Repo = args[1]
	}
}

func isFullURL(s string) bool {
	if strings.Contains(s, ":") {
		return true
	}
	return false
}

func (cfg *cloneConfig) parseFullURL(url string) {
	// if ssh
	parts := strings.Split(url, ":")
	if strings.HasPrefix(url, "git@") {
		// parse hostname from git@github.com
		cfg.Host = strings.Split(parts[0], "git@")[1]
		// remove .git
		repoString := strings.Split(parts[1], ".git")[0]
		// parse username/repo
		parts = strings.Split(repoString, "/")
		cfg.User = parts[0]
		cfg.Repo = parts[1]
	} else if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		parts = strings.Split(url, "/")
		cfg.Host = parts[2]
		cfg.User = parts[3]
		// remove .git
		cfg.Repo = strings.Split(parts[4], ".git")[0]
	} else {
		printer.Fail("couldnt parse git url")
	}
	cfg.Url = url
}

func (cfg *cloneConfig) github() {
	cfg.Url = "github.com"
	cfg.Host = "github.com"
	user := parser.GetString("github_user")
	if user != "" {
		cfg.User = user
	}
}

func (cfg *cloneConfig) gitlab() {
	cfg.Url = parser.GetString("gitlab_url")
	cfg.Host = parser.GetString("gitlab_url")
	if cfg.Url == "" {
		utilities.ErrorWithHelp(cfg.command, "\nGitlab url not provided in devctlconfig\n ")
	}

	cfg.User = parser.GetString("gitlab_user")
	if cfg.User == "" {
		printer.Warning("Gitlab user not specified, falling back to github user")
	}

}

func (cfg *cloneConfig) setSourceDir() {
	sourceDir := parser.GetString("source_dir")
	cfg.SourceDir = path.Join(sourceDir, "src", cfg.Host, cfg.User, tag)
	os.MkdirAll(cfg.SourceDir, 0755)
}

func (cfg *cloneConfig) clone() {
	var cloneURL string
	if isFullURL(cfg.Url) {
		cloneURL = cfg.Url
	} else {
		cloneURL = fmt.Sprintf("git@%s:%s/%s", cfg.Url, cfg.User, cfg.Repo)
	}

	session := sh.NewSession()
	session.SetDir(cfg.SourceDir)
	session.ShowCMD = true
	err := session.Command("git", "clone", cloneURL).Run()
	if err != nil {
		fmt.Print(err)
	} else {
		printer.Success("Clone was successful")
	}
}
