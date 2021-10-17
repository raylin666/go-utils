package rabbitmq

import (
	"github.com/streadway/amqp"
)

func (c *Client) ConsumerMessage(parameters *Parameters) (<-chan amqp.Delivery, error) {
	ch, err := c.Channel()
	if err != nil {
		return nil, err
	}

	defer ch.Close()

	_, err = ch.QueueDeclare(parameters)
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(parameters)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(parameters)
	if err != nil {
		return nil, err
	}

	return ch.Consume(parameters)
}