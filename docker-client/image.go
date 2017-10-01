package dockerClient

import (
	"io"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerImage interface {
	PullImage(image string, pull bool) (err error)
	IsImagePulled(image string) (status bool, err error)
	RemoveImage(image string) (err error)
	Images() ([]types.ImageSummary, error)
}

func (c *CLI) Images() ([]types.ImageSummary, error) {
	return c.Client.ImageList(c.Ctx, types.ImageListOptions{})
}

func (c *CLI) PullImage(image string, pull bool) (err error) {
	if exists, _ := c.IsImagePulled(image); !pull && exists {
		return nil
	}

	resp, err := c.Client.ImagePull(c.Ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer resp.Close()
	// Required to wait for container pull to finish
	io.Copy(ioutil.Discard, resp)
	return err
}

func (c *CLI) IsImagePulled(image string) (status bool, err error) {
	_, _, err = c.Client.ImageInspectWithRaw(c.Ctx, image)
	return !client.IsErrImageNotFound(err), err
}

func (c *CLI) RemoveImage(image string) (err error) {
	if exists, _ := c.IsImagePulled(image); !exists {
		return nil
	}

	_, err = c.Client.ImageRemove(c.Ctx, image, types.ImageRemoveOptions{})
	return err
}
