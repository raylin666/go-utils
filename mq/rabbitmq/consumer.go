package rabbitmq

import (
	"github.com/streadway/amqp"
)

func (c *Client) ConsumerMessage(parameters *Parameters, handler func(c *Client, ch *Channel, delivery amqp.Delivery), isCoroutineHandler bool) error {
	ch, err := c.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	_, err = ch.QueueDeclare(parameters)
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare(parameters)
	if err != nil {
		return err
	}

	err = ch.QueueBind(parameters)
	if err != nil {
		return err
	}

	delivery, err := ch.Consume(parameters)
	if err != nil {
		return err
	}

	for d := range delivery {
		if isCoroutineHandler == true {
			go func(d amqp.Delivery) {
				handler(c, ch, d)
			}(d)
		} else {
			handler(c, ch, d)
		}
	}

	return nil
}