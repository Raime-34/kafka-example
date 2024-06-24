package main

import "time"

const (
	topic            = "userInfo"
	network          = "tcp"
	host             = "localhost:9092"
	deadline         = -1
	messageWriteTime = 5 * time.Second
	messageReadTime  = 10 * time.Second
)
