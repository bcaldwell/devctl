package cmd

import (
	"strings"

	"github.com/bcaldwell/devctl/postCommand"
	"github.com/bcaldwell/devctl/shell"
	"github.com/bcaldwell/devctl/utilities"
	"github.com/spf13/cobra"
)

var force bool

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Aliases: []string{"g"},
	Use:     "git",
	Short:   "a set of git alaises",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		postCommand.RunCommand("git " + strings.Join(args, " "))
	},
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "removes local branches which where deleted remotely",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if force {
			postCommand.RunCommand("sh -c \"git fetch -p && git branch -vv | awk '/: gone]/{print $1}' | xargs git branch -D\"")
		} else {
			postCommand.RunCommand("sh -c \"git fetch -p && git branch -vv | awk '/: gone]/{print $1}' | xargs git branch -d\"")
		}
	},
}

var yoloCmd = &cobra.Command{
	Use:   "yolo",
	Short: "removes local branches which where deleted remotely",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := shell.Command("git", "add", "-A").PrintOutput()
		if !utilities.HandleError(err) {
			err = shell.Command("git", "commit", "--amend", "--no-edit").PrintOutput()
			if !utilities.HandleError(err) {
				if currentBranch, err := gitOutput("rev-parse", "--abbrev-ref", "HEAD"); err == nil {
					// for some reason its not working in go...
					postCommand.RunCommand("git push -f origin " + string(currentBranch))
					// err = shell.Command("git", "push", "-f", "-v", "origin", string(currentBranch)+":"+string(currentBranch)).PrintOutput()
					// utilities.HandleError(err)
				}
			}
		}
	},
}

//

func init() {
	devctlCmd.AddCommand(gitCmd)
	gitCmd.AddCommand(cleanCmd)
	gitCmd.AddCommand(yoloCmd)
	cleanCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "force")
}
