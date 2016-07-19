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
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/fatih/color"
	"github.com/renstrom/fuzzysearch/fuzzy"
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
	if len(args) == 0 {
		// TODO make this go to /src/github.com/user
		utilities.ErrorWithHelp(cmd, "\nMinimum of one argument is required\n ")
	}

	query := args[0]

	sourceDir := parser.GetString("source_dir")
	sourceDir = path.Join(sourceDir, "src")

	files := getFolderList(sourceDir)

	for _, i := range files {
		fmt.Println(i)
	}

	dir := findMatch(query, files)
	if dir == "" {
		color.Yellow("%s could not be found", query)
		dir = sourceDir
	}

	post := new(utilities.PostCommand)
	post.ChangeDir(dir)
	post.Write()

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
	var filted = make([]os.FileInfo, 0)
	for _, i := range files {
		if i.IsDir() && !strings.HasPrefix(i.Name(), ".") {
			filted = append(filted, i)
		}
	}
	return filted
}

func fileInfotoFolder(files []os.FileInfo, root string) []folder {
	var folders = make([]folder, 0)
	for _, i := range files {
		folders = append(folders, folder{root, i.Name()})
	}
	return folders
}

func findMatch(query string, files []folder) string {
	var folders = make([]string, 0)
	for _, i := range files {
		if strings.Contains(strings.ToLower(i.Name), strings.ToLower(query)) {
			return i.Path()
		}
		folders = append(folders, i.Path())
		// folders = append(folders, i.Name)
	}
	temp := fuzzy.RankFind(query, folders)
	sort.Sort(temp)
	fmt.Println(temp)
	return temp[0].Target
}
