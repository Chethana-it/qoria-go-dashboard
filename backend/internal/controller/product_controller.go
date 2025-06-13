package controller

import (
	"net/http"
	"strconv"

	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/data"
	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/service"
	"github.com/gin-gonic/gin"
)

// RegisterProductRoutes registers product-related endpoints
func RegisterProductRoutes(rg *gin.RouterGroup, records []data.Record) {
	rg.GET("/top-products", func(c *gin.Context) {
		limit := 20
		if q := c.Query("limit"); q != "" {
			if v, err := strconv.Atoi(q); err == nil {
				limit = v
			}
		}
		res := service.TopProducts(records, limit)
		c.JSON(http.StatusOK, res)
	})
}
