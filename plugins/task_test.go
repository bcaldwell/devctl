package plugins

import (
	"errors"
	"fmt"
	"testing"
)

type testTask struct {
	count     int
	trueCount int
	errAfter  int
	name      string
}

func (t testTask) Name() string {
	return t.name
}

func (t testTask) Check() (bool, error) {
	var err error
	if t.count >= t.errAfter {
		err = errors.New("my error")
	}
	return (t.count >= t.trueCount), err
}

func (t *testTask) Execute() error {
	t.count++
	return nil
}

type executeErrorTask struct {
	testTask
}

func (e *executeErrorTask) Execute() error {
	return errors.New("Execution error")
}

var alwaysTrue = &testTask{
	0, 0, 1000, "always true",
}
var alwaysFalse = &testTask{
	0, 1000, 1000, "always false",
}
var returnsFalseError = &testTask{
	0, 1000, 0, "returns error",
}
var returnsTrueError = &testTask{
	0, 0, 0, "returns error",
}

func TestRunTasks(t *testing.T) {
	type args struct {
		tasks []Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "all checks pass",
			args: args{
				[]Task{
					alwaysTrue, alwaysTrue,
				},
			},
			wantErr: false,
		},
		{
			name: "all checks pass",
			args: args{
				[]Task{
					alwaysFalse, alwaysFalse,
				},
			},
			wantErr: true,
		},
		{
			name: "one check needs to be run once",
			args: args{
				[]Task{
					alwaysTrue,
					&testTask{
						0, 1, 100, "run once",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "execute returns error",
			args: args{
				[]Task{
					alwaysTrue,
					&executeErrorTask{
						testTask{
							0, 100, 100, "execution error",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "check that returns error",
			args: args{
				[]Task{
					alwaysTrue,
					returnsTrueError,
				},
			},
			wantErr: true,
		},
		{
			name: "check that returns error after task is executed",
			args: args{
				[]Task{
					alwaysTrue,
					&testTask{0, 100, 1, "returns error after execution"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RunTasks(tt.args.tasks); (err != nil) != tt.wantErr {
				t.Errorf("RunTasks() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println()
		})
	}
}

func TestRunChecks(t *testing.T) {
	type args struct {
		tasks []Task
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "checks all return true",
			args: args{
				[]Task{
					alwaysTrue, alwaysTrue,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "checks all return false",
			args: args{
				[]Task{
					alwaysFalse, alwaysFalse,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "mix of true and false",
			args: args{
				[]Task{alwaysFalse, alwaysTrue},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "error returned truthy",
			args: args{
				[]Task{alwaysTrue, returnsTrueError},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "error returned falsy",
			args: args{
				[]Task{alwaysTrue, returnsFalseError},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RunChecks(tt.args.tasks)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunChecks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RunChecks() = %v, want %v", got, tt.want)
			}
		})
	}
}
