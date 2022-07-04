package main

import (
	"fmt"
	"incount/src/signal"

	"github.com/davidgaspardev/golog"
)

func main() {
	unblock := make(chan bool)

	gpio2 := uint8(2)
	gpio17 := uint8(17)
	gpio27 := uint8(27)

	receiver2 := signal.GetReceiverAtPin(gpio2)
	receiver17 := signal.GetReceiverAtPin(gpio17)
	receiver27 := signal.GetReceiverAtPin(gpio27)

	var count2 uint16
	receiver2.OnSignalChange(func(signal uint8) {
		count2++
		golog.Info(fmt.Sprintf("PIN %02d", gpio2), fmt.Sprintf("%2dº signal: %d", count2, signal))
	})

	var count17 uint16
	receiver17.OnSignalChange(func(signal uint8) {
		count17++
		golog.Info(fmt.Sprintf("PIN %02d", gpio17), fmt.Sprintf("%2dº signal: %d", count17, signal))
	})

	var count27 uint16
	receiver27.OnSignalChange(func(signal uint8) {
		count27++
		golog.Info(fmt.Sprintf("PIN %02d", gpio27), fmt.Sprintf("%2dº signal: %d", count27, signal))
	})

	for {
		if <-unblock {
			// Finish program
			break
		}
	}
}
