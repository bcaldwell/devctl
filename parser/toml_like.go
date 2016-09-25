package parser

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// WriteMapTomlLike writes a map of [string]string to a file using format key=value
func WriteMapTomlLike(m map[string]string, fileName string) error {
	r := ""

	for key, val := range m {
		r += fmt.Sprintf("%s=%s\n", key, val)
	}

	return ioutil.WriteFile(fileName, []byte(r), 0644)
}

// ReadTomlLike reads key=value file and returns map of [string]string
func ReadTomlLike(fileName string) (map[string]string, error) {
	m := make(map[string]string)
	// make sure file exists
	_, err := os.Stat(fileName)
	if err == nil {
		// regex to parse key=value
		r, err := regexp.Compile(`^([^\s#]+)\s*=\s*([^\s]+)$`)
		if err != nil {
			return m, err
		}
		// open file
		file, err := os.Open(fileName)
		if err != nil {
			return m, err
		}
		// read line by line and add each line to map m
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if res := r.FindAllStringSubmatch(scanner.Text(), 2); res != nil {
				key := res[0][1]
				val := res[0][2]
				m[key] = val
			}
		}
		return m, nil
	}
	return m, err
}
