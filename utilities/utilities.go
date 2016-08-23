package utilities

import (
	"math/rand"
	"os"

	"github.com/benjamincaldwell/devctl/printer"
	"github.com/spf13/cobra"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type RunCommand struct {
	Desc    string
	Command string
}

// ErrorWithHelp Show show error message, help menu and exit
func ErrorWithHelp(cmd *cobra.Command, message string) {
	printer.Fail(message)
	cmd.Help()
	os.Exit(1)
}

func ErrorCheck(err error, message string) {
	if err == nil {
		printer.Success("%s successful", message)
	} else {
		printer.Fail("%s failed with %s", message, err)
	}
}

// Keys returns the keys of a map
func Keys(arr map[string]string) []string {
	keys := make([]string, len(arr))
	i := 0
	for k := range arr {
		keys[i] = k
		i++
	}
	return keys
}

// HandleError prints error with fail printer if it exists
func HandleError(err error) bool {
	if err != nil {
		printer.Fail("%s", err)
		return true
	}
	return false
}

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
