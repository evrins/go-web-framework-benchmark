package server

import (
	"fmt"
	"log"

	fasthttprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// fasthttp Router
func startFastHTTPRouter(port int) {
	mux := fasthttprouter.New()
	mux.GET("/hello", fastHTTPHandler)
	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), mux.Handler))
}
