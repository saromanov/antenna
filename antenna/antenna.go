package antenna

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/saromanov/antenna/container/docker"
	"github.com/saromanov/antenna/storage"
	structs "github.com/saromanov/antenna/structs/v1"
)

// Application provides definition of the main
// interface for app
type Application struct {
	HTTPClient     http.Client
	Store          storage.Storage
	events         chan *ContainerEvent
	dockerClient   *docker.Docker
	watcher        *containerWatcher
	containers     map[string]*structs.Container
	containersLock *sync.RWMutex
}

type containerInfo struct {
	Name string
	ID   string
}

// ContainerEventType provides definition for container event handling
type ContainerEventType int

const (
	// ContainerAdd defines event for adding a new container
	ContainerAdd ContainerEventType = iota + 1
	// ContainerRemove defines event for removing old one container
	ContainerRemove
	// ListContainers returns list of running containers
	ListContainers
)

// ContainerEvent event defines events on containers
type ContainerEvent struct {
	event      ContainerEventType
	containers []*structs.Container
}

// Start provides starting of the app
func (a *Application) Start() error {
	a.events = make(chan *ContainerEvent)
	a.connectToDocker()
	a.watcher = &containerWatcher{
		dockerClient: a.dockerClient,
		events:       a.events,
	}
	a.watcher.Watch()
	a.startEventWatcher()
	return nil
}

// connectToDocker creates connection to docker via client
func (a *Application) connectToDocker() {
	client := docker.Init(&structs.ClientContainerConfig{})
	a.dockerClient = client
}

func (a *Application) addContainer() {

}

func (a *Application) removeContainer() {

}

func (a *Application) processListContainers(containers []*structs.Container) {
	for _, c := range containers {
		a.containers[c.Name] = c
	}
}

func (a *Application) startEventWatcher() {
	select {
	case event := <-a.events:
		switch event.event {
		case ContainerAdd:
			a.addContainer()
		case ContainerRemove:
			a.removeContainer()
		case ListContainers:
			a.processListContainers(event.containers)
		}

	}
}

type antenna struct {
	store      storage.Storage
	httpClient *http.Client
}

// New provides initialization on the app
func New() (*Application, error) {
	return &Application{
		Store:          nil,
		containersLock: &sync.RWMutex{},
		containers:     make(map[string]*structs.Container),
	}, nil
}

// GetContainerInfo returns info about running container
func (a *Application) GetContainerInfo(name string) (*structs.Container, error) {
	var cont *structs.Container
	var ok bool
	func() {
		a.containersLock.RLock()
		defer a.containersLock.RUnlock()
		cont, ok = a.containers[name]
	}()
	if !ok {
		return nil, fmt.Errorf("unknown container %q", name)
	}
	return cont, nil
}
