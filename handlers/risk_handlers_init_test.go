package handlers

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNewAPIService(t *testing.T) {
    apiService := NewAPIService()
	assert.NotNil(t, apiService.risks)
	assert.Empty(t, apiService.risks)
	assert.Equal(t, 0 , len(apiService.risks))
}