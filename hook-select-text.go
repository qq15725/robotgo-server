package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	hook "github.com/robotn/gohook"
	"runtime"
	"time"
)

func HandleSelectText(x int16, y int16) {
	oldVal, _ := clipboard.ReadAll()
	if runtime.GOOS == "darwin" {
		robotgo.KeyTap(robotgo.KeyC, robotgo.Cmd)
	} else {
		robotgo.KeyTap(robotgo.KeyC, robotgo.Ctrl)
	}
	val, _ := clipboard.ReadAll()
	if len(val) != 0 {
		fmt.Println(x, y)
		fmt.Println(val)
	}
	clipboard.WriteAll(oldVal)
}

func HookSelectText() {
	isHold := false
	starTime := time.Now()

	hook.Register(hook.MouseHold, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["left"] {
			starTime = time.Now()
			isHold = true
		}
	})

	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		if isHold {
			isHold = false
			diff := time.Now().Sub(starTime)
			if diff > 200*time.Millisecond {
				HandleSelectText(e.X, e.Y)
			}
		}
	})

	hook.Register(hook.MouseUp, []string{}, func(e hook.Event) {
		isHold = false
	})

	s := hook.Start()
	<-hook.Process(s)
}
