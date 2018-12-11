package container

import structs "github.com/saromanov/antenna/structs/v1"

// Container defines interface for container api
type Container interface {
	GetContainers(*structs.ListContainersOptions) ([]*structs.Container, error)
	Start() error
	Stop() error
	GetStats() *structs.ContainerStat
	Name() string
}
