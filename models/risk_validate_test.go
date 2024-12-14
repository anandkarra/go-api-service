package models

import (
    "testing"

    "github.com/google/uuid"
    "github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
    testCases := []struct {
        name        string
        risk        *Risk
        expectedErr bool
        errMessage  string
    }{
        {
            name: "Valid Risk",
            risk: &Risk{
                ID:          uuid.New(),
                State:       "open",
                Title:       "Valid Risk Title",
                Description: "Valid Description",
            },
            expectedErr: false,
        },
		{
            name: "Invalid UUID",
            risk: &Risk{
                ID:          uuid.UUID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
                State:       "open",
                Title:       "Valid Risk Title",
                Description: "Valid Description",
            },
            expectedErr: false,
			errMessage: "invalid UUID format",
        },
		{
            name: "Empty State",
            risk: &Risk{
                ID:          uuid.New(),
                State:       "",
                Title:       "Valid Risk Title",
                Description: "Valid Description",
            },
            expectedErr: true,
            errMessage:  "state is required",
        },
        {
            name: "Invalid State",
            risk: &Risk{
                ID:          uuid.New(),
                State:       "invalid-state",
                Title:       "Valid Risk Title",
                Description: "Valid Description",
            },
            expectedErr: true,
            errMessage:  "invalid risk state",
        },
        {
            name: "Empty Title",
            risk: &Risk{
                ID:          uuid.New(),
                State:       "open",
                Title:       "",
                Description: "Valid Description",
            },
            expectedErr: true,
            errMessage:  "title is required",
        },
        {
            name: "Title Too Long",
            risk: &Risk{
                ID:          uuid.New(),
                State:       "open",
                Title:       "This is a very long title that exceeds the maximum allowed length of one hundred characters and should trigger a validation error",
                Description: "Valid Description",
            },
            expectedErr: true,
            errMessage:  "title cannot be longer than 100 characters",
        },
        {
            name: "Invalid Title Characters",
            risk: &Risk{
                ID:          uuid.New(),
                State:       "open",
                Title:       "Invalid Title with @special characters!",
                Description: "Valid Description",
            },
            expectedErr: true,
            errMessage:  "title can only contain alphanumeric characters, ., -, and _",
        },
		{
            name: "Empty Description",
            risk: &Risk{
                ID:          uuid.New(),
                State:       "open",
                Title:       "Valid Risk Title",
                Description: "",
            },
            expectedErr: true,
            errMessage:  "description is required",
        },
        {
            name: "Description Too Long",
            risk: &Risk{
                ID:          uuid.New(),
                State:       "open",
                Title:       "Valid Title",
                Description: "This is a very long description that exceeds the maximum allowed length of one hundred characters and should trigger a validation error",
            },
            expectedErr: true,
            errMessage:  "description cannot be longer than 100 characters",
        },
        {
            name: "Invalid Description Characters",
            risk: &Risk{
                ID:          uuid.New(),
                State:       "open",
                Title:       "Valid Title",
                Description: "Invalid Description with @special characters!",
            },
            expectedErr: true,
            errMessage:  "description can only contain alphanumeric characters, ., -, and _",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            err := tc.risk.Validate()
            
            if tc.expectedErr {
                assert.Error(t, err)
                if tc.errMessage != "" {
                    assert.Contains(t, err.Error(), tc.errMessage)
                }
            } else {
                assert.NoError(t, err)
            }
        })
    }

    // Additional UUID validation test
    t.Run("Invalid UUID", func(t *testing.T) {
        invalidRisk := &Risk{
            ID:          uuid.UUID{}, // Zero UUID
            State:       "open",
            Title:       "Valid Title",
            Description: "Valid Description",
        }
        
        err := invalidRisk.Validate()
        assert.NoError(t, err, "Zero UUID should be replaced with a new UUID")
        assert.NotEqual(t, uuid.UUID{}, invalidRisk.ID, "UUID should be auto-generated")
    })
}