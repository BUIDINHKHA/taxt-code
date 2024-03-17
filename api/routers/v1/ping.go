package v1

import (
	"megabot/handler/api"

	"github.com/gin-gonic/gin"
)

func initPingRouter(
	rg gin.IRouter,
) {
	rg.GET("", func(context *gin.Context) {
		api.Ping(context)
	})
}
