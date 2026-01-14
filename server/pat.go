package server

import (
	"net/http"

	"github.com/bmizerany/pat"
)

// pat
func startPat(port int) {
	mux := pat.New()
	mux.Get("/hello", http.HandlerFunc(helloHandler))
	listenAndServe(port, mux)
}
