package handlers

import (
	"go-api-service/models"

	"github.com/google/uuid"
)
// APIService stores risks in-memory
type APIService struct {
    risks map[uuid.UUID]*models.Risk
}

// NewAPIService creates a new APIService
func NewAPIService() *APIService {
	return &APIService{
		risks: make(map[uuid.UUID]*models.Risk),
	}
}