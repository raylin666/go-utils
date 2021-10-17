package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Client struct {
	Conn *amqp.Connection
}

func (c *Client) Channel() (*Channel, error) {
	var ch = new(Channel)
	channel, err := c.Conn.Channel()
	if err != nil {
		return nil, err
	}

	ch.Channel = channel
	return ch, nil
}

func (c *Client) Close() error {
	return c.Conn.Close()
}


