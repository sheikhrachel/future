package handlers

import (
	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/ratelimit"
	"github.com/sheikhrachel/future/appointment_service"
)

type Handler struct {
	AppointmentService *appointment_service.AppointmentService
	RateLimiters       map[string]*ratelimit.IPRateLimiter
}

func NewHandler(cc call.Call) *Handler {
	appointmentSvc := appointment_service.NewAppointmentService(cc)
	return &Handler{
		AppointmentService: appointmentSvc,
		RateLimiters:       make(map[string]*ratelimit.IPRateLimiter),
	}
}
