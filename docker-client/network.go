package dockerClient

import "github.com/docker/docker/api/types"

type DockerNetwork interface {
	Network(id string) (network types.NetworkResource, err error)
	NetworkByName(name string) (network types.NetworkResource, err error)
	CreateNetwork(name string) (id string, err error)
}

func (c *CLI) CreateNetwork(name string) (id string, err error) {
	var network types.NetworkCreateResponse
	network, err = c.Client.NetworkCreate(c.ctx, name, types.NetworkCreate{})

	return network.ID, err
}

func (c *CLI) Network(id string) (network types.NetworkResource, err error) {
	return c.Client.NetworkInspect(c.ctx, id)
}

func (c *CLI) NetworkByName(name string) (network types.NetworkResource, err error) {
	// TODO: make this use docker args
	networks, err := c.Client.NetworkList(c.ctx, types.NetworkListOptions{})
	if err == nil {
		for _, network := range networks {
			if network.Name == name {
				return network, err
			}
		}
	}
	return network, err
}
