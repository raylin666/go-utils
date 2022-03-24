package opentelemetry

import (
	"context"
	"testing"
	"time"
)

func TestJaeger(t *testing.T) {
	tp, err := NewJaeger(
		"http://localhost:14268/api/traces",
		"test.service",
		"v1.0.0",
		"dev",
		"001",
	)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second*2)

	_, span := StartFromContext(context.Background(), "tracer", "spanName")
	defer span.End()

	defer func() {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		tp.Shutdown(ctx)
	}()
}

