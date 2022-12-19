package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/gonutz/w32/v2"
	"github.com/lucasb-eyer/go-colorful"
	"time"
)

var screenX, screenY = robotgo.GetScreenSize()
var centerX, centerY = GetScreenCenter(screenX, screenY)
var isEnabled = true
var minRT = 1 // in milliseconds
var maxRT = 2 // in milliseconds
var dc, _ = colorful.Hex("#a145a3")

const (
	ad = 0.5
)

func main() {
	fmt.Println("[ sus software utility started! ]")

	fmt.Printf("> Screen resolution: %vx%v\n\n", screenX, screenY)
	fmt.Printf("> Software enabled: %v\n", isEnabled)
	fmt.Printf("> Min response time: %vms\n", minRT)
	fmt.Printf("> Max response time: %vms\n", maxRT)
	fmt.Printf("> Detect color: %v\n\n", dc.Hex())

	//fmt.Printf("> (Use '\\' to toggle)\n\n")

	//HideConsole()

	go func() {
		for {
			isEnabled = w32.GetAsyncKeyState(w32.VK_RBUTTON) == 32768
			time.Sleep(time.Duration(10) * time.Millisecond)
		}
	}()

	// extra checks
	//for i := 1; i < 10; i++ {
	//	go CreateScreenListener(0, i)
	//}
	//go CreateScreenListener(1, 0)
	//go CreateScreenListener(-1, 0)
	//go CreateScreenListener(0, -1)

	// center check
	CreateScreenListener(0, 0)
}

func CreateScreenListener(offsetX int, offsetY int) {
	for {
		color := robotgo.GetPixelColor(centerX+offsetX, centerY+offsetY)

		//fmt.Printf("Color at pixel (%v, %v): %v\n", centerX+offsetX, centerY+offsetY, color)

		if !isEnabled {
			continue
		}

		c, _ := colorful.Hex("#" + color)
		d := c.DistanceLab(dc)

		// for debugging only
		//fmt.Printf("Distance: %v\n", d)

		if d <= ad {
			robotgo.Click()

			time.Sleep(time.Duration(20) * time.Millisecond)
			continue
		}
	}
}

func HideConsole() {
	console := w32.GetConsoleWindow()
	if console == 0 {
		return // no console attached
	}
	_, consoleProcID := w32.GetWindowThreadProcessId(console)
	if w32.GetCurrentProcessId() == consoleProcID {
		w32.ShowWindowAsync(console, w32.SW_HIDE)
	}
}

func GetScreenCenter(x int, y int) (int, int) {
	return x / 2, y / 2
}
