package main

import (
	"fmt"
	"log"
	
	"github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

// ValidRiskStates defines allowed values for a risk state
var ValidRiskStates = map[string]bool{
    "open":          true,
    "closed":        true,
    "accepted":      true,
    "investigating": true,
}

// Risk defines a risk item
type Risk struct {
    ID          uuid.UUID `json:"id"`
    State       string    `json:"state"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
}

// NewRisk creates a new risk item
func NewRisk(state, title, description string) (*Risk, error) {
    // Validate state
    if !ValidRiskStates[state] {
        return nil, fmt.Errorf("invalid risk state: %s", state)
    }

    return &Risk{
        ID:          uuid.New(),
        State:       state,
        Title:       title,
        Description: description,
    }, nil
}

// APIService stores risks in-memory
type APIService struct {
    risks map[uuid.UUID]*Risk
}

// NewAPIService creates a new APIService
func NewAPIService() *APIService {
	return &APIService{
		risks: make(map[uuid.UUID]*Risk),
	}
}

func main() {
	// New Gin router
	router := gin.Default()

	// New API service 
	riskAPIService := NewAPIService()

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