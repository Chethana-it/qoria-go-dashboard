package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHealthRoutes registers health-check routes
func RegisterHealthRoutes(rg *gin.RouterGroup) {
	rg.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
}
