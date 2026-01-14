package server

import (
	"net/http"
	"runtime"
	"time"

	"github.com/julienschmidt/httprouter"
)

// httprouter
func httpRouterHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
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

func startHTTPRouter(port int) {
	mux := httprouter.New()
	mux.GET("/hello", httpRouterHandler)
	listenAndServe(port, mux)
}
