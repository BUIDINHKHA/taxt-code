package ui

import (
	"megabot/config"
	"megabot/pkg/logger"

	"github.com/gin-gonic/gin"
)

func UiRouter(
	rg *gin.RouterGroup,
	log *logger.Logger,
	cfg *config.Environment,
) {
	rg.Use()
	{
		initTaxCodeRouter(rg.Group("/search"))
	}
}
