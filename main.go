package main

import (
	"log"

	"go-api-service/handlers"
	
	"github.com/gin-gonic/gin"
)

func main() {
	// New Gin router
	router := gin.Default()

	// New API service 
	riskAPIService := handlers.NewAPIService()

	// Routes
	v1 := router.Group("/v1/risks")
    {
        v1.POST("", riskAPIService.CreateRisk)
        v1.GET("", riskAPIService.GetRisks)
        v1.GET("/:id", riskAPIService.GetRiskByID)
    }

	// Start server
	log.Println("Starting service on port 8080")
	if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}