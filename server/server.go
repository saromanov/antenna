package server

import (
	"log"
	"net/http"

	"github.com/saromanov/antenna/storage"
)

// server defines api endpoints
type server struct {
	st storage.Storage
}

// Start provides initialization of the server
func Start(st storage.Storage, address string) {
	s := &server{st: st}
	http.HandleFunc("/v1/aggregate", s.AggregateMetrics)
	http.HandleFunc("/v1/info", s.Info)
	http.HandleFunc("/v1/search", s.Search)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
