package api

import (
	"context"
	"dummy-simpers-au/config"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func InitTracer(ctx context.Context, env *config.EnvironmentVariable) (*sdktrace.TracerProvider, error) {
	otel.SetErrorHandler(otel.ErrorHandlerFunc(func(err error) {
		log.Printf("OTEL error: %v", err)
	}))

	// Create an OTLP/gRPC exporter that sends to your Collector
	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(env.Tracer.Address), // Collector’s OTLP gRPC port
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	// stdoutExp, err := stdouttrace.New(
	// 	stdouttrace.WithPrettyPrint(),
	// )
	// if err != nil {
	// 	return nil, err
	// }

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		// sdktrace.WithBatcher(stdoutExp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(env.Tracer.ServiceName),
		)),
	)
	otel.SetTracerProvider(tp)
	log.Printf("Succes init Tracer on : %v - %v", env.Tracer.Address, env.Tracer.ServiceName)

	return tp, nil
}
