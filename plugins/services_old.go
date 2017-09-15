package plugins

// import docker "github.com/fsouza/go-dockerclient"

// import docker "github.com/fsouza/go-dockerclient"

// import (
// 	"bufio"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path"
// 	"regexp"
// 	"runtime"
// 	"strings"
// 	"sync"

// 	"crypto/sha256"

// 	"github.com/benjamincaldwell/devctl/parser"
// 	"github.com/benjamincaldwell/devctl/postCommand"
// 	"github.com/benjamincaldwell/devctl/shell"
// 	"github.com/benjamincaldwell/devctl/utilities"
// 	"github.com/benjamincaldwell/go-printer"
// )

// // func init() {
// // AddPlugin(&Docker{
// // 	ids:         make(map[string]string),
// // 	tags:        make(map[string]string),
// // 	environment: make(map[string]string),
// // })
// // }

type Service struct {
	image string
	port  string
	// port        docker.Port
	volumes     []string
	environment []string
	tag         string
}

// type Docker struct {
// 	ids         map[string]string
// 	tags        map[string]string
// 	environment map[string]string
// 	client      *docker.Client
// }

// func (d Docker) Setup() {
// 	if runtime.GOOS == "darwin" {
// 		if !utilities.CheckIfInstalled("docker") {
// 			printer.Info("Installing docker for mac")
// 			tmpfile, err := ioutil.TempFile("", "docker-install-")
// 			printer.VerboseInfo("created temporary file %s", tmpfile.Name())
// 			defer os.Remove(tmpfile.Name()) // clean up

// 			if utilities.HandleError(err, "creating a temporary file") {
// 				return
// 			}

// 			printer.VerboseInfo("Downloading docker from stable channel")
// 			data, err := utilities.HTTPDownload("https://download.docker.com/mac/stable/Docker.dmg")
// 			if utilities.HandleError(err, "downloading docker") {
// 				return
// 			}

// 			localSha := fmt.Sprintf("%x", sha256.Sum256(data))
// 			printer.VerboseInfo("local docker sha: %s", localSha)

// 			var remoteShaBytes []byte
// 			remoteShaBytes, err = utilities.HTTPDownload("https://download.docker.com/mac/stable/Docker.dmg.sha256sum")
// 			if utilities.HandleError(err, "downloading docker sha") {
// 				return
// 			}

// 			remoteSha := strings.Fields(string(remoteShaBytes))[0]
// 			printer.VerboseInfo("remote docker sha: %s", localSha)

// 			if remoteSha != localSha {
// 				printer.Fail("sha verification of docker.dmg failed")
// 				return
// 			}

// 			_, err = tmpfile.Write(data)
// 			if utilities.HandleError(err, "writing docker.dmg to temporary file") {
// 				return
// 			}
// 			err = tmpfile.Close()
// 			if utilities.HandleError(err, "closing temporary file") {
// 				return
// 			}

// 			printer.VerboseInfo("mounting docker")
// 			err = shell.Command("hdiutil", "mount", tmpfile.Name()).Run()
// 			if utilities.HandleError(err, "mounting dokcer") {
// 				return
// 			}

// 			mountPoint := "/Volumes/Docker"
// 			printer.Info("Installing docker to /Applications")
// 			err = shell.Command("sudo", "cp", "-R", mountPoint+"/Docker.app", "/Applications").Run()
// 			if utilities.HandleError(err, "copying docker to /Applications") {
// 				return
// 			}

// 			printer.VerboseInfo("unmounting docker")
// 			err = shell.Command("hdiutil", "unmount", mountPoint).Run()
// 			if utilities.HandleError(err, "unmounting docker") {
// 				return
// 			}
// 			printer.Success("Successfully install docker for mac")
// 		} else {
// 			printer.Success("docker for mac already installed")
// 		}
// 	}
// }

// func (d *Docker) PreInstall(c *parser.ProjectConfigStruct) {
// 	var err error
// 	// make sure docker is running
// 	startDocker()
// 	endpoint := "unix:///var/run/docker.sock"
// 	d.client, err = docker.NewClient(endpoint)
// 	if err != nil {
// 		printer.Fail("Failed to connect to docker")
// 		return
// 	}
// }

// func (d *Docker) Install(c *parser.ProjectConfigStruct) {

// 	printer.Info("Starting services")
// 	printer.InfoLineTop()

// 	dir, err := os.Getwd()
// 	if utilities.HandleError(err) {
// 		return
// 	}

// 	var wg sync.WaitGroup

// 	d.ids, err = parser.ReadTomlLike(".devctl/service_list")

// 	for _, serviceConf := range c.Services {
// 		var tag string
// 		var serviceName string

// 		// since serviceConf is being parsed as an interface it can be either a string of map[interface{}]interface{}
// 		// note: servce:version with no space is parsed as a string
// 		switch s := serviceConf.(type) {
// 		case string:
// 			parts := strings.Split(s, ":")
// 			serviceName = parts[0]
// 			defaultTag := serviceList[serviceName].tag
// 			tag = defaultTag
// 			if len(parts) > 1 {
// 				tag = parts[1]
// 				if defaultTag == "alpine" {
// 					tag += "-alpine"
// 				}
// 			}
// 		case map[interface{}]interface{}:
// 			for key, value := range s {
// 				serviceName = key.(string)
// 				tag = fmt.Sprint(value)
// 				break
// 			}
// 		default:
// 			printer.Fail("Invaild service configuration %v", s)
// 			return
// 		}

// 		d.tags[serviceName] = tag
// 		wg.Add(1)
// 		go startService(&wg, d, serviceName, dir)
// 	}

// 	wg.Wait()

// 	printer.InfoLineBottom()

// 	parser.WriteMapTomlLike(d.ids, ".devctl/service_list")
// 	parser.WriteMapTomlLike(d.environment, ".devctl/env")
// }

// func (d *Docker) PostInstall(c *parser.ProjectConfigStruct) {
// 	printer.Info("Running service health check")
// 	printer.InfoLineTop()

// 	for service, id := range d.ids {
// 		if status, err := containerStatus(d, id); err != nil || status != "running" {
// 			printer.ErrorBar("%s is not running", service)
// 		} else {
// 			printer.SuccessBar("%s is running", service)
// 		}
// 	}

// 	printer.InfoLineBottom()
// }

// func (d *Docker) PreScript(c *parser.ProjectConfigStruct) {
// 	printer.Info("Setting up services environment")
// 	env, err := os.Open(".devctl/env")
// 	if err != nil {
// 		printer.Fail("Couldn't read env file. Trying running devctl up")
// 		return
// 	}
// 	defer env.Close()

// 	// match blah=blah not #blah=blah
// 	r, err := regexp.Compile(`^([^\s#]+)=([^\s]+)$`)
// 	if utilities.HandleError(err) {
// 		return
// 	}

// 	startDocker()

// 	scanner := bufio.NewScanner(env)
// 	for scanner.Scan() {
// 		if res := r.FindAllStringSubmatch(scanner.Text(), 2); res != nil {
// 			envName := res[0][1]
// 			url := res[0][2]
// 			d.environment[envName] = url
// 			exportString := fmt.Sprintf("export %s=%s", envName, url)
// 			postCommand.RunCommand(exportString)
// 		}
// 	}

// 	if err := scanner.Err(); utilities.HandleError(err) {
// 		return
// 	}
// }

// func (d *Docker) Scripts(c *parser.ProjectConfigStruct) map[string]utilities.RunCommand {
// 	scripts := make(map[string]utilities.RunCommand)

// 	scripts["env"] = utilities.RunCommand{
// 		Desc:    "",
// 		Command: "printenv | grep URL",
// 	}
// 	return scripts
// }

// func (d *Docker) PostScript(c *parser.ProjectConfigStruct) {
// 	for envName := range d.environment {
// 		unsetString := fmt.Sprintf("unset %s", envName)
// 		postCommand.RunCommand(unsetString)
// 	}
// }

// func (d *Docker) Down(c *parser.ProjectConfigStruct) {
// 	var err error
// 	endpoint := "unix:///var/run/docker.sock"
// 	d.client, err = docker.NewClient(endpoint)
// 	if err != nil {
// 		printer.Fail("Failed to connect to docker")
// 		return
// 	}

// 	printer.Info("Stop services")
// 	printer.InfoLineTop()

// 	var wg sync.WaitGroup

// 	d.ids, err = parser.ReadTomlLike(".devctl/service_list")
// 	for name, id := range d.ids {
// 		go stopContainer(&wg, d, name, id)
// 		wg.Add(1)
// 	}

// 	wg.Wait()
// 	printer.InfoLineBottom()
// }

// func (d *Docker) IsProjectType(c *parser.ProjectConfigStruct) bool {
// 	return len(c.Services) > 0
// }

// func createContainer(d *Docker, serviceConfig Service, tag, dir string) (string, error) {

// 	containerName := fmt.Sprintf("devctl-%s-%s-%s", path.Base(dir), serviceConfig.image, utilities.RandStringBytesRmndr(5))

// 	opts := docker.CreateContainerOptions{
// 		Config: &docker.Config{
// 			Image:  serviceConfig.image + ":" + tag,
// 			Env:    serviceConfig.environment,
// 			Labels: map[string]string{"devctl": path.Base(dir)},
// 		},
// 		HostConfig: &docker.HostConfig{
// 			Binds:           []string{},
// 			PublishAllPorts: true,
// 			// RestartPolicy:   docker.RestartUnlessStopped(),
// 		},
// 		Name: containerName,
// 	}

// 	for _, volume := range serviceConfig.volumes {
// 		localDir := path.Join(dir, ".devctl/services", serviceConfig.image, volume)
// 		opts.HostConfig.Binds = append(opts.HostConfig.Binds, localDir+":"+volume)
// 	}

// 	container, err := d.client.CreateContainer(opts)
// 	if err != nil {
// 		if err == docker.ErrNoSuchImage {
// 			printer.InfoBar("Getting image for %s", serviceConfig.image)
// 			if err = d.client.PullImage(docker.PullImageOptions{Repository: serviceConfig.image, Tag: tag}, docker.AuthConfiguration{}); err != nil {
// 				printer.ErrorBar("Failed to get image for %s because %s", serviceConfig.image, err)
// 				return "", err
// 			}
// 			if container, err = d.client.CreateContainer(opts); err != nil {
// 				return "", err
// 			}
// 			return container.ID, nil
// 		}
// 		return "", err
// 	}

// 	return container.ID, nil
// }

// func containerStatus(d *Docker, id string) (string, error) {
// 	container, err := d.client.InspectContainer(id)
// 	if err != nil {
// 		if _, ok := err.(*docker.NoSuchContainer); ok {
// 			return "deleted", nil
// 		}
// 		return "", err
// 	}
// 	return container.State.Status, nil
// }

// func startService(wg *sync.WaitGroup, d *Docker, serviceName string, dir string) {
// 	defer wg.Done()

// 	var err error
// 	if val, ok := d.ids[serviceName]; ok {
// 		status, err := containerStatus(d, val)
// 		utilities.HandleError(err)
// 		if status == "running" {
// 			printer.SuccessBar("%s already running", serviceName)
// 			d.environment[strings.ToUpper(serviceName)+"_URL"] = "localhost:" + getDockerPort(d, serviceName)
// 			return
// 		} else if status == "deleted" {
// 			printer.InfoBar("creating service %s", serviceName)
// 			d.ids[serviceName], err = createContainer(d, serviceList[serviceName], d.tags[serviceName], dir)
// 			if err != nil {
// 				printer.ErrorBar("Failed to create service %s because %s", serviceName, err)
// 				return
// 			}
// 		} else {
// 			printer.InfoBar("starting service %s", serviceName)
// 		}
// 	} else {
// 		printer.InfoBar("creating service %s", serviceName)
// 		d.ids[serviceName], err = createContainer(d, serviceList[serviceName], d.tags[serviceName], dir)
// 		if err != nil {
// 			printer.ErrorBar("Failed to create service %s because %s", serviceName, err)
// 			return
// 		}
// 	}

// 	err = d.client.StartContainer(
// 		d.ids[serviceName],
// 		&docker.HostConfig{},
// 	)

// 	if err != nil {
// 		printer.ErrorBar("failed to start service %s because %s", serviceName, err)
// 	} else {
// 		printer.SuccessBar("Successfully started %s", serviceName)
// 	}

// 	d.environment[strings.ToUpper(serviceName)+"_URL"] = getDockerPort(d, serviceName)
// }

// func stopContainer(wg *sync.WaitGroup, d *Docker, name, id string) error {
// 	defer wg.Done()
// 	printer.InfoBar("Stoping %s", name)
// 	err := d.client.StopContainer(id, 60)
// 	if err != nil {
// 		printer.ErrorBar("Failed to stop %s because %s", name, err)
// 	} else {
// 		printer.SuccessBar("Successfully stopped %s", name)
// 	}
// 	return err
// }

// func getDockerPort(d *Docker, serviceName string) string {
// 	container, _ := d.client.InspectContainer(d.ids[serviceName])
// 	port := serviceList[serviceName].port
// 	if portBinding, ok := container.NetworkSettings.Ports[port]; ok {
// 		return portBinding[0].HostPort
// 	}
// 	return ""
// }
