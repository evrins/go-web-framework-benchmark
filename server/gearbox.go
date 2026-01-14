package server

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/gogearbox/gearbox"
)

// gearbox
func startGearbox(port int) {
	gb := gearbox.New()
	gb.Get("/hello", func(ctx gearbox.Context) {
		if cpuBound {
			pow(target)
		} else {
			if SleepTime > 0 {
				time.Sleep(SleepTimeDuration)
			} else {
				runtime.Gosched()
			}
		}
		ctx.SendString(messageStr)
	})
	err := gb.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
