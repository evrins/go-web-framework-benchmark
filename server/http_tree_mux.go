package server

import (
	"net/http"
	"runtime"
	"time"

	"github.com/dimfeld/httptreemux"
)

// httpTreeMux
func httpTreeMuxHandler(w http.ResponseWriter, _ *http.Request, vars map[string]string) {
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

func starthttpTreeMux(port int) {
	mux := httptreemux.New()
	mux.GET("/hello", httpTreeMuxHandler)
	listenAndServe(port, mux)
}
