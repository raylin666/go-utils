package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"testing"
)

func getClient() (*Client, error) {
	return New(&Options{})
}

func TestConnection(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(client)
}

func TestPublishMessage(t *testing.T) {
	client, err := getClient()
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