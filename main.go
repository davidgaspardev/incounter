package main

import (
	"fmt"
	"incount/src/signal"
)

func main() {
	receiver := signal.GetReceiverAtPin(2)

	var count uint16

	receiver.OnSignalChange(func(signal uint8) {
		fmt.Println(count, "- Signal:", signal)

		count++
	})
}
