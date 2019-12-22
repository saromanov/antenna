package storage

import structs "github.com/saromanov/antenna/structs/v1"

// Storage defines main interface for Providing
// handling of storage data
type Storage interface {
	Add(*structs.ContainerStat) error
	Close() error
	Search(*structs.ContainerStatSearch) ([]*structs.ContainerStat, error)
	Aggregate(*structs.AggregateSearchRequest) (*structs.AggregateSearchResponse, error)
	Info() map[string]interface{}
}

// Config defines configuration for Storage init
type Config struct {
	Name     string
	URL      string
	Username string
	Password string
	Database string
	Token    string
}
