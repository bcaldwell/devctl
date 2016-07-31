package plugins

import (
	"strings"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/post_command"
	"github.com/benjamincaldwell/devctl/printer"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/codeskyblue/go-sh"
)

// make languages interface{}
// up makes array of languages interface

type Node struct {
	path    string
	version string
}

func (n Node) Setup() {
	isNvmInstalled := nvmInstalled()
	if isNvmInstalled {
		printer.Info("nvm already installed")
		return
	}
	printer.Info("Installing nvm")
	err := sh.Command("curl", "-o-", "https://raw.githubusercontent.com/creationix/nvm/v0.31.3/install.sh").Command("bash").Run()
	utilities.ErrorCheck(err, "nvm install")
}

func (n Node) PreInstall(c *parser.ConfigurationStruct) {
	printer.Info("setting node version to " + c.Node.Version)
	n.version = c.Node.Version

	// check if nvm is install
	isNvmInstalled := nvmInstalled()
	if !isNvmInstalled {
		printer.Fail("nvm not installed. Install nvm and try again")
		return
	}

	// check if requested version is installed
	nodeVersions, _ := sh.Command("sh", "-c", "source ~/.nvm/nvm.sh && nvm version "+n.version).Output()
	if strings.Contains(string(nodeVersions), "N/A") {
		sh.Command("sh", "-c", "source ~/.nvm/nvm.sh && nvm install "+n.version).Output()
	}

	// set correct version in ENV
	postCommand.RunCommand("nvm use " + n.version)
}

func (n Node) Install(c *parser.ConfigurationStruct) {
	// npm install
	printer.Info("npm install")
	sh.Command("npm", "install").Output()
}

func (n Node) Scripts(c *parser.ConfigurationStruct) map[string]utilities.RunCommand {
	// return scripts struct array
	scripts := make(map[string]utilities.RunCommand)

	packageJSON := parser.Parser{}
	if err := packageJSON.Parse("package.json"); err != nil {
		return scripts
	}

	packageJSONscripts := packageJSON.GetMap("scripts")

	for name := range packageJSONscripts {
		scripts[name] = utilities.RunCommand{
			Command: "npm run " + name,
		}
	}
	return scripts
}

func (n Node) IsProjectType(c *parser.ConfigurationStruct) bool {
	if c.Node.Version != "" {
		return true
	}
	return false
}

func nvmInstalled() bool {
	_, err := sh.Command("sh", "-c", "source ~/.nvm/nvm.sh && command -v nvm").Output()
	if err != nil {
		return false
	}
	return true
}
