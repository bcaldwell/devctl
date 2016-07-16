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
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// cdCmd represents the cd command
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: cd,
}

func init() {
	devctlCmd.AddCommand(cdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func cd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		utilities.ErrorWithHelp(cmd, "\nMinimum of one argument is required\n ")
	}

	query := args[0]

	sourceDir := parser.GetString("source_dir")

	files, _ := ioutil.ReadDir(sourceDir)
	files = filterDir(files)

	match := findMatch(query, files)

	dir := sourceDir

	if match == "" {
		color.Yellow("%s could not be found", query)
	}

	dir = path.Join(sourceDir, match)

	post := new(utilities.PostCommand)
	post.ChangeDir(dir)
	post.Write()

}

func filterDir(files []os.FileInfo) []os.FileInfo {
	var filted = make([]os.FileInfo, 0)
	for _, i := range files {
		if i.IsDir() {
			filted = append(filted, i)
		}
	}
	return filted
}

func findMatch(query string, files []os.FileInfo) string {
	for _, i := range files {
		if strings.Contains(strings.ToLower(i.Name()), strings.ToLower(query)) {
			return i.Name()
		}
	}
	return ""
}
