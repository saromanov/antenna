package antenna

import (
	"net/http"
	"github.com/saromanov/antenna/storage"
)
// Application provides definition of the main
// interface for app
type Application interface{
	httpClient *http.Client 
}

type antenna struct {
	store storage.Storage
}

func New()(Application, error){
	return {
		store: nil,
	}, nil
}