package telemetry

import (
	commonv1 "github.com/omcrgnt/proto/gen/go/common/v1"
)

// Insecure controls whether the OTLP HTTP exporter skips TLS — a plain bool
// doesn't implement Validator, which ecfg requires of every non-proto leaf
// field (see ecfgtool's validateLeaf); a named bool with Usage/Validate is
// the same fix every other non-proto scalar config field in this project
// already uses (e.g. IdleMin int in user-session's session.Spec).
type Insecure bool

func (Insecure) Usage() string {
	return "Skip TLS for the OTLP exporter (true for local/dev collectors, e.g. Jaeger)"
}

func (Insecure) Validate() error {
	return nil
}

// Config configures OTLP trace export (host, port, service name).
// ecfg fills and protovalidates proto fields before Build is called.
type Config struct {
	ServiceName *commonv1.Label `ecfg:"SERVICE_NAME"`
	Host        *commonv1.Host  `ecfg:"HOST"`
	Port        *commonv1.Port  `ecfg:"PORT"`
	Insecure    Insecure        `ecfg:"INSECURE"`
}

// Build returns a Provider resource for unique.Add.
func (c Config) Build() (any, error) {
	return newProvider(
		c.Host.GetValue(),
		c.Port.GetValue(),
		bool(c.Insecure),
		c.ServiceName.GetValue(),
	)
}
