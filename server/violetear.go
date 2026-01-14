package server

import (
	"github.com/nbari/violetear"
)

// violetear
func startVioletear(port int) {
	mux := violetear.New()
	mux.HandleFunc("/hello", helloHandler)
	listenAndServe(port, mux)
}
