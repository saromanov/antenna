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
	Name         string            `json:"name,omitempty"`
	State        string            `json:"state,omitempty"`
	Status       string            `json:"status,omitempty"`
	Image        string            `json:"image,omitempty"`
	Names        []string          `json:"names,omitempty"`
	SizeRw       int64             `json:"size_rw,omitempty"`
	SizeRootFs   int64             `json:"size_root_fs,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
	RestartCount int               `json:"restart_count,omitempty"`
	Running      bool              `json:"running,omitempty"`
	Paused       bool              `json:"paused,omitempty"`
	Restarting   bool              `json:"restarting,omitempty"`
	OOMKilled    bool              `json:"oom_killed,omitempty"`
	Error        string            `json:"error,omitempty"`
	Args         []string          `json:"args,omitempty"`
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
