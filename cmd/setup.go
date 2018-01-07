package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"runtime"

	printer "github.com/bcaldwell/go-printer"

	"github.com/bcaldwell/devctl/plugins"
	"github.com/bcaldwell/devctl/postCommand"
	"github.com/bcaldwell/devctl/utilities"
	"github.com/spf13/cobra"
)

//go:generate go-bindata -prefix "../" -o bindata.go -pkg cmd ../devctl.sh

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up devctl dependencies and install shell extensions",
	Long:  ``,
	Run:   setup,
}

func init() {
	devctlCmd.AddCommand(setupCmd)
}

func setup(cmd *cobra.Command, args []string) {
	printer.InfoLineTop()

	setupShellScript()

	for _, i := range plugins.List {
		i.Setup()
	}

	printer.InfoLineBottom()
}

func setupShellScript() error {
	// create devctl.sh file in devctl home folder (HOME/.devctl)
	data, err := Asset("devctl.sh")
	utilities.Check(err, "Fetching devctl.sh file contents")

	fileName := path.Join(devctlHomeFolder, "devctl.sh")
	if utilities.Check(err, "Creating file "+fileName) {
		return err
	}
	f, err := os.Create(fileName)
	defer f.Close()

	_, err = f.Write(data)
	if utilities.Check(err, "Writing contents to "+fileName) {
		return err
	}

	profileFile := detectProfile()
	fmt.Println(profileFile)
	profileFile = path.Join(os.Getenv("HOME"), profileFile)

	devctlSourceString := fmt.Sprintf("[ -f %s ] && \\. %s # This loads devctl shell super powers", fileName, fileName)

	fileData, err := ioutil.ReadFile(profileFile)
	if utilities.HandleError(err) {
		return err
	}
	writeString := utilities.UniqueStringMerge(string(fileData), devctlSourceString)
	err = ioutil.WriteFile(profileFile, []byte(writeString), 0644)

	postCommand.RunCommand("source " + fileName)

	printer.InfoBar(printer.ColoredString("{{green:%s}} Setup shell functions"), printer.SuccessIcon)
	return nil
}

func detectProfile() string {
	var re = regexp.MustCompile(`-.*\z`)
	shell := os.Getenv("SHELL")
	shell = path.Base(shell)
	// handle possible version suffix like `zsh-5.2`
	shell = re.ReplaceAllString(shell, "")

	switch shell {
	case "zsh":
		return ".zshrc"
	case "bash":
		if runtime.GOOS == "darwin" {
			return ".bash_profile"
		}
		return ".bashrc"
	case "csh":
		return ".cshrc"
	case "fish":
		return ".config/fish/config.fish"
	case "ksh":
		return ".kshrc"
	case "sh":
		return ".bash_profile"
	case "tcsh":
		return ".tcshrc"
	}
	return ".bash_profile"
}
