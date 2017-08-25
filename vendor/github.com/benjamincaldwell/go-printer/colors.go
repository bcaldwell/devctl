package printer

import "fmt"

// Nc is the no color string constant
const NoColor = "\x1b[0m"
const NoboldString = "\033[0m"

const BoldString = "\033[1m"

const GreenColor = "\x1b[32m"
const RedColor = "\x1b[31m"

// Blue is the blue color string constant
const BlueColor = "\x1b[94m"
const YellowColor = "\x1b[33m"
const CyanColor = "\x1b[36m"

func Green(text string) string {
	return fmt.Sprintf("%s%s%s", GreenColor, text, NoColor)
}

func Greenf(text string, a ...interface{}) string {
	text = fmt.Sprintf(text, a...)
	return Green(text)
}

func Red(text string) string {
	return fmt.Sprintf("%s%s%s", RedColor, text, NoColor)
}

func Redf(text string, a ...interface{}) string {
	text = fmt.Sprintf(text, a...)
	return Red(text)
}

func Blue(text string) string {
	return fmt.Sprintf("%s%s%s", BlueColor, text, NoColor)
}

func Bluef(text string, a ...interface{}) string {
	text = fmt.Sprintf(text, a...)
	return Blue(text)
}

func Yellow(text string) string {
	return fmt.Sprintf("%s%s%s", YellowColor, text, NoColor)
}

func Yellowf(text string, a ...interface{}) string {
	text = fmt.Sprintf(text, a...)
	return Yellow(text)
}

func Cyan(text string) string {
	return fmt.Sprintf("%s%s%s", CyanColor, text, NoColor)
}

func Cyanf(text string, a ...interface{}) string {
	text = fmt.Sprintf(text, a...)
	return Cyan(text)
}

func Bold(text string) string {
	return fmt.Sprintf("%s%s%s", BoldString, text, NoboldString)
}

func Boldf(text string, a ...interface{}) string {
	text = fmt.Sprintf(text, a...)
	return Bold(text)
}
