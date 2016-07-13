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

	"github.com/codeskyblue/go-sh"
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

func init() {
	devctlCmd.AddCommand(cloneCmd)
	//cloneCmd.Flags().StringVarP(&gitUser, "gitUser", "u", "", "Non default git user")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func clone(cmd *cobra.Command, args []string) {
	repo, user := parseArgs(args)
	if user == "" {
		usrStr, _ := viper.Get("github_user").(string)
		// Handle Err
		user = usrStr
	}

	sourceDir := viper.Get("source_dir")

	gitRepo := fmt.Sprintf("git@github.com:%s/%s", user, repo)
	session := sh.NewSession()
	session.SetDir(sourceDir.(string))
	fmt.Printf("Cloning %s/%s to %s\n", user, repo, sourceDir)
	session.ShowCMD = true
	out, err := session.Command("git", "clone", gitRepo).Output()
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(out)
	}
}

func parseArgs(args []string) (string, string) {
	repo := ""
	user := ""
	if len(args) == 0 {
		fmt.Println("Minimum of one argument is required")
		os.Exit(-1)
	} else if len(args) == 1 {
		repo = args[0]
	} else if len(args) == 2 {
		user = args[0]
		repo = args[1]
	}
	return repo, user
}
