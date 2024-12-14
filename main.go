package main

import (
	"fmt"
	"log"

	"net/http"
	
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

// CreateRisk creates a new risk
func (apiSvc *APIService) CreateRisk(c *gin.Context) {
    var newRisk struct {
        State       string `json:"state" binding:"required"`
        Title       string `json:"title" binding:"required"`
        Description string `json:"description"`
    }

	// Parse JSON request body
    if err := c.ShouldBindJSON(&newRisk); err != nil { // Failed to parse input
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	// Type-cast into Risk type
    newRiskItem, err := NewRisk(newRisk.State, newRisk.Title, newRisk.Description)
    if err != nil { // Failed to type-cast
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	// Return created Risk item
    apiSvc.risks[newRiskItem.ID] = newRiskItem
    c.JSON(http.StatusCreated, newRiskItem)
}

// GetRiskByID fetches a specific risk
func (apiSvc *APIService) GetRiskByID(c *gin.Context) {
	// Parse input UUID
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil { // Invalid UUID
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
        return
    }

	// Find risk item by ID
    risk, exists := apiSvc.risks[id]
    if !exists { // Risk item not found
        c.JSON(http.StatusNotFound, gin.H{"error": "Risk not found"})
        return
    }

	// Return risk item
    c.JSON(http.StatusOK, risk)
}

// GetRisks lists all risks
func (apiSvc *APIService) GetRisks(c *gin.Context) {
	// Parse risk items from memory into slice of Risk type
    risks := make([]*Risk, 0, len(apiSvc.risks))
    for _, risk := range apiSvc.risks {
        risks = append(risks, risk)
    }

	// Return risk items
    c.JSON(http.StatusOK, risks)
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