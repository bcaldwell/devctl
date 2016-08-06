package printer

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

const nc = "\x1b[0m"
const green = "\x1b[32m"
const red = "\x1b[31m"
const blue = "\x1b[94m"
const yellow = "\x1b[33m"
const bold = "\033[1m"
const nobold = "\033[0m"

const line = "━"
const bar = "┃ "
const cornerTop = "┏"
const cornerBottom = "┗"

func Success(text string, a ...interface{}) {
	fmt.Printf(green+"✔ "+nc+text+"\n", a...)
}

func Fail(text string, a ...interface{}) {
	fmt.Printf(red+"✗ "+nc+text+"\n", a...)
}

func Error(text string, a ...interface{}) {
	fmt.Printf(red+"✗ "+nc+text+"\n", a...)
}

func Info(text string, a ...interface{}) {
	fmt.Printf(blue+"🐧  "+nc+text+"\n", a...)
}

func Warning(text string, a ...interface{}) {
	fmt.Printf(yellow+"⚠ "+nc+text+"\n", a...)
}

func SuccessBar(text string, a ...interface{}) {
	fmt.Printf(green+bar+nc+text+"\n", a...)
}

func ErrorBar(text string, a ...interface{}) {
	fmt.Printf(red+bar+nc+text+"\n", a...)
}

func InfoBar(text string, a ...interface{}) {
	fmt.Printf(blue+bar+nc+text+"\n", a...)
}

func WarningBar(text string, a ...interface{}) {
	fmt.Printf(yellow+bar+nc+text+"\n", a...)
}

func SuccessLine() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(green + strings.Repeat(line, width) + nc)
}

func ErrorLine() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(red + strings.Repeat(line, width) + nc)
}

func InfoLine() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(blue + strings.Repeat(line, width) + nc)
}

func WarningLine() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(yellow + strings.Repeat(line, width) + nc)
}

func SuccessLineTop() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(bold + green + cornerTop + strings.Repeat(line, width-1) + nc + nobold)
}

func ErrorLineTop() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(bold + red + cornerTop + strings.Repeat(line, width-1) + nc + nobold)
}

func InfoLineTop() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(bold + blue + cornerTop + strings.Repeat(line, width-1) + nc + nobold)
}

func WarningLineTop() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(bold + yellow + cornerTop + strings.Repeat(line, width-1) + nc + nobold)
}

func SuccessLineBottom() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(bold + green + cornerBottom + strings.Repeat(line, width-1) + nc + nobold)
}

func ErrorLineBottom() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(bold + red + cornerBottom + strings.Repeat(line, width-1) + nc + nobold)
}

func InfoLineBottom() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(bold + blue + cornerBottom + strings.Repeat(line, width-1) + nc + nobold)
}

func WarningLineBottom() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(bold + yellow + cornerBottom + strings.Repeat(line, width-1) + nc + nobold)
}

// GetSize returns the dimensions of the given terminal.
// https://github.com/golang/crypto/blob/master/ssh/terminal/util.go#L80
func getSize(fd int) (width, height int, err error) {
	var dimensions [4]uint16

	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions)), 0, 0, 0); err != 0 {
		return -1, -1, err
	}
	return int(dimensions[1]), int(dimensions[0]), nil
}
