package server

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// gin
func ginHandler(c *gin.Context) {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	c.Writer.Write(message)
}

func startGin(port int) {
	gin.SetMode(gin.ReleaseMode)
	mux := gin.New()
	mux.GET("/hello", ginHandler)
	err := mux.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
