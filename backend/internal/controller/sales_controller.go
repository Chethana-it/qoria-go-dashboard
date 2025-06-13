package controller

import (
	"net/http"

	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/data"
	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/service"
	"github.com/gin-gonic/gin"
)

// RegisterSalesRoutes registers sales analytics endpoints
func RegisterSalesRoutes(rg *gin.RouterGroup, records []data.Record) {
	rg.GET("/sales-by-month", func(c *gin.Context) {
		res := service.SalesByMonth(records)
		c.JSON(http.StatusOK, res)
	})
}
