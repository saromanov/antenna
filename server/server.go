package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/saromanov/antenna/storage"
	structs "github.com/saromanov/antenna/structs/v1"
)

// server defines api endpoints
type server struct {
	st storage.Storage
}

// AggregateMetrics defines api endpoint for getting aggregated
// metrics from container
func (s *server) AggregateMetrics(w http.ResponseWriter, r *http.Request) {
	response, err := s.st.Aggregate(&structs.AggregateSearchRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(response)
}

func Start(st storage.Storage, address string) {

	s := &server{st: st}
	http.HandleFunc("/aggregate", s.AggregateMetrics)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
