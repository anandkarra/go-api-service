package handlers

import (
	"net/http"
	
	"go-api-service/models"

	"github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

// CreateRisk creates a new risk
func (apiSvc *APIService) CreateRisk(c *gin.Context) {
    var newRisk struct {
        State       string `json:"state" binding:"required"`
        Title       string `json:"title" binding:"required"`
        Description string `json:"description" binding:"required"`
    }

	// Parse JSON request body
    if err := c.ShouldBindJSON(&newRisk); err != nil { // Failed to parse input
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	// Type-cast into Risk type
    newRiskItem, err := models.NewRisk(newRisk.State, newRisk.Title, newRisk.Description)
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
    risks := make([]*models.Risk, 0, len(apiSvc.risks))
    for _, risk := range apiSvc.risks {
        risks = append(risks, risk)
    }

	// Return risk items
    c.JSON(http.StatusOK, risks)
}