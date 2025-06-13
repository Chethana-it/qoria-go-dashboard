package controller

import (
	"net/http"
	"strconv"

	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/data"
	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterRegionRoutes(rg *gin.RouterGroup, records []data.Record) {
	rg.GET("/revenue-by-region", func(c *gin.Context) {
		limit := 30
		if q := c.Query("limit"); q != "" {
			if v, err := strconv.Atoi(q); err == nil {
				limit = v
			}
		}
		res := service.RevenueByRegion(records, limit)
		c.JSON(http.StatusOK, res)
	})
}
