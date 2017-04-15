package dockerClient

import (
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

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
