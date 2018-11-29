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

type antenna struct {
	store storage.Storage
	httpClient *http.Client
}

func New()(Application, error){
	return {
		store: nil,
	}, nil
}