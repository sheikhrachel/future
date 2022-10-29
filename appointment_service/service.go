package appointment_service

import (
	"github.com/sheikhrachel/future/model"
)

type Service interface {
	AddNewAppointment(appointmentId int) error
	DeleteAppointment(appointmentId int) error
	AlreadyExists(appointmentId int) (exists bool, err error)
	ValidateStart(start string) (valid bool, err error)
	GetAvailability(trainerId int) ([]string, error)
	GetBookings(trainerId int) ([]model.Appointment, error)
	TrainerExists(trainerId int) (exists bool, err error)
}
