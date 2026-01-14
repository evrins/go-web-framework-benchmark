package server

import (
	"runtime"
	"time"

	"github.com/gocraft/web"
)

// gocraftWeb
type gocraftWebContext struct{}

func gocraftWebHandler(w web.ResponseWriter, r *web.Request) {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}

func startGocraftWeb(port int) {
	mux := web.New(gocraftWebContext{})
	mux.Get("/hello", gocraftWebHandler)
	listenAndServe(port, mux)
}
