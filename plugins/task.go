package plugins

import (
	"errors"

	"github.com/benjamincaldwell/go-printer"
)

// Task is the interface used by task runner
type Task interface {
	String() string
	Check() (bool, error)
	Execute() error
}

// CheckFailedAfterTaskErr is the error that is returned if task is run but check doesnt pass afterwards
var CheckFailedAfterTaskErr = errors.New("Check failed after tasks was executed")

// RunTasks runs all tasks. First checks the check. If check returns false the task is executed.
// The check is checked again after the task is executed.
func RunTasks(tasks []Task) (err error) {
	var pass bool
	for _, task := range tasks {
		pass, err = task.Check()
		if err != nil {
			return err
		}
		if pass {
			printer.SuccessBarIcon("%s (Already completed)", task)
			continue
		}
		// run task
		printer.InfoBar("Executing: %s", task)
		err = task.Execute()
		if err != nil {
			return err
		}
		// check check after task was run
		pass, err = task.Check()
		if err != nil {
			return err
		}
		if !pass {
			printer.ErrorBarIcon("%s (Check failed after task was executed)", task)
			return CheckFailedAfterTaskErr
		}
	}
	return nil
}

// RunChecks returns the result of calling the check function of all the tasks passed in.
// Returns as soon as one returns false or an error.
// If an error occurs, false, error witll be returned
func RunChecks(tasks []Task) (pass bool, err error) {
	for _, task := range tasks {
		pass, err = task.Check()
		if !pass || err != nil {
			return false, err
		}
	}
	return true, nil
}
