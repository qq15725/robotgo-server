package main

import "github.com/go-vgo/robotgo"

var methodMap = map[string]interface{}{
	// Mouse
	"ScrollDir":    robotgo.ScrollDir,
	"Scroll":       robotgo.Scroll,
	"MilliSleep":   robotgo.MilliSleep,
	"ScrollSmooth": robotgo.ScrollSmooth,
	"Move":         robotgo.Move,
	"MoveRelative": robotgo.MoveRelative,
	"DragSmooth":   robotgo.DragSmooth,
	"Click":        robotgo.Click,
	"MoveSmooth":   robotgo.MoveSmooth,
	"Toggle":       robotgo.Toggle,
	// Keyboard
	"TypeStr":   robotgo.TypeStr,
	"Sleep":     robotgo.Sleep,
	"KeyTap":    robotgo.KeyTap,
	"KeyToggle": robotgo.KeyToggle,
	"WriteAll":  robotgo.WriteAll,
	"ReadAll":   robotgo.ReadAll,
	// Screen
	"Location":      robotgo.Location,
	"GetPixelColor": robotgo.GetPixelColor,
	"CaptureScreen": robotgo.CaptureScreen,
	"DisplaysNum":   robotgo.DisplaysNum,
}
