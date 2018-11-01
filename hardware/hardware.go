package hardware

import (
	"io/ioutil"
	"os"
	"regexp"
	"runtime"
)

const (
	procMem           = "/proc/meminfo"
	memoryTotalRegexp = regexp.MustCompile(`MemTotal:\s*([0-9]+) kB`)
)

// GetNumberOfCores returns number of cores at machine
func GetNumberOfCores() int {
	return runtime.NumCPU()
}

// GetHostname returns hostname of machine
func GetHostname() (string, error) {
	return os.Hostname()
}

// GetMemoryCapacity returns total memory from /proc/meminfo.
func GetMemoryCapacity() (uint64, error) {
	out, err := ioutil.ReadFile(procMem)
	if err != nil {
		return 0, err
	}

	memoryCapacity, err := parseCapacity(out, memoryCapacityRegexp)
	if err != nil {
		return 0, err
	}
	return memoryCapacity, err
}
