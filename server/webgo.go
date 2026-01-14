package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/naughtygopher/webgo/v7"
)

// webgo
func getWebgoRoutes() []*webgo.Route {
	return []*webgo.Route{
		{
			Name:     "hello",
			Method:   http.MethodGet,
			Pattern:  "/hello",
			Handlers: []http.HandlerFunc{helloHandler},
		},
	}
}

func startWebgo(port int) {
	cfg := webgo.Config{
		Host:         "",
		Port:         strconv.Itoa(port),
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}
	router := webgo.NewRouter(&cfg, getWebgoRoutes()...)
	router.Start()
}
