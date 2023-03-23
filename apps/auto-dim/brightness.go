package main

import "time"

type brightness int

const (
	brightness_on  brightness = 2048
	brightness_dim brightness = 2048 / 2
	brightness_off brightness = 0
)

const (
	dimAfter = time.Minute / 2
	offAfter = time.Minute
)

var lastMotion = time.Now()

func getBrightness(proximityNear, inMotion bool) brightness {
	if proximityNear {
		return brightness_off
	}

	if inMotion {
		lastMotion = time.Now()
		return brightness_on
	}
	timeSinceMotion := time.Since(lastMotion)

	if timeSinceMotion < dimAfter {
		return brightness_on
	}

	if timeSinceMotion < offAfter {
		return brightness_dim
	}

	return brightness_off
}
