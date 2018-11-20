package v1

import "time"

// ClientContainerConfig defines representation of the
// Configuration before starting of container client
type ClientContainerConfig struct {
	//CertPathEnv defines path to the cert files
	CertPathEnv string `json:"cert_path_env"`
	Endpoint    string `json:"endpoint"`
}

// Container define struct for container representation
type Container struct {
	CreatedTime time.Time `json:"created_time,omitempty"`
	Name        string    `json:"name"`
	State       string    `json:"state"`
	Status      string    `json:"status"`
	Image       string    `json:"image"`
	Names       []string  `json:"names"`
}

// ContainerStat provides definition for statistics
// on container
type ContainerStat struct {
	Stat []*Stat
}

// ContainerStat returns statistics for container
type Stat struct {
	Timestamp time.Time `json:"timestamp"`
	CPU       CPUStat   `json:"cpu_stat"`
	HasCPU    bool      `json:"has_cpu"`
}

// CPUStat defines statistics for cpu usage
type CPUStat struct {
}
