package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const ledPin = 17

func main() {
	// init
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	// LED pin output
	pin := rpio.Pin(ledPin)
	pin.Output()
	fmt.Println(pin)

	// Blink 5 times
	for i := 0; i < 5; i++ {
		fmt.Printf("i: %d", i)

		// Light on
		pin.High()
		time.Sleep(500 * time.Millisecond)

		// Light off
		pin.Low()
		time.Sleep(500 * time.Millisecond)
	}
}
