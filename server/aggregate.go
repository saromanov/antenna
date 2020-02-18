package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	structs "github.com/saromanov/antenna/structs/v1"
)

// AggregateMetrics defines api endpoint for getting aggregated
// metrics from container
func (s *server) AggregateMetrics(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "query is not defined", http.StatusInternalServerError)
		return
	}
	response, err := s.st.Aggregate(&structs.AggregateSearchRequest{
		Request: query,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
