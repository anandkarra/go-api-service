package models

import (
	"fmt"
	"github.com/google/uuid"
)

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