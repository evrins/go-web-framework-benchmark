package server

import (
	"go101.org/tinyrouter"
)

// TinyRouter
func startTinyRouter(port int) {
	routes := []tinyrouter.Route{
		{
			Method:     "GET",
			Pattern:    "/hello",
			HandleFunc: helloHandler,
		},
	}
	router := tinyrouter.New(tinyrouter.Config{Routes: routes})
	listenAndServe(port, router)
}
