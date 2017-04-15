package dockerClient

import (
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerImage interface {
	PullImage(image string) (err error)
	IsImagePulled(image string) (status bool, err error)
	RemoveImage(image string) (err error)
}

func (c *CLI) PullImage(image string) (err error) {
	if exists, _ := c.IsImagePulled(image); exists {
		return nil
	}

	resp, err := c.Client.ImagePull(c.ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer resp.Close()
	_, err = ioutil.ReadAll(resp)
	return err
}

func (c *CLI) IsImagePulled(image string) (status bool, err error) {
	_, _, err = c.Client.ImageInspectWithRaw(c.ctx, image)
	return !client.IsErrImageNotFound(err), err
}

func (c *CLI) RemoveImage(image string) (err error) {
	if exists, _ := c.IsImagePulled(image); !exists {
		return nil
	}

	_, err = c.Client.ImageRemove(c.ctx, image, types.ImageRemoveOptions{})
	return err
}
