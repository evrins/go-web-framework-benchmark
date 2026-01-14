package server

import (
	"github.com/kataras/muxie"
)

func startMuxie(port int) {
	mux := muxie.NewMux()
	mux.HandleFunc("/hello", helloHandler)
	listenAndServe(port, mux)
}
