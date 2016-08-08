package plugins

import (
	"io/ioutil"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/printer"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/utilities"
	"gopkg.in/yaml.v2"
)

func init() {
	AddPlugin(DockerCompose{})
}

type DockerCompose struct {
}

func (d DockerCompose) Setup() {
}

func (d DockerCompose) PreInstall(c *parser.ConfigurationStruct) {
	composerText, err := yaml.Marshal(c.DockerCompose)
	err = ioutil.WriteFile("./.devctl/docker-compose.yml", composerText, 0644)
	utilities.ErrorCheck(err, "writing docker compose file")
}

func (d DockerCompose) Install(c *parser.ConfigurationStruct) {
}

func (d DockerCompose) PostInstall(c *parser.ConfigurationStruct) {
	printer.Info("Starting docker compose")
	printer.InfoLineTop()
	shell.Command("docker-compose", "-f", "./.devctl/docker-compose.yml", "up", "-d").PrintOutput()
	printer.InfoLineBottom()
}

func (d DockerCompose) Scripts(c *parser.ConfigurationStruct) map[string]utilities.RunCommand {
	scripts := make(map[string]utilities.RunCommand)

	return scripts
}

func (d DockerCompose) IsProjectType(c *parser.ConfigurationStruct) bool {
	if c.DockerCompose != "" {
		return true
	}
	return false
}
