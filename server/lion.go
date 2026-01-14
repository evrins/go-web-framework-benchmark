package server

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"gopkg.in/celrenheit/lion.v1"
)

// lion
func lionHandler(c context.Context, w http.ResponseWriter, r *http.Request) {
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

func startLion(port int) {
	mux := lion.New()
	mux.GetFunc("/hello", lionHandler)
	mux.Run(fmt.Sprintf(":%d", port))
}
