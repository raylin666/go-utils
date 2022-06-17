package trace

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	TracePackageName = "github.com/raylin666/go-utils/trace"
)

type Tracer struct {
	provider *tracesdk.TracerProvider
}

func NewTracerProvider(exporter tracesdk.SpanExporter, attrs ...attribute.KeyValue) *Tracer {
	var tracer = new(Tracer)
	bsp := tracesdk.NewBatchSpanProcessor(exporter)
	tracer.provider = tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithSpanProcessor(bsp),
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
	otel.SetTracerProvider(tracer.provider)
	return tracer
}

func (t *Tracer) RegisterSpanProcessor(s tracesdk.SpanProcessor) {
	t.provider.RegisterSpanProcessor(s)
}

func (t *Tracer) UnregisterSpanProcessor(s tracesdk.SpanProcessor) {
	t.provider.UnregisterSpanProcessor(s)
}

// SchedulerRootStart 启动一个 Root 链路追踪器
func (t *Tracer) SchedulerRootStart(ctx context.Context, spanName string) (context.Context, trace.Span) {
	tr := t.provider.Tracer(TracePackageName)
	rootCtx, rootSpan := tr.Start(ctx, spanName)
	return rootCtx, rootSpan
}

func (t *Tracer) ForceFlush(ctx context.Context) error {
	return t.provider.ForceFlush(ctx)
}

func (t *Tracer) Shutdown(ctx context.Context) error {
	return t.provider.Shutdown(ctx)
}


