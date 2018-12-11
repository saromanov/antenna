package antenna

import (
	"fmt"
	"net/http"

	"github.com/saromanov/antenna/container/docker"
	"github.com/saromanov/antenna/storage"
	structs "github.com/saromanov/antenna/structs/v1"
)

// Application provides definition of the main
// interface for app
type Application struct {
	HTTPClient http.Client
	Store      storage.Storage
	events     chan *ContainerEvent
}

// ContainerEventType provides definition for container event handling
type ContainerEventType int

const (
	// ContainerAdd defines event for adding a new container
	ContainerAdd ContainerEventType = iota + 1
	// ContainerRemove defines event for removing old one container
	ContainerRemove
)

// ContainerEvent event defines events on containers
type ContainerEvent struct {
	event ContainerEventType
}

// Start provides starting of the app
func (a *Application) Start() error {
	a.events = make(chan *ContainerEvent)
	a.connectToDocker()
	go a.startEventWatcher()
	return nil
}

// connectToDocker creates connection to docker via client
func (a *Application) connectToDocker() {
	client := docker.Init(&structs.ClientContainerConfig{})
	fmt.Println(client)
}

func (a *Application) addContainer() {

}

func (a *Application) removeContainer() {

}
func (a *Application) startEventWatcher() {
	select {
	case event := <-a.events:
		switch event.event {
		case ContainerAdd:
			a.addContainer()
		case ContainerRemove:
			a.removeContainer()
		}

	}
}

type antenna struct {
	store      storage.Storage
	httpClient *http.Client
}

// New provides initialization on the app
func New() (Application, error) {
	return Application{
		Store: nil,
	}, nil
}
