package signal

import "github.com/stianeikeland/go-rpio"

type Sender struct {
	id  string
	pin rpio.Pin
}

var senders = make(map[uint8]*Sender)

func GetSenderAtPin(pin uint8) *Sender {
	if senders[pin] == nil {
		senders[pin] = &Sender{
			pin: rpio.Pin(pin),
		}

		senders[pin].pin.Output()
	}

	return senders[pin]
}

func (sender *Sender) SetSignal(signal bool) {
	if signal {
		sender.pin.High()
	} else {
		sender.pin.Low()
	}
}
