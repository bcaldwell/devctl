package languages

import (
	"strings"

	"github.com/benjamincaldwell/devctl/parser"
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

}

func (n Node) PreInstall(c *parser.ConfigurationStruct) {
	printer.Info("setting node version to " + c.Node.Version)
	n.version = c.Node.Version

	// check if nvm is install
	_, err := sh.Command("sh", "-c", "source ~/.nvm/nvm.sh && command -v nvm").Output()
	if err != nil {
		printer.Fail("nvm not installed. Install nvm and try again")
		return
	}

	// check if requested version is installed
	nodeVersions, err := sh.Command("sh", "-c", "source ~/.nvm/nvm.sh && nvm version "+n.version).Output()
	if strings.Contains(string(nodeVersions), "N/A") {
		sh.Command("sh", "-c", "source ~/.nvm/nvm.sh && nvm install "+n.version).Output()
	}

	// set correct version in ENV
	post := new(utilities.PostCommand)
	post.RunCommand("nvm use " + n.version)
	post.Write()
}

func (n Node) Install(c *parser.ConfigurationStruct) {
	// npm install
	printer.Info("npm install")
	sh.Command("npm", "install").Output()

}

// func (n Node) scripts(c parser.ConfigurationStruct) {
// 	// return scripts struct array
// }
