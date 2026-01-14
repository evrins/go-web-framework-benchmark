package server

import (
	"net/http"

	"github.com/rcrowley/go-tigertonic"
)

// Tiger Tonic
func startTigerTonic(port int) {
	mux := tigertonic.NewTrieServeMux()
	mux.Handle("GET", "/hello", http.HandlerFunc(helloHandler))
	listenAndServe(port, mux)
}
