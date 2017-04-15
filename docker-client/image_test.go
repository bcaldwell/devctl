package dockerClient

import (
	"testing"
)

func TestCLI_PullImage(t *testing.T) {
	type args struct {
		image string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "nginx",
			args:    args{"nginx"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		c := New()
		c.Connect()
		c.RemoveImage("nginx")
		t.Run(tt.name, func(t *testing.T) {
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
