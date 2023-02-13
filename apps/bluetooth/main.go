package main

// Log to journal - https://github.com/coreos/go-systemd/blob/main/journal/journal_unix.go

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
	"github.com/pkg/errors"
)

var (
	device = flag.String("device", "default", "implementation of ble")

	serviceUUID            = ble.MustParse("2aecbc54-9099-4816-9a8e-5890414d1399")
	echoCharacteristicUUID = ble.MustParse("fab48e37-6afa-4615-b2d6-002b01c6d7a7")
)

func echoCharacteristic() *ble.Characteristic {
	e := ble.NewCharacteristic(echoCharacteristicUUID)
	e.HandleWrite(ble.WriteHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter) {
		fmt.Println(req.Data())

		rsp.SetStatus(ble.ErrSuccess)
	}))

	return e
}

func main() {
	flag.Parse()

	d, err := dev.NewDevice(*device)
	if err != nil {
		log.Fatalf("can't new device : %s", err)
	}
	ble.SetDefaultDevice(d)

	testSvc := ble.NewService(serviceUUID)
	testSvc.AddCharacteristic(echoCharacteristic())

	if err := ble.AddService(testSvc); err != nil {
		log.Fatalf("can't add service: %s", err)
	}

	// Advertise for specified durantion, or until interrupted by user.
	log.Println("Starting")
	ctx := ble.WithSigHandler(context.WithCancel(context.Background()))
	chkErr(ble.AdvertiseNameAndServices(ctx, "Gopher", serviceUUID))
}

func chkErr(err error) {
	switch errors.Cause(err) {
	case nil:
	case context.DeadlineExceeded:
		fmt.Printf("done\n")
	case context.Canceled:
		fmt.Printf("canceled\n")
	default:
		log.Fatalf(err.Error())
	}
}
