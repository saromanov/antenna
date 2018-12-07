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

// Container event defines events on containers
type ContainerEvent struct {
}

// Start provides starting of the app
func (a *Application) Start() error {
	return nil
}

func (a *Application) startEventWatcher() {
	select {
	case event := <-a.events:

	}
}

type antenna struct {
	store      storage.Storage
	httpClient *http.Client
}

// New provides initialization on the app
func New() (Application, error) {
	return Application{
		store: nil,
	}, nil
}
