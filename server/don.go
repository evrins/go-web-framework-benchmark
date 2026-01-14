package server

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/abemedia/go-don"
	_ "github.com/abemedia/go-don/encoding/text"
)

// don
func donHandler(context.Context, any) ([]byte, error) {
	if cpuBound {
		pow(target)
	} else {
		if SleepTime > 0 {
			time.Sleep(SleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	return message, nil
}

func startDon(port int) {
	api := don.New(nil)
	api.Get("/hello", don.H(donHandler))
	err := api.ListenAndServe(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
