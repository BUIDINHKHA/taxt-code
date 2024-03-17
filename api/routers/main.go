package routers

import (
	"megabot/config"
	"megabot/pkg/logger"
	"megabot/routers/ui"
	v1 "megabot/routers/v1"
	"megabot/usecase"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(
	log *logger.Logger,
	cfg *config.Environment,
	config *usecase.Config,
) *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(cfg.CorsAllowOrigins, ","),
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Access-Control-Allow-Headers",
			"Authorization",
			"X-XSRF-TOKEN",
			"screenId",
			"apiOrder",
		},
		ExposeHeaders: []string{
			"Content-Disposition",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.LoadHTMLGlob("templates/*")
	uiGroup := router.Group("/ui")
	apiGroup := router.Group("/api")
	ui.UiRouter(
		// TODO: implement middleware
		uiGroup.Group(""),
		log,
		cfg,
	)

	v1.V1Router(
		// TODO: implement middleware
		apiGroup.Group("/demo"),
		log,
		config,
	)
	return router
}
