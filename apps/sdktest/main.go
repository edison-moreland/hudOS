package main

import (
	"encoding/binary"
	"github.com/edison-moreland/nreal-hud/sdk/log"
	"github.com/edison-moreland/nreal-hud/sdk/wayland"
	"github.com/edison-moreland/nreal-hud/sdk/wayland/protocol"
	"os"
	"os/signal"
)

func main() {
	client, err := wayland.NewClient()
	if err != nil {
		log.Panic().
			Err(err).
			Msg("Could not open wayland socket!")
	}
	defer client.Close()

	body := make([]byte, 4)
	binary.LittleEndian.PutUint32(body, 2)

	header := protocol.MessageHeader{
		ObjectID:      1,
		MessageLength: 4,
		Opcode:        protocol.WlDisplayGetRegistryRequestOp,
	}
	message := protocol.Message{
		MessageHeader: header,
		Data:          body,
	}

	if err := client.SendMessage(&message); err != nil {
		log.Panic().
			Err(err).
			Msg("Shit")
	}

	blockUntilSig()
}

func blockUntilSig() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	<-c
}
