package appointment_service

import (
	"time"

	"github.com/sheikhrachel/future/api_common/aws/rds"
	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/util/errutil"
	"github.com/sheikhrachel/future/model"
)

func (a *AppointmentService) AddNewAppointment(cc call.Call, appointment model.Appointment) error {
	startTime, endTime, err := getStartAndEndTime(appointment.StartsAt, appointment.EndsAt)
	if errutil.HandleError(err) {
		return err
	}
	queryOpts := rds.AppointmentQueryOpts{
		AppointmentId: appointment.AppointmentId,
		TrainerId:     appointment.TrainerId,
		UserId:        appointment.UserId,
		StartUnix:     startTime.Unix(),
		EndUnix:       endTime.Unix(),
	}
	if alreadyBooked, err := a.Postgres.TrainerAlreadyBooked(cc, queryOpts); alreadyBooked || errutil.HandleError(err) {
		return errutil.ErrTrainerAlreadyBooked
	}
	return a.Postgres.AddNewAppointment(cc, queryOpts)
}

func (a *AppointmentService) DeleteAppointment(cc call.Call, appointmentId int) error {
	queryOpts := rds.AppointmentQueryOpts{AppointmentId: appointmentId}
	return a.Postgres.DeleteAppointment(cc, queryOpts)
}

func (a *AppointmentService) AppointmentExists(cc call.Call, appointmentId int) (exists bool, err error) {
	queryOpts := rds.AppointmentQueryOpts{AppointmentId: appointmentId}
	return a.Postgres.AppointmentExists(cc, queryOpts)
}

func (a *AppointmentService) GetAvailability(cc call.Call, trainerId int, start, end string) ([]time.Time, error) {
	startTime, endTime, err := getStartAndEndTime(start, end)
	if errutil.HandleError(err) {
		return []time.Time{}, err
	}
	return a.FindAvailabilityInRange(cc, trainerId, startTime, endTime)
}

func (a *AppointmentService) GetBookings(cc call.Call, trainerId int) ([]model.Appointment, error) {
	return a.GetAllBookings(cc, trainerId)
}

func (a *AppointmentService) TrainerExists(cc call.Call, trainerId int) (exists bool, err error) {
	queryOpts := rds.AppointmentQueryOpts{TrainerId: trainerId}
	return a.Postgres.TrainerExists(cc, queryOpts)
}

func ValidateStart(start string) (valid bool, err error) {
	dateTime, err := time.Parse(time.RFC3339, start)
	if errutil.HandleError(err) {
		return false, err
	}
	return isValidDate(dateTime), nil
}

func validateDayOfWeek(date time.Time) (valid bool) {
	_, exists := weekdays[date.Weekday().String()]
	return exists
}

func validateHours(date time.Time) (valid bool) {
	// return <= 16, as the last appointment will be at 16:30
	return date.Hour() >= 8 && date.Hour() <= 16
}

func validateMinute(date time.Time) (valid bool) {
	return date.Minute() == 0 || date.Minute() == 30
}

func isValidDate(date time.Time) (valid bool) {
	return validateDayOfWeek(date) &&
		validateHours(date) &&
		validateMinute(date)
}

func getStartAndEndTime(start, end string) (startTime, endTime time.Time, err error) {
	startTime, err = time.Parse(time.RFC3339, start)
	if errutil.HandleError(err) {
		return startTime, endTime, err
	}
	endTime, err = time.Parse(time.RFC3339, end)
	if errutil.HandleError(err) {
		return startTime, endTime, err
	}
	return startTime, endTime, nil
}
