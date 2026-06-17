package telemetry

const (
	defaultHostValue        = "localhost"
	defaultPortValue        = uint32(4318)
	defaultServiceNameValue = "app"
	defaultInsecureValue    = true
)

// DefaultTrace returns the system trace Provider for telemetry/use registration.
func DefaultTrace() any {
	p, err := newProvider(defaultHostValue, defaultPortValue, defaultInsecureValue, defaultServiceNameValue)
	if err != nil {
		return &Provider{tp: noopTracerProvider()}
	}
	return p
}
