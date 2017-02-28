package client

import (
	"errors"
	"time"

	"github.com/benjamincaldwell/devctl/shell"
	printer "github.com/benjamincaldwell/go-printer"
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

	if !isDockerRunning() {
		err = startDocker()
	}
	return err
}

func isDockerRunning() bool {
	err := client.Ping()
	return err == nil
}

func startDocker() error {
	// timeout set to 20 seconds
	timeout := 20.0
	startTime := time.Now()
	if !isDockerRunning() {
		printer.Info("Starting docker")
		err := shell.Command("open", "/Applications/Docker.app").Run()
		if err != nil {
			return err
		}
		// sudo start docker ubuntu
		// sudo systemctl start docker CentOS/Red Hat Enterprise Linux/Fedora
		for !isDockerRunning() {
			if time.Since(startTime).Seconds() > timeout {
				return ErrDockerTimeout
			}
			time.Sleep(100)
		}
		printer.Success("Docker started successfully")
		return nil
	}
	return nil
}
