package handler

import "github.com/evodicka/goccupied-client/output"

// Switches the state of the LEDs, based on the occupied flag
func SwitchState(occupied bool) {
	if occupied {
		output.SwitchBusyOn()
		output.SwitchFreeOff()
	} else {
		output.SwitchFreeOn()
		output.SwitchBusyOff()
	}
}
