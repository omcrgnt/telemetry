// Package use registers telemetry system defaults in res.Global.
//
// Import for side effects at the app composition root:
//
//	import _ "github.com/omcrgnt/telemetry/use"
package use

import (
	"github.com/omcrgnt/res"
	"github.com/omcrgnt/telemetry"
)

func init() {
	_ = res.AddToGlobalWithTags(telemetry.DefaultTraceConfig(), res.TagReplaceable)
}
