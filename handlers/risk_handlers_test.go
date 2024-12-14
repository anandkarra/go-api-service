package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-api-service/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateRisk(t *testing.T) {
	apiSvc := NewAPIService()
	router := gin.Default()
	router.POST("/risks", apiSvc.CreateRisk)

	testCases := []struct {
        name        string
        risk        *models.Risk
        expectedErr bool
        expectedRespCode int
    }{
        {
            name: "Success",
            risk: &models.Risk{
                ID:          uuid.New(),
                State:       "open",
                Title:       "Valid Risk Title",
                Description: "Valid Description",
            },
            expectedErr: false,
			expectedRespCode: http.StatusCreated,
        },
		{
            name: "Invalid Input Format",
            risk: &models.Risk{},
            expectedErr: true,
			expectedRespCode: http.StatusInternalServerError,
        },
		{
            name: "Invalid Input Values",
			risk: &models.Risk{
                ID:          uuid.New(),
                State:       "invalid-state",
                Title:       "Valid Risk Title",
                Description: "Valid Description",
            },
            expectedErr: true,
			expectedRespCode: http.StatusBadRequest,
        },
	}

	for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tc.risk)
			req := httptest.NewRequest(http.MethodPost, "/risks", bytes.NewReader(requestBody))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, tc.expectedRespCode, resp.Code)

			if !tc.expectedErr {
				var risk models.Risk
				assert.NoError(t, json.Unmarshal(resp.Body.Bytes(), &risk))
				assert.Equal(t, tc.risk.State, risk.State)
				assert.Equal(t, tc.risk.Title, risk.Title)
				assert.Equal(t, tc.risk.Description, risk.Description)
			}			
		})
	}
}

func TestGetRiskByID(t *testing.T) {
	apiSvc := NewAPIService()
	router := gin.Default()
	router.GET("/risks/:id", apiSvc.GetRiskByID)
	
	// Create a valid risk item
	id := uuid.New()
	apiSvc.risks[id] = &models.Risk{
		ID:          id,
		State:       "open",
		Title:       "Valid Risk Title",
		Description: "Valid Description",
	}

	testCases := []struct {
        name             string
        id               string
        expectedErr      bool
        expectedRespCode int
        errMessage       string
    }{
        {
            name: "Success",
            id: id.String(),
            expectedErr: false,
			expectedRespCode: http.StatusOK,
        },
		{
            name: "Invalid UUID",
			id:          "invalid-uuid",
			expectedErr: true,
			expectedRespCode: http.StatusBadRequest,
		},
		{
            name: "Risk Not Found",
			id:          uuid.New().String(),
			expectedErr: true,
			expectedRespCode: http.StatusNotFound,
		},
	}
	for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/risks/"+tc.id, nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, tc.expectedRespCode, resp.Code)

			if !tc.expectedErr {
				var risk models.Risk
				assert.NoError(t, json.Unmarshal(resp.Body.Bytes(), &risk))
				assert.Equal(t, id, risk.ID)
			}
		})
	}
}

func TestGetRisks(t *testing.T) {
	apiSvc := NewAPIService()
	router := gin.Default()
	router.GET("/risks", apiSvc.GetRisks)

	// Create risk items
	id1 := uuid.New()
	id2 := uuid.New()
	apiSvc.risks[id1] = &models.Risk{
		ID:          id1,
		State:       "open",
		Title:       "Valid Risk Title",
		Description: "Valid Risk Description",
	}
	apiSvc.risks[id2] = &models.Risk{
		ID:          id2,
		State:       "closed",
		Title:       "Another Valid Risk Title",
		Description: "Another Valid Risk Description",
	}

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/risks", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var risks []models.Risk
		assert.NoError(t, json.Unmarshal(resp.Body.Bytes(), &risks))
		assert.Len(t, risks, 2)
	})
}
