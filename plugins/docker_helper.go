package plugins

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/benjamincaldwell/devctl/docker-client"
	"github.com/docker/docker/api/types/container"
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

	err = dockerClient.PullImage(s.image, true)
	if err != nil {
		panic(err)
	}

	s.Container, err = dockerClient.RunContainer(&container.Config{
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
	})

	_, err = dockerClient.ContainerExec(s.Container.ID, []string{"bash -c 'curl -o- -L https://yarnpkg.com/install.sh | bash'"})
	if err != nil {
		panic(err)
	}

	_, err = dockerClient.ContainerExec(s.Container.ID, []string{"yarn", "install", "--silent"})
	if err != nil {
		panic(err)
	}

	fmt.Print("\n")
	// syscall.Exec("/usr/local/bin/docker", []string{"docker", "exec", "-it", s.Container.ID, "/bin/bash"}, []string{})
	syscall.Exec("/usr/local/bin/docker", []string{"docker", "exec", "-it", s.Container.ID, "yarn", "start"}, []string{})
	return nil
}
