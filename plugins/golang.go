// glide
// godep
// https://github.com/FiloSottile/gvt
// https://github.com/kovetskiy/manul

package plugins

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/postCommand"
	"github.com/benjamincaldwell/devctl/printer"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/utilities"
)

type Golang struct {
}

func init() {
	AddPlugin(&Golang{})
}

func (g Golang) Setup() {
	// install go?
	// instal glide? godeps?
}

func (g Golang) PreInstall(c *parser.ConfigurationStruct) {
	sourceDir := parser.GetString("source_dir")
	// set go path
	printer.Info("adding devctl root to gopath")
	if gopath, ok := os.LookupEnv("GOPATH"); !ok || utilities.StringInSlice(sourceDir, strings.Split(gopath, ":")) {
		printer.Success("gopath already set properly")
	} else {
		postCommand.RunCommand("export GOPATH=${GOPATH}:" + sourceDir)
		printer.Success("adding devctl root to gopath")
	}

	if _, err := os.Stat("glide.yaml"); err == nil && !utilities.CheckIfInstalled("glide") {
		resp, err := http.Get("https://glide.sh/get")
		utilities.ErrorCheck(err, "glide download")
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		printer.Info("Installing glide")
		err = shell.Command("sh").SetEnv("GOPATH", "/home/benjamin/go").SetInput(string(body)).Run()
		utilities.ErrorCheck(err, "glide install")
	} else if _, err := os.Stat("Godeps/Godeps.json"); err == nil && !utilities.CheckIfInstalled("godep") {
		printer.Info("Installing godep")
		err := shell.Command("go", "get", "github.com/tools/godep").Run()
		utilities.ErrorCheck(err, "godep install")
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
		shell.Command("godep", "restore", "-v").PrintOutput()
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

func (g Golang) PreScript(c *parser.ConfigurationStruct) {
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

func (g Golang) PostScript(c *parser.ConfigurationStruct) {
}

func (g Golang) Down(c *parser.ConfigurationStruct) {
}

func (n Golang) IsProjectType(c *parser.ConfigurationStruct) bool {
	return c.Go.Version != ""
}
