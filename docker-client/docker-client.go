package dockerClient

import (
	"context"
	"errors"
	"time"

	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/go-system-detector"

	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	ErrDockerTimeout      = errors.New("Timed out waiting for docker to start")
	ErrCouldntStartDocker = errors.New("Couldn't start docker daemon. Try starting it manually")
)

type Client interface {
	Connect() (err error)
	IsDockerRunning() bool
	StartDocker() error
	dockerClient() *client.Client

	// Network methods
	Network(id string) (network types.NetworkResource, err error)
	NetworkByName(name string) (network types.NetworkResource, err error)
	CreateNetwork(name string) (id string, err error)

	// Image methods
	PullImage(image string) (err error)
	IsImagePulled(image string) (status bool, err error)
	RemoveImage(image string) (err error)
}

type CLI struct {
	Client *client.Client
	ctx    context.Context
}

func New() Client {
	return &CLI{
		ctx: context.Background(),
	}
}

func (c *CLI) Connect() (err error) {
	if c.Client == nil {
		if os.Getenv("DOCKER_API_VERSION") == "" {
			os.Setenv("DOCKER_API_VERSION", c.DockerAPIversion())
		}
		c.Client, err = client.NewEnvClient()
	}
	return err
}

func (c *CLI) DockerAPIversion() (version string) {
	return "1.26"
}

func (c *CLI) IsDockerRunning() bool {
	if c.Client == nil {
		return false
	}
	_, err := c.Client.Ping(c.ctx)
	return err == nil
}

func (c *CLI) StartDocker() error {
	// timeout set to 20 seconds
	timeout := 20.0
	startTime := time.Now()
	system, _ := systemDetector.DetectSystem()
	if !c.IsDockerRunning() {
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
		for !c.IsDockerRunning() {
			if time.Since(startTime).Seconds() > timeout {
				return ErrDockerTimeout
			}
			time.Sleep(100)
		}
		return nil
	}
	return nil
}

func (c *CLI) dockerClient() *client.Client {
	return c.Client
}
