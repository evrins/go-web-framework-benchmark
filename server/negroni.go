package server

import (
	"net/http"

	"github.com/urfave/negroni"
)

// negroni
func startNegroni(port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	n := negroni.New()
	n.UseHandler(mux)

	listenAndServe(port, n)
}
