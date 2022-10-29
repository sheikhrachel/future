package appointment_service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sheikhrachel/future/api_common/aws/rds"
	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/util/errutil"
	"github.com/sheikhrachel/future/model"
)

func (a *AppointmentService) GetAllBookings(
	cc call.Call,
	trainerId int,
) (bookings []model.Appointment, err error) {
	queryOpts := rds.AppointmentQueryOpts{TrainerId: trainerId}
	bookings, err = a.Postgres.GetAllBookings(cc, queryOpts)
	if errutil.HandleError(err) {
		return bookings, err
	}
	return bookings, nil
}

func (a *AppointmentService) GetBookingsInRange(
	cc call.Call,
	trainerId int,
	start, end time.Time,
) (bookings []time.Time, err error) {
	queryOpts := rds.AppointmentQueryOpts{
		TrainerId: trainerId,
		StartUnix: start.Unix(),
		EndUnix:   end.Unix(),
	}
	bookingStarts, err := a.Postgres.GetBookingStartsInRange(cc, queryOpts)
	if errutil.HandleError(err) {
		return bookings, err
	}
	for _, booking := range bookingStarts {
		bookingStart, err := strconv.ParseInt(booking.StartsAt, 10, 64)
		if errutil.HandleError(err) {
			cc.InfoF(fmt.Sprintf("err: %#v", err))
			continue
		}
		bookings = append(bookings, time.Unix(bookingStart, 0).In(PST))
	}
	return bookings, nil
}
