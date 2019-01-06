package container

import structs "github.com/saromanov/antenna/structs/v1"

// Container defines interface for container api
type Container interface {
	List(*structs.ListContainersOptions) ([]*structs.Container, error)
	Start(*structs.StartContainerOptions) error
	Stop() error
	GetStats(id string) *structs.ContainerStat
	Get(id string) (*structs.Container, error)
	Name() string
}
