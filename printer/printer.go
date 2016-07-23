package printer

import "fmt"

const nc = "\x1b[0m"
const green = "\x1b[32m"
const red = "\x1b[31m"
const blue = "\x1b[94m"
const yellow = "\x1b[33m"

func Success(text string, a ...interface{}) {
	fmt.Printf(green+"✔ "+nc+text+"\n", a...)
}

func Fail(text string, a ...interface{}) {
	fmt.Printf(red+"✗ "+nc+text+"\n", a...)
}

func Info(text string, a ...interface{}) {
	fmt.Printf(blue+"🐧 "+nc+text+"\n", a...)
}

func Warning(text string, a ...interface{}) {
	fmt.Printf(yellow+"⚠ "+nc+text+"\n", a...)
}
