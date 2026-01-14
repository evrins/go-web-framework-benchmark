package server

import (
	"runtime"
	"time"

	ozzo "github.com/go-ozzo/ozzo-routing"
)

// go-ozzo
func ozzoHandler(c *ozzo.Context) error {
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

func startGoozzo(port int) {
	r := ozzo.New()
	r.Get("/hello", ozzoHandler)
	listenAndServe(port, r)
}
