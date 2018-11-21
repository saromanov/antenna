package storage

import structs "github.com/saromanov/antenna/structs/v1"

// Storage defines main interface for Providing
// handling of storage data
type Storage interface {
	Add(*structs.ContainerStat) error
	Close() error
	Search() ([]*structs.ContainerStat, error)
}

// New creates storage based on name
// At the init stage, its supports only InfluxDB
func New(conf *Config) (Storage, error) {
	if conf.Name == "" || conf.Name == "influxdb" {

	}
	return nil, nil
}

// Config defines configuration for Storage init
type Config struct {
	Name     string
	URL      string
	Username string
	Password string
}
