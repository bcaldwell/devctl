package dockerClient

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

type DockerContainer interface {
	RunContainer(config *container.Config, hostConfig *container.HostConfig) (container.ContainerCreateCreatedBody, error)
	ContainerExec(id string, command []string) (resp types.ContainerExecInspect, err error)
}

func (c *CLI) RunContainer(config *container.Config, hostConfig *container.HostConfig) (container container.ContainerCreateCreatedBody, err error) {
	container, err = c.Client.ContainerCreate(c.Ctx, config, hostConfig, nil, "")
	if err != nil {
		return container, err
	}

	err = c.Client.ContainerStart(c.Ctx, container.ID, types.ContainerStartOptions{})
	return container, err
}

func (c *CLI) ContainerExec(id string, command []string) (resp types.ContainerExecInspect, err error) {
	out2, err := c.Client.ContainerExecCreate(c.Ctx, id, types.ExecConfig{
		Cmd: command,
	})

	if err != nil {
		return resp, err
	}

	err = c.Client.ContainerExecStart(c.Ctx, out2.ID, types.ExecStartCheck{})

	if err != nil {
		return resp, err
	}

	for {
		resp, err = c.Client.ContainerExecInspect(c.Ctx, out2.ID)
		if err != nil {
			return resp, err
		}
		if !resp.Running {
			break
		}
	}
	return resp, err
}
