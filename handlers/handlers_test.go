package handlers

import (
	"bufio"
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"

	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/router"
	`github.com/sheikhrachel/future/api_common/util/errutil`
	"github.com/sheikhrachel/future/model"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	router.SetCorsOnRouter(r)
	router.SetupTimeoutMiddleware(r)
	cc := call.New()
	handlers := NewHandler(cc)
	registerEndpoints(r, handlers)
	return r
}

func sendTestRequest(r *gin.Engine, method, url string, bodyBytes []byte) (w *httptest.ResponseRecorder) {
	reader := bytes.NewReader(bodyBytes)
	body := bufio.NewReader(reader)
	req, _ := http.NewRequest(method, url, body)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var (
	appointment0 = createAppointmentModel(0, 1, 1, "2019-01-24T10:00:00-08:00", "2019-01-24T10:30:00-08:00")
	appointment1 = createAppointmentModel(1, 2, 1, "2019-01-24T11:00:00-08:00", "2019-01-24T11:30:00-08:00")
	appointment2 = createAppointmentModel(2, 3, 1, "2019-01-24T12:30:00-08:00", "2019-01-24T13:00:00-08:00")
)

func (h *Handler) bookAppointments() {
	cc := call.New()
	for _, appt := range []model.Appointment{appointment0, appointment1, appointment2} {
		err := h.AppointmentService.AddNewAppointment(cc, appt)
		if errutil.HandleError(err) {
			// panic is acceptable here for test-only method
			panic(err)
		}
	}
}

func (h *Handler) cleanupAppointments() {
	cc := call.New()
	for _, appt := range []model.Appointment{appointment0, appointment1, appointment2} {
		err := h.AppointmentService.DeleteAppointment(cc, appt.AppointmentId)
		if errutil.HandleError(err) {
			// panic is acceptable here for test-only method
			panic(err)
		}
	}
}

func createAppointmentModel(id, userId, trainerId int, starts, ends string) model.Appointment {
	return model.Appointment{
		AppointmentId: id,
		UserId:        userId,
		TrainerId:     trainerId,
		StartsAt:      starts,
		EndsAt:        ends,
	}
}
