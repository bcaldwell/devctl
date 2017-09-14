package cmd

import (
	"strings"

	"github.com/benjamincaldwell/devctl/postCommand"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open various pages on remote host (github)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		repoCmd.Run(cmd, args)
	},
}

var repoCmd = &cobra.Command{
	Use:     "repo",
	Aliases: []string{"github", "gitlab", "gh"},
	Short:   "Open repo webpage in default browser",
	Run: func(cmd *cobra.Command, args []string) {
		postCommand.RunCommand("open " + gitURL())
	},
}

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Open new issue",
	Run: func(cmd *cobra.Command, args []string) {
		postCommand.RunCommand("open " + gitURL() + "/issues/new")
	},
}

var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Open projects issue page",
	Run: func(cmd *cobra.Command, args []string) {
		postCommand.RunCommand("open " + gitURL() + "/issues")
	},
}

// https://github.com/benjamincaldwell/devctl/compare/master...feature/add-open-command

var createPRCmd = &cobra.Command{
	Use:   "pr",
	Short: "Open new pull request",
	Run: func(cmd *cobra.Command, args []string) {
		if currentBranch, err := gitOutput("rev-parse", "--abbrev-ref", "HEAD"); err == nil {
			postCommand.RunCommand("open " + gitURL() + "/pull/" + string(currentBranch))
		}
	},
}

var PRsCmd = &cobra.Command{
	Use:   "prs",
	Short: "Open projects pull requests page",
	Run: func(cmd *cobra.Command, args []string) {
		postCommand.RunCommand("open " + gitURL() + "/pulls")
	},
}

func gitURL() string {
	if output, err := gitOutput("remote", "-v"); err == nil {
		urlLine := utilities.LineWith(string(output), "fetch")
		url := strings.Fields(urlLine)[1]
		url = strings.Replace(url, ":", "/", -1)
		url = strings.Replace(url, "git@", "http://", -1)
		url = strings.Replace(url, ".git", "", -1)
		return url
	}
	return ""
}

func gitOutput(args ...string) ([]byte, error) {
	return shell.Command("git", args...).Output()
}

func init() {
	devctlCmd.AddCommand(openCmd)
	openCmd.AddCommand(repoCmd)
	openCmd.AddCommand(issueCmd)
	openCmd.AddCommand(issuesCmd)
	openCmd.AddCommand(createPRCmd)
	openCmd.AddCommand(PRsCmd)
}
