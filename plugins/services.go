package plugins

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/post_command"
	"github.com/benjamincaldwell/devctl/printer"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/fsouza/go-dockerclient"
)

func init() {
	AddPlugin(&Docker{
		services: make(map[string]string),
	})
}

type Service struct {
	image       string
	port        string
	volumes     []string
	environment []string
}

type Docker struct {
	services map[string]string
	client   *docker.Client
}

func (d Docker) Setup() {
}

func (d *Docker) PreInstall(c *parser.ConfigurationStruct) {
	var err error
	// make sure docker is running
	startDocker()
	endpoint := "unix:///var/run/docker.sock"
	d.client, err = docker.NewClient(endpoint)
	if err != nil {
		printer.Fail("Failed to connect to docker")
		return
	}

}

func (d *Docker) Install(c *parser.ConfigurationStruct) {
	opts := docker.CreateContainerOptions{
		Config: &docker.Config{
			Image: "mysql",
		},
		HostConfig: &docker.HostConfig{
			// Binds: []string{
			// 	"/Users/benjamincaldwell/src/github.com/benjamincaldwell/devctl/testing_dir/src/github.com/devctl/mean-example/.devctl/services/mongo:/data/db",
			// },
			PortBindings: map[docker.Port][]docker.PortBinding{
				"27017/tcp": []docker.PortBinding{},
			},
			PublishAllPorts: true,
		},
		// Name: "mongo"}
	}

	container, _ := d.client.CreateContainer(opts)

	err := d.client.StartContainer(
		container.ID,
		&docker.HostConfig{},
	)

	fmt.Println(err)

	container, _ = d.client.InspectContainer(container.ID)
	fmt.Printf("\n\n%+v", container.NetworkSettings.Ports)
}

func (d *Docker) PostInstall(c *parser.ConfigurationStruct) {

}

func (d *Docker) PreScript(c *parser.ConfigurationStruct) {
	printer.Info("Setting up services environment")
	env, err := os.Open(".devctl/env")
	if err != nil {
		printer.Fail("Couldn't read env file. Trying running devctl up")
		return
	}
	defer env.Close()

	// match blah=blah not #blah=blah
	r, err := regexp.Compile(`^([^\s#]+)=([^\s]+)$`)
	if utilities.HandleError(err) {
		return
	}

	startDocker()

	scanner := bufio.NewScanner(env)
	for scanner.Scan() {
		if res := r.FindAllStringSubmatch(scanner.Text(), 2); res != nil {
			envName := res[0][1]
			url := res[0][2]
			d.services[envName] = url
			exportString := fmt.Sprintf("export %s=%s", envName, url)
			postCommand.RunCommand(exportString)
		}
	}

	if err := scanner.Err(); utilities.HandleError(err) {
		return
	}
}

func (d *Docker) Scripts(c *parser.ConfigurationStruct) map[string]utilities.RunCommand {
	scripts := make(map[string]utilities.RunCommand)

	scripts["env"] = utilities.RunCommand{
		Desc:    "",
		Command: "printenv | grep URL",
	}
	return scripts
}

func (d *Docker) PostScript(c *parser.ConfigurationStruct) {
	for envName, _ := range d.services {
		unsetString := fmt.Sprintf("unset %s", envName)
		postCommand.RunCommand(unsetString)
	}
}

func (d *Docker) IsProjectType(c *parser.ConfigurationStruct) bool {
	return len(c.Services) > 0
}
