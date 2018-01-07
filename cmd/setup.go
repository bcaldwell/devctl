package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/bcaldwell/devctl/postCommand"
	printer "github.com/bcaldwell/go-printer"

	"github.com/bcaldwell/devctl/plugins"
	"github.com/bcaldwell/devctl/utilities"
	"github.com/spf13/cobra"
)

//go:generate go-bindata -prefix "../" -o bindata.go -pkg cmd ../devctl.sh

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: setup,
}

func init() {
	devctlCmd.AddCommand(setupCmd)
}

func setup(cmd *cobra.Command, args []string) {
	printer.InfoLineTop()
	// create devctl.sh file in devctl home folder (HOME/.devctl)
	data, err := Asset("devctl.sh")
	utilities.Check(err, "Fetching devctl.sh file contents")

	fileName := path.Join(devctlHomeFolder, "devctl.sh")
	utilities.Check(err, "Creating file "+fileName)
	f, err := os.Create(fileName)
	defer f.Close()

	_, err = f.Write(data)
	utilities.Check(err, "Writing contents to "+fileName)

	profileFile := detectProfile()
	profileFile = path.Join(os.Getenv("HOME"), profileFile)

	devctlSourceString := fmt.Sprintf("[ -f %s ] && \\. %s # This loads devctl shell super powers", fileName, fileName)

	fileData, err := ioutil.ReadFile(profileFile)
	if utilities.HandleError(err) {
		return
	}
	writeString := utilities.UniqueStringMerge(string(fileData), devctlSourceString)
	err = ioutil.WriteFile(profileFile, []byte(writeString), 0644)

	postCommand.RunCommand("source " + fileName)

	printer.InfoBar(printer.ColoredString("{{green:%s}} Setup shell functions"), printer.SuccessIcon)

	for _, i := range plugins.List {
		i.Setup()
	}
	printer.InfoLineBottom()
}

func detectProfile() string {
	return ".zshrc"
}
