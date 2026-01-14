package server

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/config"
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
	cfg := config.LoadDefault()

	cfg.Set("server.port", port)
	cfg.Set("server.host", "0.0.0.0")

	opts := goyave.Options{
		Config: cfg,
	}

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
