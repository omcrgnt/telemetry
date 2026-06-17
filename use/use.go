// Package use registers telemetry system defaults in res.Default.
//
// Import for side effects at the app composition root (main or a meta use package):
//
//	import _ "github.com/omcrgnt/telemetry/use"
package use

import (
	"github.com/omcrgnt/res"
	"github.com/omcrgnt/telemetry"
)

func init() {
	_ = res.AddWithTags(telemetry.DefaultTrace(), res.TagReplaceable)
}
