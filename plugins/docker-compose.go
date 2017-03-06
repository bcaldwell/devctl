package plugins

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/utilities"
	"gopkg.in/yaml.v2"
	"github.com/benjamincaldwell/go-printer"
)

func init() {
	AddPlugin(&DockerCompose{})
}

var composeFile string

type DockerCompose struct {
}

func (d DockerCompose) Setup() {
}

func (d *DockerCompose) PreInstall(c *parser.ConfigurationStruct) {
	// create .devctl folder
	os.Mkdir(".devctl", 0700)

	// make sure docker is running
	if !isDockerRunning() {
		printer.Info("Starting docker")
		shell.Command("open", "/Applications/Docker.app").Run()
		for !isDockerRunning() {
			time.Sleep(100)
		}
		printer.Success("Docker started successfully")
	}

	if c.DockerComposeFile != "" {
		composeFile = c.DockerComposeFile
	} else {
		// create docker-compose.yml file
		composerText, err := yaml.Marshal(c.DockerCompose)
		err = ioutil.WriteFile("./.devctl/docker-compose.yml", composerText, 0644)
		utilities.ErrorCheck(err, "writing docker compose file")
		composeFile = "./.devctl/docker-compose.yml"
	}

}

func (d *DockerCompose) Install(c *parser.ConfigurationStruct) {
}

func (d *DockerCompose) PostInstall(c *parser.ConfigurationStruct) {
	printer.Info("Starting docker compose")
	printer.InfoLineTop()
	err := shell.Command("docker-compose", "-f", composeFile, "up", "-d").PrintOutput()
	printer.InfoLineBottom()
	utilities.ErrorCheck(err, "starting docker compose")
	if err != nil {
		printer.Info("Try running docker-compose -f %s up to debug", composeFile)
	}
}

func (d *DockerCompose) PreScript(c *parser.ConfigurationStruct) {
}

func (d *DockerCompose) Scripts(c *parser.ConfigurationStruct) map[string]utilities.RunCommand {
	scripts := make(map[string]utilities.RunCommand)

	return scripts
}

func (d *DockerCompose) PostScript(c *parser.ConfigurationStruct) {
}

func (d *DockerCompose) Down(c *parser.ConfigurationStruct) {
}

func (d *DockerCompose) IsProjectType(c *parser.ConfigurationStruct) bool {
	return c.DockerCompose != nil || c.DockerComposeFile != ""
}
