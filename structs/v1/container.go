package v1

import (
	"context"
	"time"
)

// ClientContainerConfig defines representation of the
// Configuration before starting of container client
type ClientContainerConfig struct {
	//CertPathEnv defines path to the cert files
	CertPathEnv string `json:"cert_path_env"`
	Endpoint    string `json:"endpoint"`
}

// ListContainersOptions provides options for searching containers
// copyed from docker client definition
type ListContainersOptions struct {
	All     bool
	Size    bool
	Limit   int
	Since   string
	Before  string
	Filters map[string][]string
	Context context.Context
}

// StartContainerOptions defines options for start container
type StartContainerOptions struct {
	ID string
}

// ContainerInfo returns info about running container
type ContainerInfo struct {
}

// Container define struct for container representation
type Container struct {
	CreatedTime  time.Time         `json:"created_time,omitempty"`
	Name         string            `json:"name"`
	State        string            `json:"state"`
	Status       string            `json:"status"`
	Image        string            `json:"image"`
	Names        []string          `json:"names"`
	SizeRw       int64             `json:"size_rw"`
	SizeRootFs   int64             `json:"size_root_fs"`
	Labels       map[string]string `json:"labels"`
	RestartCount int               `json:"restart_count"`
	Running      bool              `json:"running"`
	Paused       bool              `json:"paused"`
	Restarting   bool              `json:"restarting"`
	OOMKilled    bool              `json:"oom_killed"`
	Error        string            `json:"error"`
	Args         []string          `json:"args"`
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
