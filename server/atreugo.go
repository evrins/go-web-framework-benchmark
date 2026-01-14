package server

import (
	"log"
	"runtime"
	"strconv"
	"time"

	"github.com/savsgio/atreugo/v11"
)

// atreugo
func atreugoHandler(ctx *atreugo.RequestCtx) error {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}

	return ctx.TextResponse(messageStr)
}
func startAtreugo(port int) {
	server := atreugo.New(atreugo.Config{
		Addr:                          ":" + strconv.Itoa(port),
		Prefork:                       true,
		NoDefaultDate:                 true,
		NoDefaultContentType:          true,
		DisableHeaderNamesNormalizing: true,
	})
	server.GET("/hello", atreugoHandler)
	log.Fatal(server.ListenAndServe())
}
