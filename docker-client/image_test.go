package dockerClient

import (
	"context"
	"testing"

	"github.com/docker/docker/client"
)

func TestCLI_PullImage(t *testing.T) {
	type fields struct {
		Client *client.Client
		ctx    context.Context
	}
	type args struct {
		image string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CLI{
				Client: tt.fields.Client,
				ctx:    tt.fields.ctx,
			}
			if err := c.PullImage(tt.args.image); (err != nil) != tt.wantErr {
				t.Errorf("CLI.PullImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCLI_IsImagePulled(t *testing.T) {
	type args struct {
		image string
	}
	tests := []struct {
		name       string
		args       args
		wantStatus bool
		wantErr    bool
	}{
		{
			name:       "nginx",
			args:       args{"nginx"},
			wantStatus: true,
			wantErr:    false,
		},
		{
			name:       "something that doesnt exist",
			args:       args{"abobajsjhdsbafhjdgkbdfsgg"},
			wantStatus: false,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		c := New()
		c.Connect()
		c.PullImage("nginx")
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, err := c.IsImagePulled(tt.args.image)
			if (err != nil) != tt.wantErr {
				t.Errorf("CLI.IsImagePulled() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("CLI.IsImagePulled() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
