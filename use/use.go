// Package use registers telemetry system defaults in unique.Global.
//
// Import for side effects at the app composition root:
//
//	import _ "github.com/omcrgnt/telemetry/use"
package use

import (
	"github.com/omcrgnt/res/unique"
	"github.com/omcrgnt/telemetry"
)

func init() {
	unique.MustAddReplaceable(telemetry.DefaultTrace())
}
