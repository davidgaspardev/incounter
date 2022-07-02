package signal

import "github.com/stianeikeland/go-rpio"

func init() {
	// Start Raspberry GPIO
	if err := rpio.Open(); err != nil {
		panic(err)
	}
}
