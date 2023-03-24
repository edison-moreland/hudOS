package main

import (
	"time"
)

// Small daemon to dim the screen based on proximity/acceleration
// Device driver: https://github.com/torvalds/linux/blob/master/drivers/iio/light/stk3310.c

// These thresholds are totally arbitrary and might need to be tweaked in the future
const (
	proximityPath      = "/sys/bus/iio/devices/iio:device2/in_proximity_raw"
	proximityThreshold = 255

	accelerometerXPath     = "/sys/bus/iio/devices/iio:device1/in_accel_x_raw"
	accelerometerYPath     = "/sys/bus/iio/devices/iio:device1/in_accel_y_raw"
	accelerometerThreshold = 1000
)

func main() {
	lastBrightness := brightness_on
	setBacklightBrightness(lastBrightness)

	ticker := time.NewTicker(time.Second / 4)
	for range ticker.C {
		proximity := readSysfsSensor(proximityPath)
		proximityNear := proximity > proximityThreshold

		var inMotion = false
		if !proximityNear {
			accelerationY := readSysfsSensor(accelerometerYPath)
			accelerationX := readSysfsSensor(accelerometerYPath)

			inMotion = (accelerationY > accelerometerThreshold) || (accelerationX > accelerometerThreshold)
		}

		newBrightness := getBrightness(proximityNear, inMotion)
		if newBrightness != lastBrightness {
			setBacklightBrightness(newBrightness)
			lastBrightness = newBrightness
		}
	}
}
