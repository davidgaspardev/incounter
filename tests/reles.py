import sys
import RPi.GPIO as gpio
import threading

def listen_signal_by_pin(pin):
    gpio.setup(pin, gpio.IN)

    count = 0
    last_signal = 0
    while True:
        signal = gpio.input(pin)
        if signal != last_signal:
            last_signal = signal
            print(count, "[ PIN", pin, "] Signal received:", signal)
            count += 1

def setup(pins):
    print("pins getted:", pins)

    gpio.setmode(gpio.BCM)

    #for pin in pins:
    #    listen_signal_by_pin(pin)
    listeners = [threading.Thread(target=listen_signal_by_pin, args=(pin,)) for pin in pins]

    for listener in listeners:
        listener.start()
    
    for listener in listeners:
        listener.join()

def boot():
    argc = len(sys.argv)
    pins = []
    if argc > 1:
        for i in range(1, argc):
            if sys.argv[i].isnumeric():
                pins.append(int(sys.argv[i]))
    if len(pins) == 0:
        return
    
    setup(pins)

boot()