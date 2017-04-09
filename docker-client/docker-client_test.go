package dockerClient

import (
	"os"
	"testing"

	"github.com/benjamincaldwell/devctl/shell"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var firstClient Client
	t.Run("should return a CLI which impliments client", func(t *testing.T) {
		firstClient = New()
	})
	t.Run("running twice should return the same docker CLI", func(t *testing.T) {
		got := New()
		secondClient := got.dockerClient()
		if firstClient.dockerClient() != secondClient {
			t.Errorf("New did not return the same docker CLI")
		}
	})
}

func TestCLI_Connect(t *testing.T) {
	client := New()
	t.Run("Doesnt overwrite docker api env variable", func(t *testing.T) {
		os.Setenv("DOCKER_API_VERSION", "v0.0.0")
		client.Connect()
		if os.Getenv("DOCKER_API_VERSION") != "v0.0.0" {
			t.Errorf("DOCKER_API_VERSION was overwriten")
		}
	})
}

// func TestCLI_DockerAPIversion(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		c           *CLI
// 		wantVersion string
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotVersion := tt.c.DockerAPIversion(); gotVersion != tt.wantVersion {
// 				t.Errorf("CLI.DockerAPIversion() = %v, want %v", gotVersion, tt.wantVersion)
// 			}
// 		})
// 	}
// }

func TestCLI_StartDocker(t *testing.T) {
	if os.Getenv("DEVCTL_ENABLE_SLOW_TESTS") == "true" {
		client := New()
		client.Connect()
		t.Run("Starts docker if it isnt running", func(t *testing.T) {
			shell.Command("killall", "Docker").Run()
			assert.Equal(t, client.IsDockerRunning(), false)
			err := client.StartDocker()
			assert.Equal(t, err, nil)
			assert.Equal(t, client.IsDockerRunning(), true)
		})
	}
}
