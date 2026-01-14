package server

import (
	"runtime"
	"time"

	"github.com/emicklei/go-restful"
)

// goRestful
func goRestfulHandler(r *restful.Request, w *restful.Response) {
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

func startGoRestful(port int) {
	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(goRestfulHandler))
	wsContainer.Add(ws)

	listenAndServe(port, wsContainer)
}
