package server

import "time"

var (
	port              = 8080
	SleepTime         = 0
	cpuBound          bool
	target            = 15
	SleepTimeDuration time.Duration
	message           = []byte("hello world")
	messageStr        = "hello world"
	samplingPoint     = 20 // seconds
)
