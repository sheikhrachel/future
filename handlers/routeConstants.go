package handlers

const (
	StatusOK                  = 200
	StatusBadRequest          = 400
	StatusTooManyRequests     = 429
	StatusInternalServerError = 500

	// General

	PathRoot   = "/"
	PathHealth = "/health"

	// Appointments Service

	PathNewAppointment      = "/appointments/new"
	PathNewAppointmentBatch = "/appointments/new/batch"
	PathGetAvailability     = "/appointments/availability"
	PathGetBookings         = "/appointments/bookings/:trainer_id"
)
