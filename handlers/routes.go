package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/sheikhrachel/future/api_common/call"
)

func SetRoutes(router *gin.Engine, cc call.Call) {
	handlers := NewHandler(cc)
	registerEndpoints(router, handlers)
}
