package server

import (
	"runtime"
	"time"

	gf "github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// goframe
func gfHandler(r *ghttp.Request) {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	r.Response.Write(message)
}

func startGoframe(port int) {
	s := gf.Server()

	s.BindHandler("/hello", gfHandler)

	s.SetPort(port)
	s.Run()
}
