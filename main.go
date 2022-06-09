package main

import (
	"fmt"
	"incount/src/signal"
)

func main() {
	pinSignal := signal.GetPin(2)

	var count uint8

	pinSignal.OnSignalChange(func(signal uint8) {
		fmt.Println(count, "- Signal:", signal)

		count++
	})
}
