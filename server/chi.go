package server

import (
	"github.com/go-chi/chi/v5"
)

// chi
func startChi(port int) {
	// Create a router instance.
	r := chi.NewRouter()

	// Register route handler.
	r.Get("/hello", helloHandler)

	// Start Chi.
	listenAndServe(port, r)
}
