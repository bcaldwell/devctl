package plugins

import (
	"time"

	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/go-printer"
)

func isDockerRunning() bool {
	err := shell.Command("docker", "info").Run()
	return err == nil
}

func startDocker() bool {
	if !isDockerRunning() {
		printer.Info("Starting docker")
		shell.Command("open", "/Applications/Docker.app").Run()
		for !isDockerRunning() {
			time.Sleep(100)
		}
		printer.Success("Docker started successfully")
		return true
	}
	return true
}
