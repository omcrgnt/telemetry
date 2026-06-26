/*
Package telemetry provides OpenTelemetry trace export (OTLP) with resources in unique.

# Bootstrap

Blank-import the use subpackage at the app composition root (not telemetry itself):

	import _ "github.com/omcrgnt/telemetry/use"

telemetry/use registers [DefaultTrace] via unique.MustAddReplaceable.
Hardcoded dev defaults: host localhost, port 4318, insecure TLS, service name app.

# User override

Build and register a user Provider with unique.Add (replaces the system default):

	built, err := cfg.Telemetry.Build()
	unique.Global().Add(built)

[Config] uses [common.v1.Label], [common.v1.Host], and [common.v1.Port] for OTLP endpoint and service name.

# Wiring

[Provider] implements Inject (sets global TracerProvider) and Close (shutdown).
Call Inject before runner.Run so HTTP otelhttp instrumentation sees the provider.
*/
package telemetry
