import sys
import RPi.GPIO as gpio
import threading
import time

def listen_signal_by_pin(pin):
    gpio.setup(pin, gpio.IN)

    count = 0
    last_signal = 0
    while True:
        signal = gpio.input(pin)
        if signal != last_signal:
            last_signal = signal

            # Log
            now = time.strftime("%Y-%m-%d %H:%M:%S")
            count += 1
            print(f"[ {now} | PIN {pin:02d} ] INFO - {count}º signal: {signal}")

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