package server

import (
	"net/http"
	"runtime"
	"time"

	"github.com/aeilang/httpz"
)

// httpz
func httpzHandler(w http.ResponseWriter, _ *http.Request) error {
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
	return nil
}

func startHTTPZ(port int) {
	mux := httpz.NewServeMux()
	mux.Get("/hello", httpzHandler)
	listenAndServe(port, mux)
}
