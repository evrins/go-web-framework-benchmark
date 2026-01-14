package server

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/indigo-web/indigo"
	ihttp "github.com/indigo-web/indigo/http"
	"github.com/indigo-web/indigo/router/inbuilt"
)

// indigo
func indigoHandler(r *ihttp.Request) *ihttp.Response {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}

	return ihttp.Bytes(r, message)
}

func startIndigo(port int) {
	r := inbuilt.New()
	r.Get("/hello", indigoHandler)

	err := indigo.New(fmt.Sprintf(":%d", port)).Serve(r)
	if err != nil {
		log.Fatal(err)
	}
}
