package antenna

import "github.com/saromanov/antenna/hardware"

type HostInfo struct {
	Cores int
}

// hostWatcher provides checking of the host machine
// metrics 
func getStaticHostInfo() (*HostInfo, error) {
	host := &HostInfo{}
	cores := hardware.GetNumberOfCores()
	host.Cores = cores
	return host, nil
}