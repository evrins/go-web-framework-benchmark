package server

import (
	"runtime"
	"time"

	"github.com/go-playground/lars"
)

// lars
func larsHandler(c lars.Context) {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	c.Response().Write(message)
}

func startLars(port int) {
	mux := lars.New()
	mux.Get("/hello", larsHandler)
	listenAndServe(port, mux.Serve())
}
