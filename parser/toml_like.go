package parser

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func WriteMapTomlLike(m map[string]string, fileName string) error {
	r := ""

	for key, val := range m {
		r += fmt.Sprintf("%s=%s\n", key, val)
	}

	return ioutil.WriteFile(fileName, []byte(r), 0644)
}

func ReadTomlLike(fileName string) (map[string]string, error) {
	m := make(map[string]string)
	if _, err := os.Stat(fileName); err == nil {
		r, err := regexp.Compile(`^([^\s#]+)\s*=\s*([^\s]+)$`)
		if err != nil {
			return m, err
		}
		file, err := os.Open(fileName)
		if err != nil {
			return m, err
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if res := r.FindAllStringSubmatch(scanner.Text(), 2); res != nil {
				key := res[0][1]
				val := res[0][2]
				m[key] = val
			}
		}
		return m, nil
	} else {
		return m, err
	}
}
