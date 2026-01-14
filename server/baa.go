package server

import (
	"runtime"
	"strconv"
	"time"

	"gopkg.in/baa.v1"
)

// baa
func baaHandler(ctx *baa.Context) {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	ctx.Text(200, message)
}

func startBaa(port int) {
	mux := baa.New()
	mux.Get("/hello", baaHandler)
	mux.Run(":" + strconv.Itoa(port))
}
