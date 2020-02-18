package server

import (
	"encoding/json"
	"net/http"
)

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
