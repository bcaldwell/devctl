// glide
// godep
// https://github.com/FiloSottile/gvt
// https://github.com/kovetskiy/manul

package plugins

import (
	"os"
	"strings"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/post_command"
	"github.com/benjamincaldwell/devctl/printer"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/utilities"
)

type Golang struct {
}

func (g Golang) Setup() {
	// install go?
}

func (g Golang) PreInstall(c *parser.ConfigurationStruct) {
	sourceDir := parser.GetString("source_dir")
	// set go path
	printer.Info("adding devctl root to gopath")
	if gopath, ok := os.LookupEnv("GOPATH"); !ok || strings.Contains(gopath, sourceDir) {
		printer.Success("gopath already set properly")
	} else {
		postCommand.RunCommand("export GOPATH=${GOPATH}:" + sourceDir)
		printer.Success("adding devctl root to gopath")
	}
}

func (n Golang) Install(c *parser.ConfigurationStruct) {
	printer.Info("go getting dependencies")
	printer.InfoLineTop()
	shell.Command("sh", "-c", "go get $(go list ./... | grep -v /vendor/)").PrintOutput()
	printer.InfoLineBottom()
}

func (n Golang) Scripts(c *parser.ConfigurationStruct) map[string]utilities.RunCommand {
	// return scripts struct array
	scripts := make(map[string]utilities.RunCommand)

	scripts["test"] = utilities.RunCommand{
		Desc:    "run go test on non vendor packages",
		Command: "go test $(go list ./... | grep -v /vendor/)",
	}

	scripts["run"] = utilities.RunCommand{
		Desc:    "",
		Command: "go run *.go",
	}

	scripts["build"] = utilities.RunCommand{
		Desc:    "",
		Command: "go build *.go",
	}

	return scripts
}

func (n Golang) IsProjectType(c *parser.ConfigurationStruct) bool {
	if c.Go.Version != "" {
		return true
	}
	return false
}
