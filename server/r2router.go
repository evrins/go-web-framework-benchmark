package server

import (
	"net/http"
	"runtime"
	"time"

	"github.com/vanng822/r2router"
)

// R2router
func r2routerHandler(w http.ResponseWriter, req *http.Request, params r2router.Params) {
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

func startR2router(port int) {
	mux := r2router.NewRouter()
	mux.Get("/hello", r2routerHandler)
	listenAndServe(port, mux)
}
