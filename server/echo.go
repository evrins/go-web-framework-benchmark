package server

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"
)

// echo
func echoHandler(c echo.Context) error {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	c.Response().Write(message)
	return nil
}

func startEcho(port int) {
	e := echo.New()
	e.GET("/hello", echoHandler)

	err := e.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
