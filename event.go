package main

import (
	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	hook "github.com/robotn/gohook"
	"log"
	"runtime"
	"time"
)

type EventListener struct {
	server *RPCServer
}

func (s *EventListener) onSelectText(start hook.Event, end hook.Event) {
	oldVal, _ := clipboard.ReadAll()
	clipboard.WriteAll("")
	if runtime.GOOS == "darwin" {
		robotgo.KeyTap("c", "lcmd")
	} else {
		robotgo.KeyTap("c", "lctrl")
	}
	text, _ := clipboard.ReadAll()
	if len(text) != 0 {
		s.server.Notify("onSelectText", map[string]interface{}{
			"start": start,
			"end":   end,
			"text":  text,
		})
	}
	if err := clipboard.WriteAll(oldVal); err != nil {
		log.Println("Failed to clipboard.WriteAll", err)
	}
}

func (s *EventListener) onCopyText() {
	if val, err := clipboard.ReadAll(); err == nil {
		s.server.Notify("onCopyText", map[string]interface{}{
			"text": val,
		})
	}
}

func (s *EventListener) Listen() {
	var startEvent hook.Event

	hook.Register(hook.MouseHold, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["left"] {
			if startEvent == (hook.Event{}) {
				s.server.Notify("onMouseDown", e)
			}
			startEvent = e
		}
	})

	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		if startEvent != (hook.Event{}) {
			diff := e.When.Sub(startEvent.When)
			if diff > 200*time.Millisecond {
				s.server.Notify("onMouseUp", e)
				s.onSelectText(startEvent, e)
			}
			startEvent = hook.Event{}
		} else {
			s.server.Notify("onMouseDown", e)
		}
	})

	hook.Register(hook.MouseUp, []string{}, func(e hook.Event) {
		startEvent = hook.Event{}
		s.server.Notify("onMouseUp", e)
	})

	if runtime.GOOS == "darwin" {
		hook.Register(hook.KeyDown, []string{robotgo.KeyC, robotgo.Cmd}, func(e hook.Event) {
			s.onCopyText()
		})
	} else {
		hook.Register(hook.KeyDown, []string{robotgo.KeyC, robotgo.Ctrl}, func(e hook.Event) {
			s.onCopyText()
		})
	}

	hookEvent := hook.Start()

	<-hook.Process(hookEvent)
}
