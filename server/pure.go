package server

import (
	"github.com/go-playground/pure"
)

// pure
func startPure(port int) {
	p := pure.New()
	p.Get("/hello", helloHandler)
	listenAndServe(port, p.Serve())
}
