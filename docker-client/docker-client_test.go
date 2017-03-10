package dockerClient

import "testing"

func TestConnect(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "connects to docker",
			wantErr: false,
		},
		{
			name:    "connects to docker again doesnt reconnect",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_IsDockerRunning(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDockerRunning(); got != tt.want {
				t.Errorf("isDockerRunning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StartDocker(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartDocker(); (err != nil) != tt.wantErr {
				t.Errorf("startDocker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
