package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	c.HTML(http.StatusOK, "search.html", gin.H{})
}
