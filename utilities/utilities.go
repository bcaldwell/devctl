package utilities

import (
	"fmt"
	"math/rand"
	"os"

	"strings"

	"github.com/benjamincaldwell/devctl/printer"
	"github.com/benjamincaldwell/devctl/shell"
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

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func nvmInstalled() bool {
	err := shell.Command("sh", "-c", "source ~/.nvm/nvm.sh && command -v nvm").Run()
	if err != nil {
		return false
	}
	return true
}

// CheckIfInstalled checks if a binary is install. Optional argument ti source a file
func CheckIfInstalled(binary string, params ...string) bool {
	command := fmt.Sprintf("command -v %s", binary)
	err := shell.Command("sh", "-c", "command -v "+binary).Run()
	if err != nil {
		return false
	}

	if len(params) > 0 {
		command = fmt.Sprintf("source %s && command -v %s", params[0], binary)
		err := shell.Command("sh", "-c", command).Run()
		if err != nil {
			return false
		}
	}

	return false
}

func UniqueStringMerge(aString string, bString string) string {
	a := strings.Split(aString, "\n")
	b := strings.Split(bString, "\n")

	for _, line := range b {
		a = AppendIfMissing(a, line)
	}
	return strings.Join(a, "\n")
}

func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
