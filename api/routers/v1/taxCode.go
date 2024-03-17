package v1

import (
	"megabot/handler/api"
	"megabot/usecase"

	"github.com/gin-gonic/gin"
)

func initTaxCodeRouter(
	rg gin.IRouter,
	cfg *usecase.Config,
) {
	rg.POST("/search/:taxCode", api.Search(cfg))
}
