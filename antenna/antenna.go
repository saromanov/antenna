package antenna

import (
	"net/http"

	"github.com/saromanov/antenna/storage"
)

// Application provides definition of the main
// interface for app
type Application struct {
	httpClient http.Client
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
	go a.startEventWatcher()
	return nil
}

func (a *Application) startEventWatcher() {
	select {
	case event := <-a.events:
		if event.event == ContainerAdd {

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
