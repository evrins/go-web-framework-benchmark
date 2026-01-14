package server

import (
	"github.com/go-zoo/bone"
)

// bone
func startBone(port int) {
	mux := bone.New()
	mux.HandleFunc("/hello", helloHandler)
	listenAndServe(port, mux)
}
