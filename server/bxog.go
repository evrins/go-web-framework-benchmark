package server

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	bxog "github.com/claygod/Bxog"
)

// bxog
func bxogHandler(w http.ResponseWriter, req *http.Request, r *bxog.Router) {
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

func startBxog(port int) {
	mux := bxog.New()
	mux.Add("/hello", bxogHandler)
	mux.Start(fmt.Sprintf(":%d", port))
}
