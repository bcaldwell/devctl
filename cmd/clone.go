package cmd

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/bcaldwell/devctl/parser"
	"github.com/bcaldwell/devctl/postCommand"
	"github.com/bcaldwell/devctl/shell"
	"github.com/bcaldwell/devctl/utilities"
	"github.com/bcaldwell/go-printer"
	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Aliases: []string{"c"},
	Use:     "clone",
	Short:   "Clone the github repository to the source_dir",
	Long:    ``,
	Run:     clone,
}

var gitlab bool
var tag string

func init() {
	devctlCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().BoolVarP(&gitlab, "gitlab", "l", false, "Clone from gitlab URL")
	cloneCmd.Flags().StringVarP(&tag, "tag", "t", "", "subfolder to clone repo to")
}

type cloneConfig struct {
	Repo      string
	User      string
	Host      string
	URL       string
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
	cfg.setURL()
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

func (cfg *cloneConfig) parseFullURL(URL string) {
	r, _ := regexp.Compile(`^(?:git@|https?:\/\/(?:.*?@)?)([^\/:]*)(?:\/|:)(.*?)\/(.*?)(?:.git)?$`)
	res := r.FindAllStringSubmatch(URL, -1)

	cfg.Host = res[0][1]
	cfg.User = res[0][2]
	cfg.Repo = res[0][3]
	cfg.URL = URL
}

func (cfg *cloneConfig) github() {
	cfg.URL = "github.com"
	cfg.Host = "github.com"
	user := parser.DevctlConfig.GithubUser
	if user != "" {
		cfg.User = user
	}
}

func (cfg *cloneConfig) gitlab() {
	cfg.URL = parser.DevctlConfig.GitlabURL
	cfg.Host = parser.DevctlConfig.GitlabURL
	if cfg.URL == "" {
		printer.Warning("Gitlab URL not provided, falling back to gitlab.com")
		cfg.URL = "gitlab.com"
		cfg.Host = "gitlab.com"
	}

	cfg.User = parser.DevctlConfig.GitlabUser
	if cfg.User == "" {
		printer.Warning("Gitlab user not specified, falling back to github user")
		cfg.User = parser.DevctlConfig.GithubUser
	}

}

func (cfg *cloneConfig) setSourceDir() {
	sourceDir := parser.DevctlConfig.SourceDir
	if sourceDir == "" {
		printer.Warning("Source Directory is blank, falling back to $HOME")
		sourceDir = os.Getenv("HOME")
	}
	cfg.SourceDir = path.Join(sourceDir, cfg.Host, cfg.User, tag)
	os.MkdirAll(cfg.SourceDir, 0755)
}

func (cfg *cloneConfig) setURL() {
	if !isFullURL(cfg.URL) {
		cfg.URL = fmt.Sprintf("git@%s:%s/%s", cfg.URL, cfg.User, cfg.Repo)
	}
}

func (cfg *cloneConfig) clone() {
	sourceDir := path.Join(cfg.SourceDir, cfg.Repo)
	sourceInfo, folderExists := os.Stat(sourceDir)

	if os.IsNotExist(folderExists) {
		session := shell.Session()
		session.SetDir(cfg.SourceDir)

		err := session.Command("git", "clone", cfg.URL).PrintOutput()
		if err != nil {
			printer.Fail("Failed to clone %s", cfg.URL)
			return
		}
		printer.Success("Clone was successful")
	} else if sourceInfo.IsDir() {
		printer.Success("Project already cloned")
	} else {
		printer.Warning("Destination is a file")
		return
	}
	postCommand.ChangeDir(path.Join(cfg.SourceDir, cfg.Repo))
}
