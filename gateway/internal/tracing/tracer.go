package tracing

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

func NewTracer(exp oteltrace.SpanExporter, name string) (trace.Tracer, *oteltrace.TracerProvider) {
	tp := newTraceProvider(exp)

	otel.SetTracerProvider(tp)

	t := tp.Tracer(name)

	return t, tp
}

func NewOtlExporter(ctx context.Context, ep string) oteltrace.SpanExporter {
	insecureOpt := otlptracehttp.WithInsecure()

	epOpt := otlptracehttp.WithEndpoint(ep)

	exp, err := otlptracehttp.New(ctx, insecureOpt, epOpt)
	if err != nil {
		panic("failed to init span exporter: " + err.Error())
	}

	return exp
}

func newTraceProvider(exp oteltrace.SpanExporter) *oteltrace.TracerProvider {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("myapp"),
		),
	)

	if err != nil {
		panic("failed to init trace provider: " + err.Error())
	}

	return oteltrace.NewTracerProvider(
		oteltrace.WithBatcher(exp),
		oteltrace.WithResource(r),
	)
}

const (
	CtxContextKey utils.ContextKey = "tracer"
)

func WrapTracer(ctx context.Context, t trace.Tracer) context.Context {
	return context.WithValue(ctx, CtxContextKey, t)
}

func ExtractTracer(ctx context.Context) trace.Tracer {
	return ctx.Value(CtxContextKey).(trace.Tracer)
}
