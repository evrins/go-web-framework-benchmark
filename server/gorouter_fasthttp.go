package server

import (
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
	"github.com/vardius/gorouter/v4"
)

func startGorouterFastHTTP(port int) {
	router := gorouter.NewFastHTTPRouter()
	router.GET("/hello", fastHTTPHandler)
	err := fasthttp.ListenAndServe(":"+strconv.Itoa(port), router.HandleFastHTTP)
	if err != nil {
		log.Fatal(err)
	}
}
