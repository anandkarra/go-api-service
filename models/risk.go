package models

import (
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