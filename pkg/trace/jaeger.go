package trace

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	oteljaeger "go.opentelemetry.io/otel/exporters/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

var _ Trace = (*jaeger)(nil)

type jaeger struct {
	*jaegerOption

	*sdktrace.TracerProvider
}

func NewJaeger(endpointOption oteljaeger.EndpointOption, opts ...JaegerOption) (Trace, error) {
	var opt = new(jaegerOption)
	for _, f := range opts {
		f(opt)
	}

	// Create the Jaeger exporter
	// 可以直接连 collector 也可以连 agent
	/**
	agent 连接方式.	jaeger.WithAgentEndpoint(
						jaeger.WithAgentHost("127.0.0.1"),
						jaeger.WithAgentPort("6831"),
					)

	collector 连接方式.	jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url))
	*/
	exporter, err := oteljaeger.New(endpointOption)
	if err != nil {
		return nil, err
	}

	opt.attributes = append(
		opt.attributes,
		semconv.ServiceNameKey.String(opt.serviceName),
		semconv.ServiceVersionKey.String(opt.serviceVersion),
		semconv.ServiceNamespaceKey.String(opt.serviceNamespace),
		semconv.DeploymentEnvironmentKey.String(opt.environment),
		semconv.ServiceInstanceIDKey.String(opt.id),
	)

	return &jaeger{
		jaegerOption: opt,
		TracerProvider: provider(exporter, opt.attributes...),
	}, nil
}

func (j *jaeger) GetTracerProvider() *sdktrace.TracerProvider {
	return j.TracerProvider
}

func (j *jaeger) Shutdown(ctx context.Context) error {
	return j.TracerProvider.Shutdown(ctx)
}

type JaegerOption func(*jaegerOption)

type jaegerOption struct {
	serviceName      string
	serviceVersion   string
	serviceNamespace string
	environment      string
	id               string

	attributes []attribute.KeyValue
}

func WithJaegerServiceName(name string) JaegerOption {
	return func(option *jaegerOption) {
		option.serviceName = name
	}
}

func WithJaegerServiceVersion(version string) JaegerOption {
	return func(option *jaegerOption) {
		option.serviceVersion = version
	}
}

func WithJaegerServiceNamespace(namespace string) JaegerOption {
	return func(option *jaegerOption) {
		option.serviceNamespace = namespace
	}
}

func WithJaegerEnvironment(environment string) JaegerOption {
	return func(option *jaegerOption) {
		option.environment = environment
	}
}

func WithJaegerID(id string) JaegerOption {
	return func(option *jaegerOption) {
		option.id = id
	}
}

func WithJaegerAttributes(attrs ...attribute.KeyValue) JaegerOption {
	return func(option *jaegerOption) {
		option.attributes = attrs
	}
}
