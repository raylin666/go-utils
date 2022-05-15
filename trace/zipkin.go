package trace

import (
	exporterZipkin "go.opentelemetry.io/otel/exporters/zipkin"
)

type Zipkin struct {
	exporter *exporterZipkin.Exporter
}

func NewZipkin(collectorURL string) (*Zipkin, error) {
	var zipkin = new(Zipkin)
	exporter, err := exporterZipkin.New(collectorURL)
	if err != nil {
		return nil, err
	}

	zipkin.exporter = exporter
	return zipkin, nil
}

func (z *Zipkin) Exporter() *exporterZipkin.Exporter {
	return z.exporter
}
