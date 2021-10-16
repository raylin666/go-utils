package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func conn() (*Client, error) {
	var opts = new(Options)
	opts.Network = "tcp"
	opts.Addr = "127.0.0.1:6379"
	opts.Password = "myredis"
	opts.DB = 0
	return New(context.TODO(), opts)
}

func TestKeys(t *testing.T) {
	client, err := conn()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(client.Keys("*"))
}

func TestPipeline(t *testing.T) {
	client, err := conn()
	if err != nil {
		t.Fatal(err)
	}

	cmd, err := client.TxPipeline(func(pipe redis.Pipeliner) {
		pipe.Set(context.TODO(), "name", "kaka", time.Hour)
		pipe.Get(context.TODO(), "name")
	})

	t.Log(cmd)

	t.Log(client.Command())
}