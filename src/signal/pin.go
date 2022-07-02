package signal

import (
	"github.com/stianeikeland/go-rpio"
)

func init() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
}

type Pin struct {
	code   string
	signal rpio.Pin
}

var sensors = make(map[uint8]*Pin)

func GetPin(pinNumber uint8) *Pin {
	if sensors[pinNumber] == nil {
		// Register pin
		sensors[pinNumber] = &Pin{
			signal: rpio.Pin(pinNumber),
		}

		// Set pin to signal input
		sensors[pinNumber].signal.Input()
	}

	return sensors[pinNumber]
}

func (pin *Pin) OnSignalChange(action func(signal uint8)) {
	var lastSignal rpio.State // Start with 0

	for {
		signal := pin.signal.Read()

		if signal != lastSignal {
			lastSignal = signal

			action(uint8(signal))
		}
	}
}
