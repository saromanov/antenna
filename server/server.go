package server

import (
	"log"
	"net/http"
)

func AggregateMetrics(w http.ResponseWriter, r *http.Request) {

}

func Start(address string) {
	http.HandleFunc("/aggregate", AggregateMetrics)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
