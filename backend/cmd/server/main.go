package main

import (
	"log"

	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/controller"
	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/data"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1) Load your CSV once at startup
	records, err := data.LoadRecords("../../data/GO_test_5m.csv")
	if err != nil {
		log.Fatalf("Failed to load CSV data: %v", err)
	}
	log.Printf("Loaded %d records\n", len(records))

	// 2) Initialize Gin
	router := gin.Default()
	router.Use(cors.Default())

	// 3) Health route (no prefix)
	controller.RegisterHealthRoutes(router.Group(""))

	// 4) Create API v1 group
	api := router.Group("/api/v1")

	// 5) Register controllers for each analytics group
	controller.RegisterCountryRoutes(api, records)
	controller.RegisterProductRoutes(api, records)
	controller.RegisterSalesRoutes(api, records)
	controller.RegisterRegionRoutes(api, records)

	// 6) Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
