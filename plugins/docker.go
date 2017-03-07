package plugins

import (
	"github.com/benjamincaldwell/devctl/docker"
	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/utilities"
)

type Docker struct{}

// install docker task
type installDocker struct{}

func (i *installDocker) String() string {
	return "Installing docker"
}

func (i *installDocker) Check() (bool, error) {
	return utilities.CheckIfInstalled("docker"), nil
}

func (i *installDocker) Execute() error {
	return nil
}

// Start docker task
type startDocker struct{}

func (s *startDocker) String() string {
	return "Starting Docker"
}

func (s *startDocker) Check() (bool, error) {
	client.Connect()
	return client.IsDockerRunning(), nil
}

func (s *startDocker) Execute() error {
	return client.StartDocker()
}

func (d Docker) String() string {
	return "Docker setup"
}

func (d Docker) Setup() {

}

func (d Docker) UpTasks(*parser.ConfigurationStruct) (tasks []Task, err error) {
	tasks = []Task{
		&installDocker{},
		&startDocker{},
	}
	return tasks, err
}

func (d Docker) PreInstall(c *parser.ConfigurationStruct) {

}

func (n Docker) Install(c *parser.ConfigurationStruct) {
}

func (n Docker) PostInstall(c *parser.ConfigurationStruct) {
}

func (d Docker) PreScript(c *parser.ConfigurationStruct) {
}

func (d Docker) Scripts(c *parser.ConfigurationStruct) (scripts map[string]utilities.RunCommand) {
	return scripts
}

func (d Docker) PostScript(c *parser.ConfigurationStruct) {
}

func (d Docker) Down(c *parser.ConfigurationStruct) {
}

func (n Docker) IsProjectType(c *parser.ConfigurationStruct) bool {
	return true
}
