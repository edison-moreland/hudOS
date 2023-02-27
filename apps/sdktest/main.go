package main

import (
	"encoding/binary"
	"github.com/edison-moreland/nreal-hud/sdk/log"
	"github.com/edison-moreland/nreal-hud/sdk/wayland"
	"github.com/edison-moreland/nreal-hud/sdk/wayland/protocol"
	"os"
	"os/signal"
)

type registryListener struct {
	protocol.UnimplementedWlRegistryListener
}

func (r *registryListener) Global(e protocol.WlRegistryGlobalEvent) {
	log.Info().
		Uint32("name", e.Name).
		Str("interface", e.Interface).
		Uint32("version", e.Version).
		Msg("Global")
}

func main() {
	defer log.PanicHandler()

	defer log.Info().Msg("sdktest done.")
	log.Info().Msg("sdktest starting!")

	client, err := wayland.NewClient()
	if err != nil {
		log.Panic().
			Err(err).
			Msg("Could not open wayland socket!")
	}
	defer client.Close()

	client.GetDisplay()
	registry := client.Registry.NewWlRegistry()
	registry.AttachListener(&registryListener{})

	body := make([]byte, 4)
	binary.LittleEndian.PutUint32(body, registry.GetId())

	header := protocol.MessageHeader{
		ObjectID:      1,
		MessageLength: 4,
		Opcode:        0,
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
