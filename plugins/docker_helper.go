package plugins

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"syscall"

	"github.com/benjamincaldwell/devctl/docker-client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Start docker container task
type startContainerTask struct {
	client    dockerClient.Client
	image     string
	Container container.ContainerCreateCreatedBody
}

func (s *startContainerTask) String() string {
	return fmt.Sprintf("Starting %s", s.image)
}

func (s *startContainerTask) Check() (bool, error) {
	return false, nil
}

func (s *startContainerTask) Execute() error {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePull(ctx, s.image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()
	// Required to wait for container pull to finish
	io.Copy(ioutil.Discard, out)

	s.Container, err = cli.ContainerCreate(ctx, &container.Config{
		Image: s.image,
		Cmd:   []string{"tail", "-f", "/dev/null"},
		Volumes: map[string]struct{}{
			pwd: {},
		},
		WorkingDir: pwd,
	}, &container.HostConfig{
		Binds: []string{
			fmt.Sprintf("%s:%s", pwd, pwd),
		},
	}, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, s.Container.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	out_2, err := cli.ContainerExecCreate(ctx, s.Container.ID, types.ExecConfig{
		Cmd: []string{"bash -c 'curl -o- -L https://yarnpkg.com/install.sh | bash'"},
	})

	if err != nil {
		panic(err)
	}

	err = cli.ContainerExecStart(ctx, out_2.ID, types.ExecStartCheck{})

	if err != nil {
		panic(err)
	}

	out_2, err = cli.ContainerExecCreate(ctx, s.Container.ID, types.ExecConfig{
		Cmd: []string{"afdg", "install", "--silent"},
		// seems like a hack. Not sure what is going on
		AttachStdout: true,
		AttachStderr: true,
	})

	if err != nil {
		panic(err)
	}

	err = cli.ContainerExecStart(ctx, out_2.ID, types.ExecStartCheck{})

	if err != nil {
		panic(err)
	}

	fmt.Print("\n\n")
	// syscall.Exec("/usr/local/bin/docker", []string{"docker", "exec", "-it", s.Container.ID, "/bin/bash"}, []string{})
	syscall.Exec("/usr/local/bin/docker", []string{"docker", "exec", "-it", s.Container.ID, "yarn", "start"}, []string{})
	return nil
}
