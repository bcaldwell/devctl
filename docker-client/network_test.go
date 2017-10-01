package dockerClient

import (
	"reflect"
	"testing"

	"github.com/docker/docker/api/types"
)

func TestCLI_CreateNetwork(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		c       *CLI
		args    args
		wantID  string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, err := tt.c.CreateNetwork(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CLI.CreateNetwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotID != tt.wantID {
				t.Errorf("CLI.CreateNetwork() = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}

func TestCLI_Network(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name        string
		c           *CLI
		args        args
		wantNetwork types.NetworkResource
		wantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNetwork, err := tt.c.NetworkByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CLI.Network() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNetwork, tt.wantNetwork) {
				t.Errorf("CLI.Network() = %v, want %v", gotNetwork, tt.wantNetwork)
			}
		})
	}
}

func TestCLI_NetworkByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		c           *CLI
		args        args
		wantNetwork types.NetworkResource
		wantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNetwork, err := tt.c.NetworkByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CLI.NetworkByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNetwork, tt.wantNetwork) {
				t.Errorf("CLI.NetworkByName() = %v, want %v", gotNetwork, tt.wantNetwork)
			}
		})
	}
}
