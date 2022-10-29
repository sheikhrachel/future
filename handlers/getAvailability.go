package handlers

import (
	"github.com/gin-gonic/gin"
	timeout "github.com/s-wijaya/gin-timeout"

	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/router"
	"github.com/sheikhrachel/future/api_common/util/errutil"
)

func (h *Handler) GetAvailability(c *gin.Context) {
	timeout.APIWrapper(c, func(c *gin.Context) (int, interface{}) {
		if h.shouldRateLimit(c.ClientIP(), PathGetAvailability) {
			return StatusTooManyRequests, nil
		}
		var request GetAvailabilityRequest
		err := router.UnmarshallRequestBody(c.Request.Body, &request)
		if errutil.HandleError(err) {
			return StatusBadRequest, gin.H{"message": err.Error()}
		}
		cc := call.New()

		// TODO: move start and end validation into the endpoint to early return 400

		// fetch availability with valid params
		availability, err := h.AppointmentService.GetAvailability(cc, request.TrainerId, request.StartsAt, request.EndsAt)
		if errutil.HandleError(err) {
			return StatusInternalServerError, gin.H{"message": err.Error()}
		}
		return StatusOK, gin.H{"trainer_id": request.TrainerId, "upcoming_availability": availability}
	})
}

type GetAvailabilityRequest struct {
	TrainerId int    `json:"trainer_id"`
	StartsAt  string `json:"starts_at"`
	EndsAt    string `json:"ends_at"`
}
