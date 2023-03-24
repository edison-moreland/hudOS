package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	porcupine "github.com/Picovoice/porcupine/binding/go"
	"github.com/gen2brain/malgo"
)

func main() {
	samples := make(chan []int16)
	chunkedSamples := make(chan []int16)

	sampleRate, chunkSize, stopeWakeword, err := startWakeWordListener(chunkedSamples)
	if err != nil {
		panic(err)
	}
	defer stopeWakeword()

	startSampleChunker(samples, chunkedSamples, chunkSize)

	stopMicrophone, err := startMicrophoneStream(samples, sampleRate)
	if err != nil {
		panic(err)
	}
	defer stopMicrophone()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	fmt.Println("Shutting down")
}

func startSampleChunker(samplesIn <-chan []int16, chunkedSamplesOut chan<- []int16, chunkSize int) {
	go func() {
		buffer := make([]int16, 0, chunkSize*2)
		for samples := range samplesIn {
			buffer = append(buffer, samples...)
			for len(buffer) >= chunkSize {
				chunk := make([]int16, chunkSize)
				copy(chunk, buffer[:chunkSize])
				chunkedSamplesOut <- chunk
				buffer = buffer[:0]
			}
		}
	}()
}

func startWakeWordListener(chunkedSamplesIn <-chan []int16) (uint32, int, func(), error) {
	porcupineEngine := porcupine.Porcupine{
		BuiltInKeywords: []porcupine.BuiltInKeyword{
			porcupine.COMPUTER,
		},
	}
	if err := porcupineEngine.Init(); err != nil {
		return 0, 0, func() {}, err
	}

	go func() {
		for sample := range chunkedSamplesIn {
			fmt.Println("Processing chunk")
			wakeword, err := porcupineEngine.Process(sample)
			if err != nil {
				panic(err)
			}

			fmt.Println(wakeword)
		}
	}()

	return uint32(porcupine.SampleRate), porcupine.FrameLength, func() {
		porcupineEngine.Delete()
	}, nil
}

func startMicrophoneStream(samplesOut chan<- []int16, sampleRate uint32) (func(), error) {
	malgoCtx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		fmt.Println(message)
	})
	if err != nil {
		return func() {}, err
	}
	closeMalgo := func() {
		malgoCtx.Uninit()
		malgoCtx.Free()
	}

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Capture)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = 1
	deviceConfig.SampleRate = sampleRate

	device, err := malgo.InitDevice(malgoCtx.Context, deviceConfig, malgo.DeviceCallbacks{
		Data: func(_, pInputSamples []byte, framecount uint32) {
			reader := bytes.NewBuffer(pInputSamples)

			convertedSample := make([]int16, framecount/2)
			for i := 0; i < len(convertedSample); i++ {
				err := binary.Read(reader, binary.LittleEndian, &convertedSample[i])
				if err != nil {
					panic(err)
				}
			}

			samplesOut <- convertedSample
		},
	})
	if err != nil {
		closeMalgo()
		return func() {}, err
	}

	device.Start()

	return func() {
		device.Uninit()
		closeMalgo()
	}, nil
}
