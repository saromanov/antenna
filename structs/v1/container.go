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
	ID           string            `json:"ID"`
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
	Timestamp      time.Time `json:"timestamp"`
	CPU            CPUStat   `json:"cpu_stat"`
	HasCPU         bool      `json:"has_cpu"`
	NumProcs       int32     `json:"num_procs"`
	Cache          int64     `json:"cache"`
	Limit          int64     `json:"limit"`
	Usage          int64     `json:"usage"`
	MaxUsage       int64     `json:"max_usage"`
	Image          string    `json:"image"`
	Name           string    `json:"name"`
	ReadSizeBytes  int64     `json:"read_size_bytes"`
	WriteSizeBytes int64     `json:"write_size_bytes"`
	TxPackets      int64     `json:"tx_packets,omitempty"`
	TxDropped      int64     `json:"tx_dropped,omitempty"`
	RxPackets      int64     `json:"rx_packets,omitempty"`
	TxErrors       int64     `json:"tx_errors,omitempty"`
	TxBytes        int64     `json:"tx_bytes,omitempty"`
}

// ContainerStat returns statistics for container
type Stat struct {
	Timestamp time.Time `json:"timestamp"`
	CPU       CPUStat   `json:"cpu_stat"`
	HasCPU    bool      `json:"has_cpu"`
}

// CPUStat defines statistics for cpu usage
type CPUStat struct {
	TotalUsage      int64 `json:"total_usage"`
	SystemCPUUsage  int64 `json:"system_cpu_usage"`
	OnlineCPUs      int64 `json:"online_cpus"`
	UsageInUsermode int64 `json:"usage_in_usermode"`
}

// ContainerStatSearch provides searching
// of stat on the storage
type ContainerStatSearch struct {
	FromTimestamp time.Time `json:"timestamp"`
	ToTimestamp   time.Time `json:"timestamp"`
	Request       string    `json:"request"`
}

// AggregateSearchRequest provides request
// for aggregation for storage
type AggregateSearchRequest struct {
	Request string `json:"request"`
}

// AggregateSearchResponse provides response for
// aggregation for storage
type AggregateSearchResponse struct {
	Count uint64 `json:"count"`
}
