package server

import (
	"io"
	"log"
	"runtime"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
)

// goJsonRest
func goJSONRestHandler(w rest.ResponseWriter, req *rest.Request) {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	iow := w.(io.Writer)
	iow.Write(message)
}

func startGoJsonRest() {
	api := rest.NewApi()
	router, err := rest.MakeRouter(
		rest.Get("/hello", goJSONRestHandler))

	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	listenAndServe(port, api.MakeHandler())
}
