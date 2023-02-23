package wayland

import (
	"github.com/edison-moreland/nreal-hud/sdk/log"
	"github.com/edison-moreland/nreal-hud/sdk/wayland/protocol"
	"io"
	"net"
)

type Client struct {
	registry *protocol.ObjectRegistry
	conn     net.Conn
}

func NewClient() (client *Client, err error) {
	client.registry = protocol.NewRegistry()

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
	var display protocol.WlDisplay
}

func (c *Client) SendMessage(message *protocol.Message) error {
	rawMessage, err := message.MarshalBinary()
	if err != nil {
		return err
	}

	log.Log().Msg("Writing test message!")
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
	eventOut := make(chan protocol.Message)

	go func() {
		log.Log().Msg("Starting event pipe")

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

			eventOut <- message
		}
	}()

	return eventOut
}

func (c *Client) startEventDispatch(events <-chan protocol.Message) {

}
