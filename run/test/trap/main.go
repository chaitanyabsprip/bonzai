package main

import (
	"syscall"
	"time"

	"github.com/chaitanyabsprip/bonzai/run"
)

func main() {
	run.Trap(nil, syscall.SIGINT, syscall.SIGTERM)
	for {
		time.Sleep(1 * time.Millisecond)
	}
}
