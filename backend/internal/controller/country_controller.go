package controller

import (
	"net/http"
	"strconv"

	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/data"
	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterCountryRoutes(rg *gin.RouterGroup, records []data.Record) {

	rg.GET("/revenue-by-country", func(c *gin.Context) {
		res := service.RevenueByCountry(records)
		c.JSON(http.StatusOK, res)
	})

	rg.GET("/country-product-stats", func(c *gin.Context) {
		full := service.CountryProductTable(records)

		limit := 100
		offset := 0
		if q := c.Query("limit"); q != "" {
			if v, err := strconv.Atoi(q); err == nil {
				limit = v
			}
		}
		if q := c.Query("offset"); q != "" {
			if v, err := strconv.Atoi(q); err == nil {
				offset = v
			}
		}

		if offset < 0 {
			offset = 0
		}
		if offset > len(full) {
			offset = len(full)
		}
		end := offset + limit
		if end > len(full) {
			end = len(full)
		}
		page := full[offset:end]
		c.JSON(http.StatusOK, gin.H{
			"total":  len(full),
			"offset": offset,
			"limit":  limit,
			"data":   page,
		})
	})
}
