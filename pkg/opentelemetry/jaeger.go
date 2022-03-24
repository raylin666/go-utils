package opentelemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func NewJaeger(url, serviceName, serviceVersion, environment, id string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	// 可以直接连 collector 也可以连 agent
	exporter, err := jaeger.New(
		jaeger.WithAgentEndpoint(
			//jaeger.WithAgentHost("127.0.0.1"),
			jaeger.WithAgentPort("6831"),
		),
		// jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)),
	)
	if err != nil {
		return nil, err
	}

	tp, err := tracerProvider(serviceName, serviceVersion, environment, id, exporter)
	if err != nil {
		return nil, err
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

