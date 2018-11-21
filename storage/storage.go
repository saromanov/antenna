package storage

import structs "github.com/saromanov/antenna/structs/v1"

// Storage defines main interface for Providing
// handling of storage data
type Storage interface {
	New(*Config) (Storage, error)
	Add(*structs.ContainerStat) error
	Close() error
	Search() ([]*structs.ContainerStat, error)
}

// Config defines configuration for Storage init
type Config struct {
	Name     string
	URL      string
	Username string
	Password string
}
