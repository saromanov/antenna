package hardware

import (
	"os"
	"runtime"
)

// GetNumberOfCores returns number of cores at machine
func GetNumberOfCores() int {
	return runtime.NumCPU()
}

// GetHostname returns hostname of machine
func GetHostname() (string, error) {
	return os.Hostname()
}
