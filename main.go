package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/kindlyfire/go-keylogger"
	"github.com/lucasb-eyer/go-colorful"
	"math/rand"
	"time"
)

var screenX, screenY = robotgo.GetScreenSize()
var centerX, centerY = GetScreenCenter(screenX, screenY)
var isEnabled = true
var minResponseTime = 100 // in milliseconds
var maxResponseTime = 120 // in milliseconds
var detectColor = "#9a24ab"
var lastColor = ""
var purple, _ = colorful.Hex(detectColor)

const (
	delayKeyFetch   = 10 // in milliseconds
	delayColorFetch = 5  // in milliseconds
)

func main() {
	fmt.Println("[ sus software utility started! ]")

	fmt.Printf("> Screen resolution: %v x %v\n\n", screenX, screenY)
	fmt.Printf("> Software enabled: %v\n", isEnabled)
	fmt.Printf("> Min response time: %v ms\n", minResponseTime)
	fmt.Printf("> Max response time: %v ms\n", maxResponseTime)
	fmt.Printf("> Detect color: %v\n\n", detectColor)

	fmt.Printf("> (Use '\\' to toggle)\n\n")

	go func() {
		kl := keylogger.NewKeylogger()

		for {
			key := kl.GetKey()

			// only for debugging
			//if !key.Empty {
			//	fmt.Printf("Keycode: %v | Rune: %v\n", key.Keycode, key.Rune)
			//}

			// keycode = \
			if (!key.Empty) && (key.Keycode == 220 && key.Rune == 92) {
				isEnabled = !isEnabled
				fmt.Printf("> Software enabled: %v\n", isEnabled)
			}

			time.Sleep(delayKeyFetch * time.Millisecond)
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

		if d <= 0.5 {
			robotgo.KeyToggle("a", "up")
			robotgo.KeyToggle("s", "up")

			robotgo.MouseSleep = rand.Intn(maxResponseTime-minResponseTime) + minResponseTime
			robotgo.Click()
		}

		// for debugging only
		//fmt.Printf("Color of pixel at (%d, %d) is 0x%s\n", centerX, centerY, color)

		lastColor = color

		//time.Sleep(delayColorFetch * time.Millisecond)
	}
}

func GetScreenCenter(x int, y int) (int, int) {
	return x / 2, y / 2
}
