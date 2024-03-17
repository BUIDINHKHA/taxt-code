package v1

import (
	"github.com/gin-gonic/gin"
	"megabot/pkg/logger"
	"megabot/usecase"
)

func V1Router(
	rg *gin.RouterGroup,
	log *logger.Logger,
	config *usecase.Config,
) {
	rg.Use()
	{
		initPingRouter(rg.Group("/ping"))
		initTaxCodeRouter(rg.Group("/tax_code"), config)
	}
}
