package antenna

import "github.com/saromanov/antenna/hardware"

type HostInfo struct {
	Cores int
	Hostname string
}

// hostWatcher provides checking of the host machine
// metrics 
func getStaticHostInfo() (*HostInfo, error) {
	host := &HostInfo{}
	cores := hardware.GetNumberOfCores()
	host.Cores = cores
	hn, err := hardware.GetHostname()
	if err != nil {
		return nil, err
	}
	host.Hostname = hn
	return host, nil
}