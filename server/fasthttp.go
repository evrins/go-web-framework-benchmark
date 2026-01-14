package server

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/valyala/fasthttp"
)

// fasthttp
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}

	ctx.Write(message)
}

func startFasthttp(port int) {
	s := &fasthttp.Server{
		Handler:                       fastHTTPHandler,
		GetOnly:                       true,
		NoDefaultDate:                 true,
		NoDefaultContentType:          true,
		DisableHeaderNamesNormalizing: true,
	}

	err := s.ListenAndServe(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
