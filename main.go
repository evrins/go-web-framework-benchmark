package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/smallnest/go-web-framework-benchmark/server"
	"github.com/spf13/pflag"
)

func main() {
	var webFramework string
	var port int
	var samplingPoint int
	var cpuBound bool
	var sleepTime int

	pflag.StringVarP(&webFramework, "webframework", "w", "web framework to test", "default")
	pflag.IntVarP(&port, "port", "p", 8080, "port to listen on")
	pflag.IntVarP(&samplingPoint, "samplingPoint", "s", 20, "sampling point in seconds")
	pflag.BoolVarP(&cpuBound, "cpuBound", "c", false, "use cpu bound")
	pflag.IntVarP(&sleepTime, "sleepTime", "t", 10, "sleep time in mills")

	pflag.Parse()

	if cpuBound {
		sleepTime = 0
	}

	startSampling(samplingPoint)

	wf, err := server.GetWebFramework(webFramework)
	if err != nil {
		log.Fatal(err)
	}
	wf(port)
}

func startSampling(samplingPoint int) {
	samplingPointDuration := time.Duration(samplingPoint) * time.Second

	go func() {
		time.Sleep(samplingPointDuration)
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		var u uint64 = 1024 * 1024
		fmt.Printf("TotalAlloc: %d\n", mem.TotalAlloc/u)
		fmt.Printf("Alloc: %d\n", mem.Alloc/u)
		fmt.Printf("HeapAlloc: %d\n", mem.HeapAlloc/u)
		fmt.Printf("HeapSys: %d\n", mem.HeapSys/u)
	}()
}
