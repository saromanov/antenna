package storage

import structs "github.com/saromanov/antenna/structs/v1"

type Storage interface {
	Add(*structs.ContainerStat) error
	Close() error
	Search() ([]*structs.ContainerStat, error)
}

// New creates storage based on name
// At the init stage, its supports only InfluxDB
// So, "name" attribute is unused at this moment
func New(name string) (Storage, error) {
	return nil, nil
}
