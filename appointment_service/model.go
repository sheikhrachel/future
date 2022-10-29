package appointment_service

import "time"

var (
	weekdays = map[string]struct{}{
		"Monday":    {},
		"Tuesday":   {},
		"Wednesday": {},
		"Thursday":  {},
		"Friday":    {},
	}

	PST, _ = time.LoadLocation("America/Los_Angeles")
)
