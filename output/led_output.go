package output

import (
	"github.com/stianeikeland/go-rpio"
)

var (
	busyPin = rpio.Pin(10)
	freePin = rpio.Pin(9)
)

// Initializes the LED output pins
func InitLEDs() {
	busyPin.Output()
	freePin.Output()
}

// Switches GPIO10 on
func SwitchBusyOn() {
	busyPin.High()
}

// Switches GPIO10 off
func SwitchBusyOff() {
	busyPin.Low()
}

// Switches GPIO09 on
func SwitchFreeOn() {
	freePin.High()
}

// Switches GPIO09 off
func SwitchFreeOff() {
	freePin.Low()
}
