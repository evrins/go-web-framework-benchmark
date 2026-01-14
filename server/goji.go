package server

import (
	"net/http"
	"runtime"
	"time"

	goji "goji.io"
	gojipat "goji.io/pat"
)

// goji
func gojiHandler(w http.ResponseWriter, r *http.Request) {
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

func startGoji(port int) {
	mux := goji.NewMux()
	mux.HandleFunc(gojipat.Get("/hello"), gojiHandler)
	listenAndServe(port, mux)
}
