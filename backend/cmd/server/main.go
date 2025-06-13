package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.GET("/api/revenue-by-country", func(c *gin.Context) {
		// TODO: Replace with real logic
		demo := []map[string]interface{}{
			{"country": "US", "revenue": 12345.67},
			{"country": "IN", "revenue": 8910.11},
		}
		c.JSON(http.StatusOK, demo)
	})

	router.Run(":8080")
}
