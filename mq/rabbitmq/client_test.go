package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"testing"
)

func get_client() (*Client, error) {
	return New(&Options{})
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
		Name: "linshan",
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

	delivery, err := client.ConsumerMessage(&Parameters{
		ExchangeName: "test",
		QueueName: "test:queue",
	})

	if err != nil {
		t.Fatal(err)
	}

	for d := range delivery {
		t.Log(d)
		t.Log(d.MessageId)
		t.Log(d.ConsumerTag)
		t.Log(d.MessageCount)
		t.Log(d.Body)
	}
}