package server

import (
	"fmt"
	"runtime"
	"time"

	"github.com/dinever/golf"
)

func golfHandler(ctx *golf.Context) {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	ctx.Send(messageStr)
}

func startGolf(port int) {
	app := golf.New()
	app.Get("/hello", golfHandler)
	app.Run(fmt.Sprintf(":%d", port))
}
