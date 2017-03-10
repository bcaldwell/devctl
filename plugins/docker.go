package plugins

import (
	"fmt"

	"github.com/benjamincaldwell/devctl/docker-client"
	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/utilities"
)

type Docker struct {
}

// install docker task
type installDocker struct {
	client dockerClient.Client
}

func (i *installDocker) String() string {
	return "Installing docker"
}

func (i *installDocker) Check() (bool, error) {
	return utilities.CheckIfInstalled("docker"), nil
}

func (i *installDocker) Execute() error {
	// TODO make this not expect brew to work
	return shell.Command("brew", "cask", "install", "docker").PrintOutput()
}

// Start docker task
type startDocker struct {
	client dockerClient.Client
}

func (s *startDocker) String() string {
	return "Starting and configuring Docker"
}

func (s *startDocker) Check() (bool, error) {
	err := s.client.Connect()
	if err != nil {
		return false, err
	}
	return s.client.IsDockerRunning(), nil
}

func (s *startDocker) Execute() error {
	return s.client.StartDocker()
}

// Docker network tasks
type createNetwork struct {
	client dockerClient.Client
	name   string
}

func (c *createNetwork) String() string {
	return fmt.Sprintf("Creating %s docker network", c.name)
}

func (c *createNetwork) Check() (bool, error) {
	network, err := c.client.NetworkByName(c.name)
	if err != nil {
		return false, err
	} else if network.ID != "" {
		return true, err
	}
	return false, err
}

func (c *createNetwork) Execute() error {
	_, err := c.client.CreateNetwork(c.name)
	return err
}

// Docker plugin struct
func (d Docker) String() string {
	return "Docker setup"
}

func (d Docker) Setup() {

}

func (d Docker) UpTasks(*parser.ConfigurationStruct) (tasks []Task, err error) {
	client := dockerClient.New()
	tasks = []Task{
		&installDocker{client},
		&startDocker{client},
		&createNetwork{client, "traefik-devctl"},
		&createNetwork{client, "project-name"},
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
