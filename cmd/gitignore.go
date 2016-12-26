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
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/benjamincaldwell/devctl/printer"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/renstrom/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"
)

// gitignoreCmd represents the gitignore command
var gitignoreCmd = &cobra.Command{
	Aliases: []string{"i"},
	Use:     "gitignore",
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		const fileName = ".gitignore"
		if len(args) == 0 {
			utilities.ErrorWithHelp(cmd, "Minimum of one argument is required\nUse "+printer.Blue+"devctl generate gitignore list"+printer.Nc+" to list all posibilities\n")
		}
		language := strings.ToLower(args[0])

		// get list of gitignore
		printer.Info("Fetching gitignore list from github")
		resp, err := http.Get("http://api.github.com/gitignore/templates")
		defer resp.Body.Close()
		if utilities.HandleError(err) {
			return
		}
		availableTemplates := new([]string)
		err = json.NewDecoder(resp.Body).Decode(availableTemplates)
		if utilities.HandleError(err) {
			return
		}

		if language == "List" || language == "L" {
			listTemplates(*availableTemplates)
		}

		// fuzzy find messes up with caps so remove them
		translation := map[string]string{}
		for index, tlp := range *availableTemplates {
			lower := strings.ToLower(tlp)
			translation[lower] = tlp
			(*availableTemplates)[index] = lower
		}

		// fuzzy find closes result
		fuzzyFind := fuzzy.RankFind(language, *availableTemplates)
		sort.Sort(fuzzyFind)

		if len(fuzzyFind) == 0 {
			printer.Error("Unable to find template for " + language)
			listTemplates(*availableTemplates)
		} else {
			templateLanguage := translation[fuzzyFind[0].Source]
			printer.Info("Fetching template for " + templateLanguage)

			resp, err := http.Get("http://api.github.com/gitignore/templates/" + templateLanguage)
			defer resp.Body.Close()
			if utilities.HandleError(err) {
				return
			}

			gitignore := new(Gitignore)
			err = json.NewDecoder(resp.Body).Decode(gitignore)
			if utilities.HandleError(err) {
				return
			}

			if _, err := os.Stat(fileName); os.IsNotExist(err) {
				printer.Info("Creating .gitignore file")
				f, err := os.Create(fileName)
				if utilities.HandleError(err) {
					return
				}
				f.Close()
			} else {
				printer.Info("Merging with existing .gitignore file")
			}

			var fileData []byte
			fileData, err = ioutil.ReadFile(fileName)
			if utilities.HandleError(err) {
				return
			}
			writeString := utilities.UniqueStringMerge(string(fileData), *(gitignore.Source))
			err = ioutil.WriteFile(fileName, []byte(writeString), 0644)
		}
	},
}

func init() {
	generateCmd.AddCommand(gitignoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gitignoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitignoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// Gitignore represents a .gitignore file as returned by the GitHub API.
type Gitignore struct {
	Name   *string `json:"name,omitempty"`
	Source *string `json:"source,omitempty"`
}

func listTemplates(availableTemplates []string) {
	printer.InfoLineTop()
	for _, template := range availableTemplates {
		printer.InfoBar(template)
	}
	printer.InfoLineBottom()
	return
}
