package server

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/beego/beego"
	beegoContext "github.com/beego/beego/context"
)

// beego
func beegoHandler(ctx *beegoContext.Context) {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	ctx.WriteString(messageStr)
}

func startBeego(port int) {
	beego.BConfig.RunMode = beego.PROD
	beego.BeeLogger.Close()
	mux := beego.NewControllerRegister()
	mux.Get("/hello", beegoHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
