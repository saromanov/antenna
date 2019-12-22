package server

import (
	"encoding/json"
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
	query := r.FormValue("query")
	if query == "" {
		http.Error(w, "query is not defined", http.StatusInternalServerError)
		return
	}
	response, err := s.st.Aggregate(&structs.AggregateSearchRequest{
		Request: query,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to aggregate query: %v", err), http.StatusInternalServerError)
		return
	}

	responseOutput := mapAggregateResponse(response)
	result, err := json.Marshal(responseOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (s *server) Info(w http.ResponseWriter, r *http.Request) {
	response := s.st.Info()
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
	http.HandleFunc("/v1.info", s.Info)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
