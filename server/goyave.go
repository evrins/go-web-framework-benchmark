package server

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"goyave.dev/goyave/v5"
)

// Goyave
func goyaveHandler(r *goyave.Response, req *goyave.Request) {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	r.String(http.StatusOK, "hello")
}

func getGoyaveRoutes(server *goyave.Server, router *goyave.Router) {
	router.Get("/hello", goyaveHandler)
}

func startGoyave(port int) {
	opts := goyave.Options{}
	opts.Config.Set("server.port", port)
	opts.Config.Set("server.host", "0.0.0.0")

	server, err := goyave.New(opts)
	if err != nil {
		log.Fatal(err)
	}

	server.RegisterRoutes(getGoyaveRoutes)

	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
