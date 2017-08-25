package printer

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"syscall"
	"unsafe"
)

// Enables verbose printing
var Verbose = false

const line = "━"
const bar = "┃ "
const cornerTop = "┏"
const cornerBottom = "┗"

const leftPad = 2

var InfoIcon = "🐧"
var SuccessIcon = "✔"
var ErrorIcon = "✗"
var WarningIcon = "⚠"

var redRegex = regexp.MustCompile(`{{red:(.*?)}}`)
var blueRegex = regexp.MustCompile(`{{blue:(.*?)}}`)
var greenRegex = regexp.MustCompile(`{{green:(.*?)}}`)
var cyanRegex = regexp.MustCompile(`{{cyan:(.*?)}}`)
var boldRegex = regexp.MustCompile(`{{bold:(.*?)}}`)

func init() {
	if os.Getenv("VERBOSE") == "true" {
		Verbose = true
	}
}

func Success(text string, a ...interface{}) {
	fmt.Printf(Green(SuccessIcon+" ")+text+"\n", a...)
}

func Fail(text string, a ...interface{}) {
	Error(text, a...)
}

func Error(text string, a ...interface{}) {
	fmt.Printf(Red(ErrorIcon+" ")+text+"\n", a...)
}

func Info(text string, a ...interface{}) {
	fmt.Printf(Blue(InfoIcon+"  ")+text+"\n", a...)
}

func Warning(text string, a ...interface{}) {
	fmt.Printf(Yellow(WarningIcon+" ")+text+"\n", a...)
}

func SuccessBar(text string, a ...interface{}) {
	fmt.Printf(Green(bar)+text+"\n", a...)
}

func ErrorBar(text string, a ...interface{}) {
	fmt.Printf(Red(bar)+text+"\n", a...)
}

func InfoBar(text string, a ...interface{}) {
	fmt.Printf(Blue(bar)+text+"\n", a...)
}

func WarningBar(text string, a ...interface{}) {
	fmt.Printf(Yellow(bar)+text+"\n", a...)
}

func SuccessBarIcon(text string, a ...interface{}) {
	fmt.Printf(Greenf("%s%s ", bar, SuccessIcon)+text+"\n", a...)
}

func ErrorBarIcon(text string, a ...interface{}) {
	fmt.Printf(Redf("%s%s ", bar, ErrorIcon)+text+"\n", a...)
}

func InfoBarIcon(text string, a ...interface{}) {
	fmt.Printf(Bluef("%s%s ", bar, InfoIcon)+text+"\n", a...)
}

func WarningBarIcon(text string, a ...interface{}) {
	fmt.Printf(Yellowf("%s%s ", bar, WarningIcon)+text+"\n", a...)
}

func SuccessLine() {
	width := getWidth()
	fmt.Printf(Green(strings.Repeat(line, width)))
}

func ErrorLine() {
	width := getWidth()
	fmt.Printf(Red(strings.Repeat(line, width)))
}

func InfoLine() {
	width := getWidth()
	fmt.Printf(Blue(strings.Repeat(line, width)))
}

func WarningLine() {
	width := getWidth()
	fmt.Printf(Yellow(strings.Repeat(line, width)))
}

func SuccessLineText(text string, a ...interface{}) {
	text = fmt.Sprintf(text, a...)
	width := getWidth()
	width = width - 2 - leftPad - len(text)
	fmt.Printf("%s %s %s", Green(strings.Repeat(line, leftPad)), text, Green(strings.Repeat(line, width)))
}

func ErrorLineText(text string, a ...interface{}) {
	text = fmt.Sprintf(text, a...)
	width := getWidth()
	width = width - 2 - leftPad - len(text)
	fmt.Printf("%s %s %s", Red(strings.Repeat(line, leftPad)), text, Red(strings.Repeat(line, width)))
}

func InfoLineText(text string, a ...interface{}) {
	text = fmt.Sprintf(text, a...)
	width := getWidth()
	width = width - 2 - leftPad - len(text)
	fmt.Printf("%s %s %s", Blue(strings.Repeat(line, leftPad)), text, Blue(strings.Repeat(line, width)))
}

func WarningLineText(text string, a ...interface{}) {
	text = fmt.Sprintf(text, a...)
	width := getWidth()
	width = width - 2 - leftPad - len(text)
	fmt.Printf("%s %s %s", Yellow(strings.Repeat(line, leftPad)), text, Yellow(strings.Repeat(line, width)))
}

func SuccessLineTop() {
	width := getWidth()
	fmt.Printf(Bold(Green(cornerTop + strings.Repeat(line, width-1))))
}

func ErrorLineTop() {
	width := getWidth()
	fmt.Printf(Bold(Red(cornerTop + strings.Repeat(line, width-1))))
}

func InfoLineTop() {
	width := getWidth()
	fmt.Printf(Bold(Blue(cornerTop + strings.Repeat(line, width-1))))
}

func WarningLineTop() {
	width := getWidth()
	fmt.Printf(Bold(Yellow(cornerTop + strings.Repeat(line, width-1))))
}

func SuccessLineTextTop(text string, a ...interface{}) {
	text = fmt.Sprintf(text, a...)
	width := getWidth()
	width = width - 1 - 2 - leftPad - len(text)
	fmt.Printf(Bold("%s %s %s"), Green(cornerTop+strings.Repeat(line, leftPad)), text, Green(strings.Repeat(line, width)))
}

func ErrorLineTextTop(text string, a ...interface{}) {
	text = fmt.Sprintf(text, a...)
	width := getWidth()
	width = width - 1 - 2 - leftPad - len(text)
	fmt.Printf(Bold("%s %s %s"), Red(cornerTop+strings.Repeat(line, leftPad)), text, Red(strings.Repeat(line, width)))
}

func InfoLineTextTop(text string, a ...interface{}) {
	text = fmt.Sprintf(text, a...)
	width := getWidth()
	width = width - 1 - 2 - leftPad - len(text)
	fmt.Printf(Bold("%s %s %s"), Blue(cornerTop+strings.Repeat(line, leftPad)), text, Blue(strings.Repeat(line, width)))
}

func WarningLineTextTop(text string, a ...interface{}) {
	text = fmt.Sprintf(text, a...)
	width := getWidth()
	width = width - 1 - 2 - leftPad - len(text)
	fmt.Printf(Bold("%s %s %s"), Yellow(cornerTop+strings.Repeat(line, leftPad)), text, Yellow(strings.Repeat(line, width)))
}

func SuccessLineBottom() {
	width := getWidth()
	fmt.Printf(Bold(Green(cornerBottom + strings.Repeat(line, width-1))))
}

func ErrorLineBottom() {
	width := getWidth()
	fmt.Printf(Bold(Red(cornerBottom + strings.Repeat(line, width-1))))
}

func InfoLineBottom() {
	width := getWidth()
	fmt.Printf(Bold(Blue(cornerBottom + strings.Repeat(line, width-1))))
}

func WarningLineBottom() {
	width := getWidth()
	fmt.Printf(Bold(Yellow(cornerBottom + strings.Repeat(line, width-1))))
}

func VerboseSuccess(text string, a ...interface{}) {
	if Verbose {
		Success(text, a...)
	}
}

func VerboseFail(text string, a ...interface{}) {
	if Verbose {
		Fail(text, a...)
	}
}

func VerboseError(text string, a ...interface{}) {
	if Verbose {
		Error(text, a...)
	}
}

func VerboseInfo(text string, a ...interface{}) {
	if Verbose {
		Info(text, a...)
	}
}

func VerboseWarning(text string, a ...interface{}) {
	if Verbose {
		Warning(text, a...)
	}
}

func VerboseSuccessBar(text string, a ...interface{}) {
	if Verbose {
		SuccessBar(text, a...)
	}
}

func VerboseErrorBar(text string, a ...interface{}) {
	if Verbose {
		ErrorBar(text, a...)
	}
}

func VerboseInfoBar(text string, a ...interface{}) {
	if Verbose {
		InfoBar(text, a...)
	}
}

func VerboseWarningBar(text string, a ...interface{}) {
	if Verbose {
		WarningBar(text, a...)
	}
}

func PrintColoredln(text string, a ...interface{}) {
	PrintColored(text+"\n", a...)
}

func PrintColored(text string, a ...interface{}) {
	text = ColoredString(text)
	fmt.Printf(text, a...)
}

func ColoredString(text string) string {
	text = blueRegex.ReplaceAllString(text, Blue("$1"))
	text = redRegex.ReplaceAllString(text, Red("$1"))
	text = greenRegex.ReplaceAllString(text, Green("$1"))
	text = cyanRegex.ReplaceAllString(text, Cyan("$1"))
	text = boldRegex.ReplaceAllString(text, Bold("$1"))
	return text
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

func getWidth() int {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	return int(math.Max(float64(width), 20.0))
}
