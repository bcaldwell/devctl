package utilities

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"

	"strings"

	"github.com/bcaldwell/devctl/shell"
	"github.com/bcaldwell/go-printer"
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

func Check(err error, message string) bool {
	if err != nil {
		printer.Fail("%s failed with %s", message, err)
		return true
	}
	return false
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
func HandleError(err error, args ...string) bool {
	if err != nil {
		if len(args) > 1 {
			message := args[0]
			printer.Fail("%s failed with %s", message, err)
		} else {
			printer.Fail("%s", err)
		}
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

// CheckIfInstalled usees command -v to check if a binary is current in the path. Returns true if it is installed
func CheckIfInstalled(binary string) bool {
	err := shell.Command("command", "-v", binary).Run()
	return (err == nil)
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

func RemoveFromPath(s string) string {
	pathParts := strings.Split(os.Getenv("PATH"), ":")
	var filtedParts []string
	for _, part := range pathParts {
		if part != s {
			filtedParts = append(filtedParts, part)
		}
	}
	return strings.Join(filtedParts, ":")
}

func LocationHash() string {
	dir, err := os.Getwd()
	if err == nil {
		return fmt.Sprintf("%x", md5.Sum([]byte(dir)))
	}
	HandleError(err)
	return ""
}

func HTTPDownload(uri string) ([]byte, error) {
	res, err := http.Get(uri)
	HandleError(err)

	defer res.Body.Close()
	d, err := ioutil.ReadAll(res.Body)
	HandleError(err)

	return d, err
}

func WriteFile(dst string, d []byte) error {
	err := ioutil.WriteFile(dst, d, 0444)
	HandleError(err)
	return err
}

func DownloadToFile(uri string, dst string) error {
	d, err := HTTPDownload(uri)
	if err == nil {
		return WriteFile(dst, d)
	}
	return err
}

func LineWith(lines string, text string) string {
	for _, line := range strings.Split(lines, "\n") {
		if strings.Contains(line, text) {
			return line
		}
	}
	return ""
}
