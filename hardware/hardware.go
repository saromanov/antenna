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

// Info provides definition for hardware info
type Info struct {
	CPUInfo        string
	MemoryCapacity uint64
}

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

// GetInfo retruns machien info
func GetInfo() (*Info, error) {
	cpuInfo, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return nil, fmt.Errorf("unable to get cpu info: %v", err)
	}
	memCap, err := GetMemoryCapacity()
	if err != nil {
		return nil, fmt.Errorf("unable to get memory capacity: %v", err)
	}
	return &Info{
		CPUInfo:        string(cpuInfo),
		MemoryCapacity: memCap,
	}, nil
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
