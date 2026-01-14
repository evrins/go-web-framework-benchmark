package server

import (
	"github.com/gorilla/mux"
)

// gorilla
func startGorilla(port int) {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler).Methods("GET")
	listenAndServe(port, router)
}
