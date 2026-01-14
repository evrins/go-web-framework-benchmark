package server

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/teambition/gear"
)

// gear
func startGear(port int) {
	app := gear.New()
	router := gear.NewRouter()

	router.Get("/hello", func(c *gear.Context) error {
		if cpuBound {
			pow(target)
		} else {
			if SleepTime > 0 {
				time.Sleep(SleepTimeDuration)
			} else {
				runtime.Gosched()
			}
		}
		return c.HTML(200, messageStr)
	})
	app.UseHandler(router)
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
