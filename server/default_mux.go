package server

import (
	"net/http"
	"runtime"
	"time"
)

// default mux
func helloHandler(w http.ResponseWriter, r *http.Request) {
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

func startDefaultMux(port int) {
	http.HandleFunc("/hello", helloHandler)
	listenAndServe(port, nil)
}
