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

	"github.com/bcaldwell/devctl/parser"
	"github.com/bcaldwell/devctl/postCommand"
	"github.com/bcaldwell/devctl/shell"
	"github.com/bcaldwell/devctl/utilities"
	"github.com/bcaldwell/go-printer"
)

type Golang struct {
}

func (g Golang) String() string {
	return "Go"
}

func (g Golang) Setup() {
	// install go?
	// instal glide? godeps?
}

func (g Golang) Check(c *parser.ProjectConfigStruct) bool {
	return true
}

func (g Golang) Up(c *parser.ProjectConfigStruct) {

}

func (g Golang) PreInstall(c *parser.ProjectConfigStruct) {
	// sourceDir := parser.GetString("source_dir")
	sourceDir := ""
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
		// TODO: BAD. also make GOATPH/bin dir
		err = shell.Command("sh").SetEnv("GOPATH", "/home/benjamin/go").SetInput(string(body)).PrintOutput()
		utilities.ErrorCheck(err, "glide install")
	} else if _, err := os.Stat("Godeps/Godeps.json"); err == nil && !utilities.CheckIfInstalled("godep") {
		printer.Info("Installing godep")
		err := shell.Command("go", "get", "github.com/tools/godep").Run()
		utilities.ErrorCheck(err, "godep install")
	}
}

func (g Golang) Install(c *parser.ProjectConfigStruct) {
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

func (g Golang) PostInstall(c *parser.ProjectConfigStruct) {
}

func (g Golang) PreScript(c *parser.ProjectConfigStruct) {
}

func (g Golang) Scripts(c *parser.ProjectConfigStruct) map[string]utilities.RunCommand {
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

func (g Golang) PostScript(c *parser.ProjectConfigStruct) {
}

func (g Golang) Down(c *parser.ProjectConfigStruct) {
}

func (g Golang) IsProjectType(c *parser.ProjectConfigStruct) bool {
	return c.Go.Version != ""
}
