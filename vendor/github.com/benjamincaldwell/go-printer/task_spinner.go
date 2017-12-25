package printer

import (
	"fmt"
	"sync"
	"time"
)

var spinStates = []string{"⠋", "⠙", "⠚", "⠞", "⠖", "⠦", "⠴", "⠲", "⠳", "⠓"}
var taskSpinners []*TaskSpinner
var spinnersMutex = new(sync.Mutex)

const (
	running = iota
	success
	failure
)

func NewTaskSpinner(name string) *TaskSpinner {
	ch := make(chan string)
	t := TaskSpinner{
		Name:       name,
		Ch:         ch,
		tick:       0,
		State:      running,
		SpinStates: spinStates,
	}
	t.init()
	return &t
}

type TaskSpinner struct {
	Prefix     string
	State      int
	FinalMSG   string
	MSG        string
	Name       string
	Ch         chan string
	tick       int
	m          sync.Mutex
	SpinStates []string
	wg         *sync.WaitGroup
}

func (t *TaskSpinner) init() {
	go func(t *TaskSpinner) {
		for {
			msg, more := <-t.Ch
			t.m.Lock()

			//false when channel is closed
			if !more {
				if t.State == running {
					t.State = success
				}
				t.m.Unlock()
				return
			}
			t.MSG = msg
			t.m.Unlock()
		}
	}(t)
}

// Fail set the task to the failure state. This will also stop the spinner and close its channel.
func (t *TaskSpinner) Fail() {
	t.m.Lock()
	defer t.m.Unlock()

	close(t.Ch)
	t.State = failure
}

// Success set the task to the succesful state. This will also stop the spinner and close its channel.
func (t *TaskSpinner) Success() {
	t.m.Lock()
	defer t.m.Unlock()

	close(t.Ch)
	t.State = success
}

func (t *TaskSpinner) draw() {
	t.m.Lock()
	defer t.m.Unlock()

	spinStatesLen := len(t.SpinStates)

	spin := t.SpinStates[t.tick%spinStatesLen]
	t.tick++

	spinColor := 35
	if (t.tick % spinStatesLen) < spinStatesLen/2 {
		spinColor = 32
	}

	msg := t.Name
	if t.MSG != "" {
		msg += " [" + t.MSG + "]"
	}

	if t.State != running && t.FinalMSG != "" {
		msg = t.FinalMSG
	}

	switch t.State {
	case running:
		fmt.Printf("\r%s\x1b[%dm%s\x1b[0m %s\x1b[K\n", t.Prefix, spinColor, spin, msg)
	case success:
		fmt.Printf("%s\x1b[K\x1b[32m✓\x1b[0m %s\x1b[1B\r", t.Prefix, msg)
	case failure:
		fmt.Printf("%s\x1b[K\x1b[31m✗\x1b[0m %s\x1b[1B\r", t.Prefix, msg)
	}

	// todo think of better way todo this, because this is shit
	if t.State != running && t.wg != nil {
		t.wg.Done()
		t.wg = nil
	}
}

func AddSpinner(name string) *TaskSpinner {
	spinnersMutex.Lock()
	defer spinnersMutex.Unlock()

	t := NewTaskSpinner(name)
	taskSpinners = append(taskSpinners, t)
	return t
}

func StartSpinners() (wg *sync.WaitGroup) {
	// avoid array being modified while spinners are running
	spinnersMutex.Lock()

	taskCount := len(taskSpinners)
	wg = &sync.WaitGroup{}
	wg.Add(taskCount)

	if taskCount == 0 {
		return wg
	}

	for i := range taskSpinners {
		taskSpinners[i].wg = wg
	}

	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		defer ticker.Stop()
		defer spinnersMutex.Unlock()
		drawSpinners()
		for range ticker.C {
			if len(taskSpinners) == 0 {
				ticker.Stop()
				continue
			}
			fmt.Printf("\x1b[%dF", taskCount)
			drawSpinners()
		}
	}()

	return wg
}

func ResetSpinners() {
	// spinnersMutex.Lock()
	// defer spinnersMutex.Unlock()
	for _, spinner := range taskSpinners {
		if spinner.State == running {
			spinner.Success()
		}
	}
	taskSpinners = []*TaskSpinner{}
}

func WaitAllSpinners() {
	if len(taskSpinners) == 0 {
		return
	}
	wg := StartSpinners()
	wg.Wait()
	fmt.Printf("\x1b[%dF", len(taskSpinners))
	drawSpinners()
}

func drawSpinners() {
	for _, spinner := range taskSpinners {
		spinner.draw()
	}
}
