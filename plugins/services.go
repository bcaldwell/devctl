package plugins

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
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
	services    map[string]string
	environment map[string]string
	client      *docker.Client
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

	printer.Info("Starting services")
	printer.InfoLineTop()

	dir, err := os.Getwd()
	if utilities.HandleError(err) {
		return
	}

	d.services, err = readTomlLike(".devctl/service_list")

	for _, service := range c.Services {
		if val, ok := d.services[service]; ok {
			status, err := containerStatus(d, val)
			utilities.HandleError(err)
			if status == "running" {
				printer.SuccessBar("%s already running", service)
				return
			} else if status == "deleted" {
				printer.InfoBar("creating service %s", service)
				d.services[service], err = createContainer(d, serviceList[service], dir)
				if err != nil {
					printer.ErrorBar("Failed to create service %s with error %s", service, err)
					return
				}
			} else {
				printer.InfoBar("starting service %s", service)
			}
		} else {
			printer.InfoBar("creating service %s", service)
			d.services[service], err = createContainer(d, serviceList[service], dir)
			if err != nil {
				printer.ErrorBar("Failed to create service %s with error %s", service, err)
				return
			}
		}

		err = d.client.StartContainer(
			d.services[service],
			&docker.HostConfig{},
		)

		if err != nil {
			printer.ErrorBar("failed to start service %s with error %s", service, err)
		} else {
			printer.SuccessBar("Successfully started %s", service)
		}
	}
	printer.InfoLineBottom()

	writeMapTomlLike(d.services, ".devctl/service_list")

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
			d.environment[envName] = url
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
	for envName, _ := range d.environment {
		unsetString := fmt.Sprintf("unset %s", envName)
		postCommand.RunCommand(unsetString)
	}
}

func (d *Docker) IsProjectType(c *parser.ConfigurationStruct) bool {
	return len(c.Services) > 0
}

func writeMapTomlLike(m map[string]string, fileName string) error {
	r := ""

	for key, val := range m {
		r += fmt.Sprintf("%s=%s", key, val)
	}

	return ioutil.WriteFile(fileName, []byte(r), 0644)
}

func readTomlLike(fileName string) (map[string]string, error) {
	m := make(map[string]string)
	if _, err := os.Stat(fileName); err == nil {
		r, err := regexp.Compile(`^([^\s#]+)=([^\s]+)$`)
		if err != nil {
			return m, err
		}
		file, err := os.Open(fileName)
		if err != nil {
			return m, err
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if res := r.FindAllStringSubmatch(scanner.Text(), 2); res != nil {
				key := res[0][1]
				val := res[0][2]
				m[key] = val
			}
		}
	}
	return m, nil
}

func createContainer(d *Docker, serviceConfig Service, dir string) (string, error) {

	containerName := fmt.Sprintf("devctl-%s-%s-%s", path.Base(dir), serviceConfig.image, utilities.RandStringBytesRmndr(5))

	opts := docker.CreateContainerOptions{
		Config: &docker.Config{
			Image: serviceConfig.image,
			Env:   serviceConfig.environment,
		},
		HostConfig: &docker.HostConfig{
			Binds:           []string{},
			PublishAllPorts: true,
		},
		Name: containerName,
	}

	for _, volume := range serviceConfig.volumes {
		localDir := path.Join(dir, ".devctl/services", serviceConfig.image, volume)
		opts.HostConfig.Binds = append(opts.HostConfig.Binds, localDir+":"+volume)
	}

	container, err := d.client.CreateContainer(opts)
	if err != nil {
		return "", err
	}

	return container.ID, nil
}

func containerStatus(d *Docker, id string) (string, error) {
	container, err := d.client.InspectContainer(id)
	if err != nil {
		if _, ok := err.(*docker.NoSuchContainer); ok {
			return "deleted", nil
		}
		return "", err
	}
	return container.State.Status, nil
}
