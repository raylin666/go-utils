package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"testing"
)

func get_client() (*Client, error) {
	var config amqp.Config
	config.Vhost = "test"
	return New(&Options{
		Config: config,
	})
}

func TestConnection(t *testing.T) {
	client, err := get_client()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(client)
}

func TestPublishMessage(t *testing.T) {
	client, err := get_client()
	if err != nil {
		t.Fatal(err)
	}

	data, err := json.Marshal(struct {
		Name string
	}{
		Name: "yangtiantian",
	})

	var published amqp.Publishing
	published.Body = data
	queue, err := client.PublishMessage(&Parameters{
		ExchangeName: "test",
		QueueName: "test:queue",
	}, Publishing{published})

	if err != nil {
		t.Fatal(err)
	}

	t.Log("Publish Message Success.", queue)
}

func TestConsumeMessage(t *testing.T) {
	client, err := get_client()
	if err != nil {
		t.Fatal(err)
	}

	err = client.ConsumerMessage(&Parameters{
		ExchangeName: "test",
		QueueName: "test:queue",
	}, func(c *Client, ch *Channel, delivery amqp.Delivery) {
		t.Log(string(delivery.Body))
		// ch.Ack(delivery.DeliveryTag, false)
	}, true)

	if err != nil {
		t.Fatal(err)
	}
}