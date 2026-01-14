package server

import (
	"net/http"
	"runtime"
	"time"

	"github.com/naoina/denco"
)

// denco
func dencoHandler(w http.ResponseWriter, r *http.Request, params denco.Params) {
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

func startDenco(port int) {
	mux := denco.NewMux()
	handler, _ := mux.Build([]denco.Handler{mux.GET("/hello", dencoHandler)})
	listenAndServe(port, handler)
}
