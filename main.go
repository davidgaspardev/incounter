package main

import (
	"fmt"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	fmt.Println("Hello World!")

	err := rpio.Open()
	if err != nil {
		panic(err)
	}

	defer rpio.Close()

	pin := rpio.Pin(2)
	pin.Input()

	fmt.Println("[ PIN 2 ] Read:", pin.Read())
}
