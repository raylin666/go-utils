package trace

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

const TracerName = "github.com/raylin666/go-utils/pkg/trace"

type Trace interface {
	GetTracerProvider() *tracesdk.TracerProvider
	Shutdown(ctx context.Context) error
}

func provider(exporter tracesdk.SpanExporter, attrs ...attribute.KeyValue) *tracesdk.TracerProvider {
	var tracerProvider = tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exporter),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			attrs...
		)),
	)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tracerProvider)
	return tracerProvider
}


