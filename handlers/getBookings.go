package handlers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	timeout "github.com/s-wijaya/gin-timeout"

	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/util/errutil"
	"github.com/sheikhrachel/future/appointment_service"
	"github.com/sheikhrachel/future/model"
)

func (h *Handler) GetBookings(c *gin.Context) {
	timeout.APIWrapper(c, func(c *gin.Context) (int, interface{}) {
		if h.shouldRateLimit(c.ClientIP(), PathGetBookings) {
			return StatusTooManyRequests, nil
		}
		trainerId, err := getTrainerId(c)
		if errutil.HandleError(err) {
			return StatusBadRequest, gin.H{"message": err}
		}

		cc := call.New()

		bookings, err := h.AppointmentService.GetBookings(cc, trainerId)
		if errutil.HandleError(err) {
			return StatusInternalServerError, gin.H{"message": err}
		}
		if bookings != nil {
			bookings = fixOutputBookingDates(bookings)
		}
		return StatusOK, gin.H{"trainer_id": trainerId, "bookings": bookings}
	})
}

func getTrainerId(c *gin.Context) (trainerId int, err error) {
	trainerIdString := c.Param("trainer_id")
	trainerIdi64, err := strconv.ParseInt(trainerIdString, 10, 64)
	if errutil.HandleError(err) {
		return 0, err
	}
	return int(trainerIdi64), nil
}

func fixOutputBookingDates(bookings []model.Appointment) (fixedBookings []model.Appointment) {
	for b, booking := range bookings {
		bookingStartInt, _ := strconv.ParseInt(booking.StartsAt, 10, 64)
		bookingStart := time.Unix(bookingStartInt, 0)
		bookingEndInt, _ := strconv.ParseInt(booking.EndsAt, 10, 64)
		bookingEnd := time.Unix(bookingEndInt, 0)

		bookings[b].StartsAt = time.Unix(bookingStart.Unix(), 0).In(appointment_service.PST).String()
		bookings[b].EndsAt = time.Unix(bookingEnd.Unix(), 0).In(appointment_service.PST).String()
	}
	return bookings
}
