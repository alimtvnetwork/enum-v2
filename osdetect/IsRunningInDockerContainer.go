package osdetect

import (
	"os"
)

// IsRunningInDockerContainer
//
//  https://paulbradley.org/indocker/
//
//  docker creates a .dockerenv file at the root
//  of the directory tree inside the container.
//  if this file exists then the viewer is running
//  from inside a container so return true
func IsRunningInDockerContainer() bool {
	if _, err := os.Stat(dockerDetectPath); err == nil {
		return true
	}
	
	return false
}
