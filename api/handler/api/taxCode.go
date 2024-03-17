package api

import (
	"errors"
	"megabot/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(cfg *usecase.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		taxCode := c.Param("taxCode")
		if len(taxCode) == 0 {
			c.JSON(http.StatusBadRequest, errors.New("missing tax code"))
			return
		}
		output, err := cfg.Search(taxCode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
