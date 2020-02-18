package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	structs "github.com/saromanov/antenna/structs/v1"
)

func (s *server) Search(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "query is not defined", http.StatusInternalServerError)
		return
	}
	response, err := s.st.Search(&structs.ContainerStatSearch{
		Request: query,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, fmt.Sprintf("unable to aggregate query: %v", err), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, fmt.Sprintf("unable to marshal response: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
