package models

import (
	"github.com/google/uuid"
)

// NewRisk creates a new risk item
func NewRisk(state, title, description string) (*Risk, error) {
    risk := &Risk{
        ID:          uuid.New(),
        State:       state,
        Title:       title,
        Description: description,
    }

	// Validate Risk
    if err := risk.Validate(); err != nil {
        return nil, err
    }

    return risk, nil
}