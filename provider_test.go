package telemetry

import (
	"context"
	"reflect"
	"testing"

	commonv1 "github.com/omcrgnt/proto/gen/go/common/v1"
	"github.com/omcrgnt/res/unique"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func newTestProvider(t *testing.T) (*Provider, *tracetest.InMemoryExporter) {
	t.Helper()

	exp := tracetest.NewInMemoryExporter()
	tp := sdktrace.NewTracerProvider(sdktrace.WithSyncer(exp))
	return &Provider{tp: tp}, exp
}

func setupUseDefaults(t *testing.T) *unique.Registry {
	t.Helper()
	u := unique.New()
	u.MustAddReplaceable(DefaultTrace())
	return u
}

func wireProvider(t *testing.T, u *unique.Registry) *Provider {
	t.Helper()
	raw, err := u.GetOneByType(reflect.TypeOf(&Provider{}))
	if err != nil {
		t.Fatal(err)
	}
	p, ok := raw.(*Provider)
	if !ok {
		t.Fatalf("expected *Provider, got %T", raw)
	}
	p.Inject(nil)
	return p
}

func TestProvider_InjectRecordsSpan(t *testing.T) {
	p, exp := newTestProvider(t)
	p.Inject(nil)

	_, span := otel.Tracer("test").Start(context.Background(), "op")
	span.End()

	spans := exp.GetSpans()
	if len(spans) != 1 {
		t.Fatalf("spans: got %d want 1", len(spans))
	}
	if spans[0].Name != "op" {
		t.Fatalf("span name: got %q", spans[0].Name)
	}

	if err := p.Close(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func TestConfig_Build(t *testing.T) {
	raw, err := (Config{
		ServiceName: &commonv1.Label{Value: "demo"},
		Host:        &commonv1.Host{Value: "127.0.0.1"},
		Port:        &commonv1.Port{Value: 4318},
		Insecure:    true,
	}).Build()
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := raw.(*Provider); !ok {
		t.Fatalf("got %T", raw)
	}
}

func TestFormatEndpoint(t *testing.T) {
	if got := formatEndpoint("localhost", 4318); got != "localhost:4318" {
		t.Fatalf("got %q", got)
	}
}

func TestRegistry_useDefault(t *testing.T) {
	u := setupUseDefaults(t)
	wireProvider(t, u)

	_, span := otel.Tracer("test").Start(context.Background(), "resolve-default")
	defer span.End()

	if !span.SpanContext().TraceID().IsValid() {
		t.Fatal("expected valid trace id after inject")
	}
}

func TestRegistry_configOverride(t *testing.T) {
	u := setupUseDefaults(t)

	built, err := (Config{
		ServiceName: &commonv1.Label{Value: "demo"},
		Host:        &commonv1.Host{Value: "127.0.0.1"},
		Port:        &commonv1.Port{Value: 4318},
		Insecure:    true,
	}).Build()
	if err != nil {
		t.Fatal(err)
	}
	if err := u.Add(built); err != nil {
		t.Fatal(err)
	}
	built.(*Provider).Inject(nil)

	_, span := otel.Tracer("test").Start(context.Background(), "user-override")
	defer span.End()

	if !span.SpanContext().TraceID().IsValid() {
		t.Fatal("expected valid trace id after user override")
	}
}
