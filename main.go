package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/lucasb-eyer/go-colorful"
	"math/rand"
)

var screenX, screenY = robotgo.GetScreenSize()
var centerX, centerY = GetScreenCenter(screenX, screenY)
var isEnabled = true
var minResponseTime = 80  // in milliseconds
var maxResponseTime = 130 // in milliseconds
var detectColor = "#9a24ab"
var lastColor = ""
var purple, _ = colorful.Hex(detectColor)

func main() {
	fmt.Println("[ sus software utility started! ]")

	fmt.Printf("> Screen resolution: %v x %v\n\n", screenX, screenY)
	fmt.Printf("> Software enabled: %v\n", isEnabled)
	fmt.Printf("> Min response time: %v\n", minResponseTime)
	fmt.Printf("> Max response time: %v\n", maxResponseTime)
	fmt.Printf("> Detect color: %v\n\n", detectColor)

	fmt.Printf("> (Use 'p' to toggle)\n\n")

	go func() {
		for {
			if robotgo.AddEvent("p") {
				isEnabled = !isEnabled
				fmt.Printf("> Software enabled: %v\n", isEnabled)
			}
		}
	}()

	for {
		color := robotgo.GetPixelColor(centerX, centerY)

		if !isEnabled {
			continue
		}

		if lastColor == color {
			continue
		}

		c, _ := colorful.Hex("#" + color)
		d := c.DistanceLab(purple)

		// for debugging only
		//fmt.Printf("Distance: %v\n", d)

		if d <= 0.4 {
			robotgo.KeyToggle("a", "up")
			robotgo.KeyToggle("s", "up")

			robotgo.MouseSleep = rand.Intn(maxResponseTime-minResponseTime) + minResponseTime
			robotgo.Click()
		}

		// for debugging only
		//fmt.Printf("Color of pixel at (%d, %d) is 0x%s\n", centerX, centerY, color)

		lastColor = color
	}
}

func GetScreenCenter(x int, y int) (int, int) {
	return x / 2, y / 2
}
