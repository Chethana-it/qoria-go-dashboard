package main

import (
	"log"

	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/controller"
	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/data"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	records, err := data.LoadRecords("../../data/GO_test_5m.csv")
	if err != nil {
		log.Fatalf("Failed to load CSV data: %v", err)
	}
	log.Printf("Loaded %d records\n", len(records))

	router := gin.Default()
	router.Use(cors.Default())

	controller.RegisterHealthRoutes(router.Group(""))

	api := router.Group("/api/v1")

	controller.RegisterCountryRoutes(api, records)
	controller.RegisterProductRoutes(api, records)
	controller.RegisterSalesRoutes(api, records)
	controller.RegisterRegionRoutes(api, records)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
