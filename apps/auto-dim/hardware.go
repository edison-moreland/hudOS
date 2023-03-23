package main

import (
	"os"
	"strconv"
)

const backlightPath = "/sys/class/backlight/backlight/brightness"

func setBacklightBrightness(value brightness) {
	backlight, err := os.OpenFile(backlightPath, os.O_WRONLY, 0)
	if err != nil {
		panic(err)
	}
	defer backlight.Close()

	_, err = backlight.WriteString(strconv.Itoa(int(value)))
	if err != nil {
		panic(err)
	}
}

func readSysfsSensor(sensorPath string) int {
	sensor, err := os.Open(sensorPath)
	if err != nil {
		panic(err)
	}
	defer sensor.Close()

	var readBuffer = make([]byte, 8)
	b, err := sensor.Read(readBuffer)
	if err != nil {
		panic(err)
	}

	val, err := strconv.Atoi(string(readBuffer[:b-1])) // -1 to strip newline
	if err != nil {
		panic(err)
	}

	return val
}
