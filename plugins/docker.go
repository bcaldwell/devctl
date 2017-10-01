package plugins

import (
	"fmt"

	"github.com/benjamincaldwell/devctl/docker-client"
	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/utilities"
)

// install docker task
type installDocker struct {
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
}

func (s *startDocker) String() string {
	return "Starting and configuring Docker"
}

func (s *startDocker) Check() (bool, error) {
	err := dockerClient.Connect()
	if err != nil {
		return false, err
	}
	return dockerClient.IsDockerRunning(), nil
}

func (s *startDocker) Execute() error {
	return dockerClient.StartDocker()
}

// Docker network tasks
type createNetwork struct {
	name string
}

func (c *createNetwork) String() string {
	return fmt.Sprintf("Creating %s docker network", c.name)
}

func (c *createNetwork) Check() (bool, error) {
	network, err := dockerClient.NetworkByName(c.name)
	if err != nil {
		return false, err
	} else if network.ID != "" {
		return true, err
	}
	return false, err
}

func (c *createNetwork) Execute() error {
	_, err := dockerClient.CreateNetwork(c.name)
	return err
}

// Docker plugin struct
type Docker struct{}

func (d Docker) String() string {
	return "Docker setup"
}

func (d Docker) Setup() {
	RunTask(&installDocker{})
}

func (d Docker) UpTasks(config *parser.ProjectConfigStruct) (tasks [][]Task, err error) {
	stage1 := []Task{
		&installDocker{},
		&startDocker{},
		&createNetwork{"traefik-devctl"},
		&createNetwork{"project-name"},
	}
	tasks = append(tasks, stage1)

	return tasks, err
}

func (d Docker) PreInstall(c *parser.ProjectConfigStruct) {

}

func (n Docker) Install(c *parser.ProjectConfigStruct) {
}

func (n Docker) PostInstall(c *parser.ProjectConfigStruct) {
}

func (d Docker) PreScript(c *parser.ProjectConfigStruct) {
}

func (d Docker) Scripts(c *parser.ProjectConfigStruct) (scripts map[string]utilities.RunCommand) {
	return scripts
}

func (d Docker) PostScript(c *parser.ProjectConfigStruct) {
}

func (d Docker) Down(c *parser.ProjectConfigStruct) {
}

func (n Docker) IsProjectType(c *parser.ProjectConfigStruct) bool {
	return true
}
