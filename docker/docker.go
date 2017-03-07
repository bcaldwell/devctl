package client

import (
	"errors"
	"time"

	"github.com/benjamincaldwell/devctl/shell"
	systemDetector "github.com/benjamincaldwell/go-system-detector"
	"github.com/fsouza/go-dockerclient"
)

var (
	client                *docker.Client
	ErrDockerTimeout      = errors.New("Timed out waiting for docker to start")
	ErrCouldntStartDocker = errors.New("Couldn't start docker daemon. Try starting it manually")
)

// Connect connects to the docker client. Also ensures docker is running
func Connect() (err error) {
	if client == nil {
		endpoint := "unix:///var/run/docker.sock"
		client, err = docker.NewClient(endpoint)
	}

	// if !IsDockerRunning() {
	// 	err = StartDocker()
	// }
	return err
}

func IsDockerRunning() bool {
	if client == nil {
		return false
	}
	err := client.Ping()
	return err == nil
}

func StartDocker() error {
	// timeout set to 20 seconds
	timeout := 20.0
	startTime := time.Now()
	system, _ := systemDetector.DetectSystem()
	if !IsDockerRunning() {
		var command string
		if system.ID == "darwin" {
			command = "open /Applications/Docker.app"
		} else if system.ID == "ubuntu" {
			command = "sudo start docker"
		}
		err := shell.Command("sh", "-c", command).Run()
		if err != nil {
			return err
		}
		// sudo start docker ubuntu
		// sudo systemctl start docker CentOS/Red Hat Enterprise Linux/Fedora
		for !IsDockerRunning() {
			if time.Since(startTime).Seconds() > timeout {
				return ErrDockerTimeout
			}
			time.Sleep(100)
		}
		return nil
	}
	return nil
}
