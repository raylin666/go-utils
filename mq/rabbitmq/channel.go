package rabbitmq

import "github.com/streadway/amqp"

type Channel struct {
	*amqp.Channel
}

func (ch *Channel) QueueDeclare(parameters *Parameters) (amqp.Queue, error) {
	return ch.Channel.QueueDeclare(
		parameters.QueueName,
		parameters.Durable,
		parameters.AutoDelete,
		parameters.Exclusive,
		parameters.NoWait,
		parameters.Args)
}

func (ch *Channel) QueueBind(parameters *Parameters) error {
	if parameters.QueueKey == "" {
		parameters.QueueKey = parameters.QueueName
	}

	return ch.Channel.QueueBind(
		parameters.QueueName,
		parameters.QueueKey,
		parameters.ExchangeName,
		parameters.NoWait,
		parameters.Args)
}

func (ch *Channel) ExchangeDeclare(parameters *Parameters) error {
	if parameters.ExchangeType == "" {
		parameters.ExchangeType = "direct"
	}

	return ch.Channel.ExchangeDeclare(
		parameters.ExchangeName,
		parameters.ExchangeType,
		parameters.Durable,
		parameters.AutoDelete,
		parameters.Internal,
		parameters.NoWait,
		parameters.Args)
}

func (ch *Channel) Publish(parameters *Parameters, publishing Publishing) error {
	if publishing.ContentType == "" {
		publishing.ContentType = "text/plain"
	}

	return ch.Channel.Publish(
		parameters.ExchangeName,
		parameters.QueueKey,
		parameters.Mandatory,
		parameters.Immediate,
		publishing.Publishing)
}

func (ch *Channel) Consume(parameters *Parameters) (<-chan amqp.Delivery, error) {
	return ch.Channel.Consume(
		parameters.QueueName,
		parameters.ConsumerTag,
		parameters.AutoAck,
		parameters.Exclusive,
		parameters.NoLocal,
		parameters.NoWait,
		parameters.Args)
}

func (ch *Channel) Close() error {
	return ch.Channel.Close()
}
