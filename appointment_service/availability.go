package appointment_service

import (
	"time"

	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/util/errutil"
)

func (a *AppointmentService) FindAvailabilityInRange(
	cc call.Call,
	trainerId int,
	start, end time.Time,
) (availability []time.Time, err error) {
	bookingsStartsInRange, err := a.GetBookingsInRange(cc, trainerId, start, end)
	if errutil.HandleError(err) {
		return availability, err
	}
	possibleAvailabilityInRange := getValidAvailabilityInRange(start, end)

	return filterBookingsFromPossibleAvailability(bookingsStartsInRange, possibleAvailabilityInRange), nil
}

// TODO: refine loop to skip over nights/weekends to reduce calls to isValidDate
func getValidAvailabilityInRange(start, end time.Time) (availability []time.Time) {
	curr := start
	// end is the cap, and appointments are 30 minutes, so the last available
	// appointment is 30 min before the end of the range
	for curr != end {
		if isValidDate(curr) {
			availability = append(availability, time.Unix(curr.Unix(), 0).In(PST))
		}
		curr = curr.Add(30 * time.Minute)
	}
	return availability
}

// TODO: update GetBookingStartsInRange to handle the bookingStartsMap creation?
func filterBookingsFromPossibleAvailability(
	bookingsStartsInRange, possibleAvailabilityInRange []time.Time,
) (availability []time.Time) {
	bookingStartsMap := map[time.Time]struct{}{}
	for _, bookingStarts := range bookingsStartsInRange {
		bookingStartsMap[bookingStarts] = struct{}{}
	}
	for _, possibleAvailability := range possibleAvailabilityInRange {
		if _, exists := bookingStartsMap[possibleAvailability]; !exists {
			availability = append(availability, possibleAvailability)
		}
	}
	return availability
}
