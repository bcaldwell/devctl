package dockerClient

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

// A default client the is already connected to Docker
var DefaultClient = New()

func Connect() error {
	return DefaultClient.Connect()
}

func IsDockerRunning() bool {
	return DefaultClient.IsDockerRunning()
}

func StartDocker() error {
	return DefaultClient.StartDocker()
}

// Network
func NetworkByName(s string) (types.NetworkResource, error) {
	return DefaultClient.NetworkByName(s)
}

func NetworkByID(s string) (types.NetworkResource, error) {
	return DefaultClient.NetworkByID(s)
}

func CreateNetwork(s string) (id string, err error) {
	return DefaultClient.CreateNetwork(s)
}

// Image
func PullImage(image string, pull bool) (err error) {
	return DefaultClient.PullImage(image, pull)
}

func IsImagePulled(image string) (status bool, err error) {
	return DefaultClient.IsImagePulled(image)
}

func RemoveImage(image string) (err error) {
	return DefaultClient.RemoveImage(image)
}

func Images() ([]types.ImageSummary, error) {
	return DefaultClient.Images()
}

// RunContainer creates and starts a container with given options
func RunContainer(config *container.Config, hostConfig *container.HostConfig) (container.ContainerCreateCreatedBody, error) {
	return DefaultClient.RunContainer(config, hostConfig)
}

func ContainerExec(id string, command []string) (resp types.ContainerExecInspect, err error) {
	return DefaultClient.ContainerExec(id, command)
}
