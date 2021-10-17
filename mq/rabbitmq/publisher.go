package rabbitmq

import "github.com/streadway/amqp"

type Publishing struct {
	amqp.Publishing
}

func (c *Client) PublishMessage(parameters *Parameters, publishing Publishing) (*amqp.Queue, error) {
	ch, err := c.Channel()
	if err != nil {
		return nil, err
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(parameters)
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

	return &queue, ch.Publish(parameters, publishing)
}
