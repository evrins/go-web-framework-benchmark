package server

import (
	"net/http"

	"github.com/vardius/gorouter/v4"
)

// gorouter
func startGorouter(port int) {
	router := gorouter.New()
	router.GET("/hello", http.HandlerFunc(helloHandler))
	listenAndServe(port, router)
}
