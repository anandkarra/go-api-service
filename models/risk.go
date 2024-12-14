package models

import (
	"regexp"

	"github.com/google/uuid"
)

var (
	// ValidRiskStates defines allowed values for a risk state
	ValidRiskStates = map[string]bool{
		"open":          true,
		"closed":        true,
		"accepted":      true,
		"investigating": true,
	}

	// Regex for validating title and description
	safeTextRegex = regexp.MustCompile(`^[a-zA-Z0-9\.\-_ ]+$`)
    uuidRegex     = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
)

// Risk defines a risk item
type Risk struct {
    ID          uuid.UUID `json:"id"`
    State       string    `json:"state"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
}