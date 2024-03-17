package ui

import (
	"megabot/handler/ui"

	"github.com/gin-gonic/gin"
)

func initTaxCodeRouter(
	r gin.IRouter,
) {
	r.GET("", func(context *gin.Context) {
		ui.Search(context)
	})
}
