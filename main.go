package main

import (
	"context"
	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.KeySleep = 100
	robotgo.MouseSleep = 100

	go HookSelectText()

	NewRobotgoServer().Serve(context.Background())
}
