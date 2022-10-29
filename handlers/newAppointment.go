package handlers

import (
	"github.com/gin-gonic/gin"
	timeout "github.com/s-wijaya/gin-timeout"

	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/router"
	"github.com/sheikhrachel/future/api_common/util/errutil"
	"github.com/sheikhrachel/future/appointment_service"
	"github.com/sheikhrachel/future/model"
)

func (h *Handler) NewAppointment(c *gin.Context) {
	timeout.APIWrapper(c, func(c *gin.Context) (int, interface{}) {
		if h.shouldRateLimit(c.ClientIP(), PathNewAppointment) {
			return StatusTooManyRequests, nil
		}
		var appointment model.Appointment
		err := router.UnmarshallRequestBody(c.Request.Body, &appointment)
		if errutil.HandleError(err) {
			return StatusBadRequest, gin.H{"message": err.Error()}
		}
		validStart, err := appointment_service.ValidateStart(appointment.StartsAt)
		if errutil.HandleError(err) {
			return StatusBadRequest, gin.H{"message": err}
		}
		if !validStart {
			return StatusBadRequest, gin.H{"message": errutil.ErrInvalidStart.Error()}
		}

		cc := call.New()
		exists, err := h.AppointmentService.AppointmentExists(cc, appointment.AppointmentId)
		if errutil.HandleError(err) || exists {
			return StatusBadRequest, gin.H{"message": errutil.ErrDuplicateAppointment.Error()}
		}
		if err = h.AppointmentService.AddNewAppointment(cc, appointment); errutil.HandleError(err) {
			return StatusInternalServerError, gin.H{"message": err.Error()}
		}
		return StatusOK, gin.H{"appointment_added": appointment}
	})
}
