package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/bcaldwell/devctl/utilities"
	"github.com/bcaldwell/go-printer"
	"github.com/renstrom/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Aliases: []string{"g"},
	Use:     "generate",
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("generate called")
	},
}

// gitignoreCmd represents the gitignore command
var gitignoreCmd = &cobra.Command{
	Aliases: []string{"i"},
	Use:     "gitignore",
	Short:   "generate gitignore file for project",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		const fileName = ".gitignore"
		if len(args) == 0 {
			printer.Error("Minimum of one argument is required\nUse " + printer.Blue("devctl generate gitignore list") + " to list all posibilities\n")
		}
		language := strings.ToLower(args[0])

		// get list of gitignore
		printer.Info("Fetching gitignore list from github")
		resp, err := http.Get("http://api.github.com/gitignore/templates")
		if utilities.HandleError(err) {
			return
		}
		defer resp.Body.Close()
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
			if utilities.HandleError(err) {
				return
			}
			defer resp.Body.Close()

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
			if utilities.HandleError(err, "Reading "+fileName) {
				return
			}
			writeString := utilities.UniqueStringMerge(string(fileData), *(gitignore.Source))
			err = ioutil.WriteFile(fileName, []byte(writeString), 0644)
			if utilities.HandleError(err, "Appending source string to "+fileName) {
				return
			}
		}
	},
}

func init() {
	devctlCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(gitignoreCmd)
}

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
