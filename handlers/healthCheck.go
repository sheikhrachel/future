package handlers

import (
	"github.com/gin-gonic/gin"
	timeout "github.com/s-wijaya/gin-timeout"
)

// HealthCheck is used by our deployment validation step to measure application health
func (h *Handler) HealthCheck(c *gin.Context) {
	timeout.APIWrapper(c, func(c *gin.Context) (int, interface{}) {
		return StatusOK, gin.H{"message": "health"}
	})
}
