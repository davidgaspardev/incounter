package signal

import (
	"time"

	"github.com/stianeikeland/go-rpio"
)

// Receivers are responsible for getting the input signals
// from a certain pin (GPIO).
type Receiver struct {
	id  string
	pin rpio.Pin
}

// Receivers collections
var receivers = make(map[uint8]*Receiver)

func GetReceiverAtPin(pin uint8) *Receiver {
	if receivers[pin] == nil {
		// Register pin
		receivers[pin] = &Receiver{
			pin: rpio.Pin(pin),
		}

		// Set pin to signal input
		receivers[pin].pin.Input()
	}

	return receivers[pin]
}

func (receiver *Receiver) OnSignalChange(action func(signal uint8)) {
	go func() {
		var lastSignal rpio.State // Start with 0

		for {
			signal := receiver.pin.Read()

			if signal != lastSignal {
				lastSignal = signal

				action(uint8(signal))
			}

			time.Sleep(750 * time.Microsecond)
		}
	}()
}
