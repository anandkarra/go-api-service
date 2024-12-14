package models

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// Validate Risk fields
func (r *Risk) Validate() error {
	// Validate ID is not empty
    if r.ID == uuid.Nil {
        r.ID = uuid.New()
    } else {
        // Check UUID format using regex
        uuidStr := r.ID.String()
        if !uuidRegex.MatchString(uuidStr) {
            return fmt.Errorf("invalid UUID format")
        }
    }

    // Trim whitespace
    r.State = strings.TrimSpace(r.State)
    r.Title = strings.TrimSpace(r.Title)
    r.Description = strings.TrimSpace(r.Description)

    // Validate state is not empty
    if r.State == "" {
        return fmt.Errorf("state is required")
    }
	// Check if state is valid
    if !ValidRiskStates[r.State] {
        return fmt.Errorf("invalid risk state: %s. state can only be open, closed, accepted or investigating", r.State)
    }

	// Validate title is not empty
    if r.Title == "" {
        return fmt.Errorf("title is required")
    }
	// Validate title length
    if len(r.Title) > 100 {
        return fmt.Errorf("title cannot be longer than 100 characters")
    }
	// Validate title characters
    if !safeTextRegex.MatchString(r.Title) {
        return fmt.Errorf("title can only contain alphanumeric characters, ., -, and _")
    }

	// Validate description is not empty
	if r.Description == "" {
        return fmt.Errorf("description is required")
    }
	// Validate description length
	if len(r.Description) > 100 {
		return fmt.Errorf("description cannot be longer than 100 characters")
	}
	// Validate description characters
	if !safeTextRegex.MatchString(r.Description) {
		return fmt.Errorf("description can only contain alphanumeric characters, ., -, and _")
	}

    return nil
}