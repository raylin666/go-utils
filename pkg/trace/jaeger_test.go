package trace

import (
	oteljaeger "go.opentelemetry.io/otel/exporters/jaeger"
	"testing"
)

func TestJaeger(t *testing.T) {
	trace, err := NewJaeger(
		oteljaeger.WithAgentEndpoint(
			oteljaeger.WithAgentHost("127.0.0.1"),
			oteljaeger.WithAgentPort("6831"),
		),
		WithJaegerServiceName("test.service"))
	if err != nil {
		t.Fatal(err)
	}


	t.Log(trace)
}