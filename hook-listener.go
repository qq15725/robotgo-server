package main

import (
	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	hook "github.com/robotn/gohook"
	"log"
	"runtime"
	"time"
)

type HookListener struct {
	server *RPCServer
}

func (s *HookListener) handleSelectText(x int16, y int16) {
	oldVal, _ := clipboard.ReadAll()
	if runtime.GOOS == "darwin" {
		robotgo.KeyTap(robotgo.KeyC, robotgo.Cmd)
	} else {
		robotgo.KeyTap(robotgo.KeyC, robotgo.Ctrl)
	}
	text, _ := clipboard.ReadAll()
	if len(text) != 0 {
		s.server.Notify("onSelectText", map[string]interface{}{
			"x":    x,
			"y":    y,
			"text": text,
		})
	}
	if err := clipboard.WriteAll(oldVal); err != nil {
		log.Println("Failed to clipboard.WriteAll", err)
	}
}

func (s *HookListener) Listen() {
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
				s.handleSelectText(e.X, e.Y)
			}
		}
	})

	hook.Register(hook.MouseUp, []string{}, func(e hook.Event) {
		isHold = false
	})

	hookEvent := hook.Start()

	<-hook.Process(hookEvent)
}
