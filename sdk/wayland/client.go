package wayland

import (
	"github.com/edison-moreland/nreal-hud/sdk/log"
	"github.com/edison-moreland/nreal-hud/sdk/wayland/protocol"
	"io"
	"net"
)

type Client struct {
	Registry *protocol.ObjectRegistry
	conn     net.Conn
}

func NewClient() (client *Client, err error) {
	client = &Client{}
	client.Registry = protocol.NewRegistry()

	client.conn, err = OpenSocket()
	if err != nil {
		return nil, err
	}

	events := client.startEventPipe()
	client.startEventDispatch(events)

	return
}

func (c *Client) Close() error {
	return c.Close()
}

func (c *Client) GetDisplay() protocol.WlDisplay {
	proxy, err := c.Registry.GetProxy(1)
	if err != nil {
		log.Panic().
			Err(err).
			Msg("WLDisplay proxy should always exist")
	}

	return protocol.NewWlDisplay(proxy)
}

func (c *Client) SendMessage(message *protocol.Message) error {
	rawMessage, err := message.MarshalBinary()
	if err != nil {
		return err
	}

	log.Debug().Msg("Writing test message!")
	bytesWritten, err := c.conn.Write(rawMessage)
	if err != nil {
		return err
	}

	if bytesWritten != len(rawMessage) {
		log.Panic().
			Msg(":(")
	}

	return nil
}

func (c *Client) startEventPipe() <-chan protocol.Message {
	eventOut := make(chan protocol.Message, 5)

	go func() {
		defer log.PanicHandler()
		defer log.Debug().Msg("Event pipe done.")
		log.Debug().Msg("Starting event pipe.")

		var eventHeader [8]byte
		eventBuffer := make([]byte, 500) // Totally guessing
		for {

			// TODO: Can we get this down to one read?
			_, err := io.ReadFull(c.conn, eventHeader[:])
			if err != nil {
				log.Error().
					Err(err).
					Msg("Error reading header from server, bailing!")
				break
			}

			var header protocol.MessageHeader
			_ = header.UnmarshallBinary(eventHeader[:])

			// TODO: Set an upper bound on this
			if cap(eventBuffer) < int(header.MessageLength) {
				eventBuffer = make([]byte, int(header.MessageLength))
			}

			_, err = io.ReadFull(c.conn, eventBuffer[:header.MessageLength])
			if err != nil {
				log.Error().
					Err(err).
					Msg("Error reading message from server, bailing!")
				break
			}

			message := protocol.Message{
				MessageHeader: header,
				Data:          eventBuffer[:header.MessageLength],
			}

			log.Debug().
				Hex("header", eventHeader[:]).
				Hex("body", eventBuffer[:header.MessageLength]).
				Msg("Raw event")

			log.Info().
				Uint32("object_id", header.ObjectID).
				Uint16("opcode", header.Opcode).
				Uint16("body_length", header.MessageLength).
				Msg("Emitting event")

			eventOut <- message
		}
	}()

	return eventOut
}

func (c *Client) startEventDispatch(events <-chan protocol.Message) {
	go func() {
		defer log.PanicHandler()
		defer log.Debug().Msg("Event dispatch done.")
		log.Debug().Msg("Starting event dispatch")

		for event := range events {
			log.Info().
				Uint32("object_id", event.ObjectID).
				Uint16("opcode", event.Opcode).
				Msg("Dispatching event")

			d, err := c.Registry.GetDispatch(event.ObjectID)
			if err != nil {
				log.Warn().
					Err(err).
					Uint32("object_id", event.ObjectID).
					Msg("Unknown object id!")
				continue
			}

			err = d.Dispatch(event.MessageHeader, event.Data)
			if err != nil {
				log.Warn().
					Err(err).
					Uint32("object_id", event.ObjectID).
					Uint16("opcode", event.Opcode).
					Msg("Could not dispatch event!")
				continue
			}
		}
	}()
}
