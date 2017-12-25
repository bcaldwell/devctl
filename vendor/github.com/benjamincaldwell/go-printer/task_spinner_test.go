package printer

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTaskSpinner(t *testing.T) {
	fmt.Println()
	s := AddSpinner("hello")
	w := AddSpinner("bye")
	s.FinalMSG = "well, I'm not done but I tried. I failed... :("
	s.Prefix = Blue(bar)
	w.Prefix = Blue(bar)
	go func() {
		time.Sleep(1 * time.Second)
		s.Ch <- "suppppppp"
		w.Ch <- "sleep"
		time.Sleep(1 * time.Second)
		s.Fail()
		time.Sleep(1 * time.Second)
		w.Success()
		time.Sleep(1 * time.Second)
	}()
	WaitAllSpinners()
}
