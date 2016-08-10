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

func init() {
	AddPlugin(Golang{})
}

func (g Golang) Setup() {
	// install go?
	// instal glide? godeps?
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
	printer.Info("getting go dependencies")
	if _, err := os.Stat("glide.yaml"); err == nil {
		// glide
		printer.Info("using glide")
		printer.InfoLineTop()
		shell.Command("glide", "install").PrintOutput()
	} else if _, err := os.Stat("Godeps/Godeps.json"); err == nil {
		// godep
		printer.Info("using godep")
		printer.InfoLineTop()
		shell.Command("godep", "restore").PrintOutput()
	} else {
		// go get
		printer.Info("using go get")
		printer.InfoLineTop()
		shell.Command("sh", "-c", "go get $(go list ./... | grep -v /vendor/)").PrintOutput()
	}
	printer.InfoLineBottom()
}

func (n Golang) PostInstall(c *parser.ConfigurationStruct) {
}

func (n Golang) Scripts(c *parser.ConfigurationStruct) map[string]utilities.RunCommand {
	// return scripts struct array
	scripts := make(map[string]utilities.RunCommand)

	scripts["test"] = utilities.RunCommand{
		Desc:    "run go test on non vendor packages",
		Command: "go test -v $(go list ./... | grep -v /vendor/)",
	}

	scripts["run"] = utilities.RunCommand{
		Desc:    "",
		Command: "go run *.go",
	}

	scripts["build"] = utilities.RunCommand{
		Desc:    "",
		Command: "go build *.go",
	}

	scripts["install"] = utilities.RunCommand{
		Desc:    "",
		Command: "go install *.go",
	}

	return scripts
}

func (n Golang) IsProjectType(c *parser.ConfigurationStruct) bool {
	return c.Go.Version != ""
}
