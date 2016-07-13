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
	"strings"

	"github.com/codeskyblue/go-sh"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func clone(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		errorWithHelp(cmd, "\nMinimum of one argument is required\n ")
	}

	var gitRepo string
	hostURL := "github.com"

	repo, user, url := parseArgs(args)

	if gitlab {
		hostURL = viper.Get("gitlab_url").(string)
		if hostURL == "" {
			errorWithHelp(cmd, "\nGitlab url not provided in devctlconfig\n ")
		}
		if user == "" {
			user = viper.Get("gitlab_user").(string)
		}
		if user == "" {
			color.Yellow("Gitlab user not specified, falling back to github user")
		}
	}

	if url != "" {
		gitRepo = url
	} else {
		if user == "" {
			user = viper.Get("github_user").(string)
		}
		gitRepo = fmt.Sprintf("git@%s:%s/%s", hostURL, user, repo)
	}

	sourceDir := viper.Get("source_dir")

	session := sh.NewSession()
	session.SetDir(sourceDir.(string))
	fmt.Printf("Cloning %s/%s to %s\n", user, repo, sourceDir)
	session.ShowCMD = true
	err := session.Command("git", "clone", gitRepo).Run()
	if err != nil {
		fmt.Print(err)
	} else {
		// fmt.Print(out)
		color.Green("Clone was successful")
	}
}

func parseArgs(args []string) (string, string, string) {

	if isFullURL(args[0]) {
		return "", "", args[0]
	}

	args = strings.Split(args[0], "/")

	repo := ""
	user := ""
	if len(args) == 1 {
		repo = args[0]
	} else if len(args) == 2 {
		user = args[0]
		repo = args[1]
	}
	return repo, user, ""
}

func isFullURL(s string) bool {
	if len(strings.Split(s, ":")) == 2 {
		return true
	} else if strings.HasPrefix(s, "https://") || strings.HasPrefix(s, "http://") {
		return true
	}
	return false
}

func errorWithHelp(cmd *cobra.Command, message string) {
	color.Red(message)
	cmd.Help()
	os.Exit(1)
}
