package server

import (
	"fmt"
	"log"
	"runtime"
	"time"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// fasthttprouting
func fastHTTPRoutingHandler(c *routing.Context) error {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	c.Write(message)
	return nil
}

func startFastHTTPRouting(port int) {
	mux := routing.New()
	mux.Get("/hello", fastHTTPRoutingHandler)
	err := fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), mux.HandleRequest)
	if err != nil {
		log.Fatal(err)
	}
}
