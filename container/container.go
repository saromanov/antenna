package container

// Container defines interface for container api
type Container interface {
	GetContainers()
	Start()
	Stop()
}
