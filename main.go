package main

import (
	"github.com/gin-gonic/gin"

	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/router"
	"github.com/sheikhrachel/future/handlers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	router.SetCorsOnRouter(r)
	router.SetupTimeoutMiddleware(r)
	cc := call.New()
	handlers.SetRoutes(r, cc)
	router.StartRouter(r, cc)
}
