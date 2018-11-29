package antenna

import (
	"net/http"
	"github.com/saromanov/antenna/storage"
)
// Application provides definition of the main
// interface for app
type Application struct {
	httpClient http.Client
	Store storage.Store
}

// Start provides starting of the app
func (a*Application) Start() error {
	return nil
}
type antenna struct {
	store storage.Storage
	httpClient *http.Client
}

func New()(Application, error){
	return {
		store: nil,
	}, nil
}