package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"github.com/sheikhrachel/future/api_common/ratelimit"
)

var (
	tokensPerSecond       = 10
	maxBurstSize    int64 = 25
)

func registerEndpoints(r *gin.Engine, h *Handler) {
	endpoints := []struct {
		method, path string
		handlerFunc  gin.HandlerFunc
	}{
		{http.MethodGet, PathRoot, h.HealthCheck},
		{http.MethodGet, PathHealth, h.HealthCheck},
		{http.MethodPost, PathNewAppointment, h.NewAppointment},
		{http.MethodPost, PathNewAppointmentBatch, h.NewAppointmentBatch},
		{http.MethodGet, PathGetAvailability, h.GetAvailability},
		{http.MethodGet, PathGetBookings, h.GetBookings},
	}
	for _, endpoint := range endpoints {
		registerEndpoint(r, h, endpoint.method, endpoint.path, endpoint.handlerFunc)
	}
}

func registerEndpoint(
	r *gin.Engine,
	h *Handler,
	method, path string,
	handlerFunc gin.HandlerFunc,
) {
	h.RateLimiters[path] = ratelimit.NewIPRateLimiter(rate.Limit(tokensPerSecond), maxBurstSize)
	switch method {
	case http.MethodGet:
		r.GET(path, handlerFunc)
	case http.MethodPost:
		r.POST(path, handlerFunc)
	}
}
