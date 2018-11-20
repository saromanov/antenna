package storage

import structs "github.com/saromanov/antenna/structs/v1"

type Storage interface {
	Add(*structs.ContainerStat) error
	Close() error
	Search() ([]*structs.ContainerStat, error)
}
