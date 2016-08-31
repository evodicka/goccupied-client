package handler

import (
	"github.com/evodicka/goccupied-client/output"
	"github.com/stianeikeland/go-rpio"
)

// Initializes GPIO
func Init() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	output.InitLEDs()
}

// Deactivates all Leds and closes GPIO
func Close() {
	output.SwitchBusyOff()
	output.SwitchFreeOff()
	rpio.Close()
}
