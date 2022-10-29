package errutil

import "errors"

var (
	ErrInvalidStart         = errors.New("invalid start received")
	ErrDuplicateAppointment = errors.New("appointment id already exists")
	ErrTrainerAlreadyBooked = errors.New("trainer already booked at this start time")
)
