package osdetect

import (
	"io/ioutil"
	"strings"
)

// https://t.ly/U9NG
func readTrimmedFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	
	return strings.TrimSpace(string(data))
}
