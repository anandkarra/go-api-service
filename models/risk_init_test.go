package models

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNewRisk(t *testing.T) {
    testCases := []struct {
		name string
        state        string
		title string
		description string
        risk        *Risk
        expectedErr bool
        errMessage  string
    }{
		{
            name: "Valid Risk",
            state: "open",
			title: "Valid Risk Title",
			description: "Valid Description",
            expectedErr: false,
        },
		{
            name: "Invalid Risk",
            state: "invalid-state",
			title: "Valid Risk Title",
			description: "Valid Description",
            expectedErr: true,
			errMessage:  "invalid risk state",
        },
	}
	for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            risk, err := NewRisk(tc.state, tc.title, tc.description)
            
            if tc.expectedErr {
                assert.Error(t, err)
                if tc.errMessage != "" {
                    assert.Contains(t, err.Error(), tc.errMessage)
                }
            } else {
                assert.NoError(t, err)
				assert.Contains(t, risk.State, tc.state)
				assert.Contains(t, risk.Title, tc.title)
				assert.Contains(t, risk.Description, tc.description)
            }
        })
    }
}