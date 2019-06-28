package antenna

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

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
	startTime      time.Time
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
	name       string
}

// Start provides starting of the app
func (a *Application) Start() error {
	a.startTime = time.Now().UTC()
	a.events = make(chan *ContainerEvent)
	a.containersLock = &sync.RWMutex{}
	a.containers = make(map[string]*structs.Container)
	a.connectToDocker()
	a.watcher = &containerWatcher{
		dockerClient: a.dockerClient,
		events:       a.events,
	}
	go a.watcher.Watch()
	a.startEventWatcher()
	return nil
}

// connectToDocker creates connection to docker via client
func (a *Application) connectToDocker() {
	client := docker.Init(&structs.ClientContainerConfig{})
	a.dockerClient = client
}

// getContainers returns map of containers
func (a *Application) getContainers() map[string]*structs.Container {
	return a.containers
}

// addContainer creating of the event after adding of the new container
func (a *Application) addContainer() {
	fmt.Println("Adding container")
}

func (a *Application) removeContainer(name string) {
	func() {
		a.containersLock.RLock()
		defer a.containersLock.RUnlock()
		delete(a.containers, name)
	}()
}

func (a *Application) processListContainers(containers []*structs.Container) {
	a.containersLock.RLock()
	defer a.containersLock.RUnlock()
	old := a.containers
	for _, c := range containers {
		container, _ := a.dockerClient.Get(c.ID)
		stats := a.dockerClient.GetStats(container.ID)
		if err := a.insertStats(stats); err != nil {
			fmt.Println("unable to insert stats: ", err)
		}
		fmt.Println(container.Name, container.Running)
		a.containers[c.Name] = c
	}

	go func(p, n map[string]*structs.Container) {
		if len(old) > len(a.containers) {
			for _, c := range old {
				found := false
				for _, c2 := range a.containers {
					if c.ID == c2.ID {
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Container was removed: ", c.ID)
				}
			}
		}
		return
	}(old, a.containers)
	/*if oldSize < len(a.containers) {
		a.events <- &ContainerEvent{
			event: ContainerAdd,
		}
	}*/
}

// insertStat provides inserting of the container stat to the storage
func (a *Application) insertStats(stat *structs.ContainerStat) error {
	if a.Store == nil {
		return errors.New("storage is not defined")
	}
	return a.Store.Add(stat)
}

func (a *Application) startEventWatcher() {
	for {
		select {
		case event := <-a.events:
			switch event.event {
			case ContainerAdd:
				a.addContainer()
			case ContainerRemove:
				a.removeContainer(event.name)
			case ListContainers:
				a.processListContainers(event.containers)
			}

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
