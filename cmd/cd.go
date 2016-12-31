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
	"path/filepath"
	"sort"
	"strings"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/post_command"
	"github.com/benjamincaldwell/devctl/printer"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/renstrom/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"
)

// cdCmd represents the cd command
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Change to project directory",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: cd,
}

type folder struct {
	Root, Name string
}

func (f folder) Path() string {
	return path.Join(f.Root, f.Name)
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
	var dir string
	sourceDir := parser.GetString("source_dir")

	if len(args) == 0 {
		user := parser.GetString("github_user")
		dir = path.Join(sourceDir, "src/github.com", user)
	} else {
		query := args[0]

		sourceDir = path.Join(sourceDir, "src")

		files := getFolderList(sourceDir)

		dir = findMatch(query, files)
		if dir == "" {
			printer.Info("%s could not be found", query)
			user := parser.GetString("github_user")
			dir = path.Join(sourceDir, "github.com", user)
		}
	}

	postCommand.ChangeDir(dir)
}

func getFolderList(sourceDir string, params ...int) []folder {
	depth := 0
	if len(params) > 0 {
		depth = params[0]
	}

	files, _ := ioutil.ReadDir(sourceDir)

	folders := filterDir(files)

	currentFolder := folder{filepath.Dir(sourceDir), filepath.Base(sourceDir)}

	if isVersionControlled(files) || depth > 3 {
		var _folders = make([]folder, 0)
		return append(_folders, currentFolder)
	}

	depth++
	var allFolders = make([]folder, 0)
	allFolders = append(allFolders, currentFolder)
	for _, i := range folders {
		temp := getFolderList(path.Join(sourceDir, i.Name()), depth)
		allFolders = append(allFolders, temp...)
	}

	return allFolders
}

func isVersionControlled(files []os.FileInfo) bool {
	for _, i := range files {
		if i.Name() == ".git" {
			return true
		}
	}
	return false
}

func filterDir(files []os.FileInfo) []os.FileInfo {
	var filtered = make([]os.FileInfo, 0)
	for _, i := range files {
		if i.IsDir() && !strings.HasPrefix(i.Name(), ".") {
			filtered = append(filtered, i)
		}
	}
	return filtered
}

func fileInfotoFolder(files []os.FileInfo, root string) []folder {
	var folders = make([]folder, 0)
	for _, i := range files {
		folders = append(folders, folder{root, i.Name()})
	}
	return folders
}

func createDirTranslation(files []folder, length int) map[string]string {
	translation := map[string]string{}

	for _, i := range files {
		parts := strings.Split(i.Path(), "/")
		translation[strings.Join(parts[len(parts)-length:], "/")] = i.Path()
	}

	return translation
}

func findMatch(query string, files []folder) string {

	for i := 1; i <= 3; i++ {
		translation := createDirTranslation(files, i)

		folders := utilities.Keys(translation)
		fuzzyFind := fuzzy.RankFind(query, folders)
		sort.Sort(fuzzyFind)
		if len(fuzzyFind) > 0 {
			return translation[fuzzyFind[0].Target]
		}
	}
	return ""
}
