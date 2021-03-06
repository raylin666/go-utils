package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func conn() (Client, error) {
	var opts = new(Options)
	opts.Network = "tcp"
	opts.Addr = "127.0.0.1:6379"
	opts.Password = "123456"
	opts.DB = 0
	return NewClient(context.Background(), opts)
}

func TestKeys(t *testing.T) {
	client, err := conn()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(client.Keys(context.Background(), "*"))
}

func TestPipeline(t *testing.T) {
	client, err := conn()
	if err != nil {
		t.Fatal(err)
	}

	var ctx = context.Background()
	cmd, err := client.TxPipeline(ctx, func(pipe redis.Pipeliner) {
		pipe.Set(context.TODO(), "name", "kaka", time.Hour)
		pipe.Get(context.TODO(), "name")
	})

	t.Log(cmd)

	t.Log(client.Command(ctx))
}