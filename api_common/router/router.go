package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	timeout "github.com/s-wijaya/gin-timeout"

	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/util/errutil"
)

const (
	timeoutVal        = 10 * time.Second
	errRequestTimeout = "request timeout"
)

var (
	corsConfig = cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}
)

func StartRouter(r *gin.Engine, cc call.Call) {
	cc.InfoF("router starting on 8080")
	err := r.Run()
	if errutil.HandleError(err) {
		cc.InfoF(fmt.Sprintf("StartRouter failed to start server with err: %+v", err.Error()))
	} // listen and serve on 0.0.0.0:8080
}

func SetCorsOnRouter(r *gin.Engine) {
	r.Use(cors.New(corsConfig))
}

func SetupTimeoutMiddleware(r *gin.Engine) {
	timeoutMsg := gin.H{"message": errRequestTimeout}
	r.Use(timeout.TimeoutHandler(timeoutVal, http.StatusRequestTimeout, timeoutMsg))
}
