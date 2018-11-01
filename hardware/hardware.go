package hardware

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"runtime"
	"strconv"
)

const procMem = "/proc/meminfo"

var memoryTotalRegexp = regexp.MustCompile(`MemTotal:\s*([0-9]+) kB`)

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
	f, err := ioutil.ReadFile(procMem)
	if err != nil {
		return 0, err
	}

	memoryCapacity, err := parse(f, memoryTotalRegexp)
	if err != nil {
		return 0, err
	}
	return memoryCapacity, err
}

// parse is a helpful method for parsing of values
func parse(b []byte, r *regexp.Regexp) (uint64, error) {
	matches := r.FindSubmatch(b)
	if len(matches) != 2 {
		return 0, fmt.Errorf("failed to match regexp in output: %q", string(b))
	}
	m, err := strconv.ParseUint(string(matches[1]), 10, 64)
	if err != nil {
		return 0, err
	}

	// Convert to bytes.
	return m * 1024, err
}
