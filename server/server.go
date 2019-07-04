package server

import (
	"encoding/json"
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
	query := r.FormValue("query")
	response, err := s.st.Aggregate(&structs.AggregateSearchRequest{
		Request: query,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func Start(st storage.Storage, address string) {

	s := &server{st: st}
	http.HandleFunc("/v1/aggregate", s.AggregateMetrics)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
