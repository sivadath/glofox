package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/sivadath/glofox/internal/errors"
	"github.com/sivadath/glofox/mocks"
	"github.com/sivadath/glofox/models"
	"github.com/sivadath/glofox/routes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setUpRouter initializes and returns a new Gin router instance.
func setUpRouter() *gin.Engine {
	return gin.Default()
}

func TestCreateClass(t *testing.T) {
	tests := []struct {
		name             string
		storage          func(reporter gomock.TestReporter) *mocks.MockStorage
		requestBody      map[string]interface{}
		expectedResponse map[string]interface{}
		expectedStatus   int
	}{
		{
			name: "Successful Class Creation",
			storage: func(reporter gomock.TestReporter) *mocks.MockStorage {
				db := mocks.NewMockStorage(gomock.NewController(reporter))
				db.EXPECT().
					AddClass(gomock.Any(), gomock.Any()).
					Return(models.Class{
						Name:      "Yoga",
						StartDate: models.Date(time.Date(2025, 2, 12, 0, 0, 0, 0, time.UTC)),
						EndDate:   models.Date(time.Date(2025, 2, 25, 0, 0, 0, 0, time.UTC)),
						Capacity:  10},
						nil).
					Times(1)
				return db
			},
			requestBody: map[string]interface{}{
				"name":       "Yoga",
				"start_date": "2025-02-12",
				"end_date":   "2025-02-25",
				"capacity":   10,
			},
			expectedResponse: map[string]interface{}{
				"id":         "",
				"name":       "Yoga",
				"start_date": "2025-02-12",
				"end_date":   "2025-02-25",
				"capacity":   float64(10),
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Invalid request - Date Mismatch",
			storage: func(reporter gomock.TestReporter) *mocks.MockStorage {
				return mocks.NewMockStorage(gomock.NewController(reporter))
			},
			requestBody: map[string]interface{}{
				"name":       "Yoga",
				"start_date": "2025-02-28",
				"end_date":   "2025-02-25",
				"capacity":   10,
			},
			expectedResponse: map[string]interface{}{
				"error": errors.ErrDateMismatch.Message,
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Database Failure",
			storage: func(reporter gomock.TestReporter) *mocks.MockStorage {
				db := mocks.NewMockStorage(gomock.NewController(reporter))
				db.EXPECT().
					AddClass(gomock.Any(), gomock.Any()).
					Return(models.Class{}, fmt.Errorf("database error")).
					Times(1)
				return db
			},
			requestBody: map[string]interface{}{
				"name":       "Yoga",
				"start_date": "2025-02-12",
				"end_date":   "2025-02-25",
				"capacity":   10,
			},
			expectedResponse: map[string]interface{}{
				"error": "database error",
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setUpRouter()
			routes.RegisterClassRoutes(r, tt.storage(t))

			reqBody, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, path.Join(routes.Version, routes.EndPointClasses), bytes.NewBuffer(reqBody))
			assert.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			respBody, err := io.ReadAll(w.Body)
			log.Print("bosy", string(respBody))
			assert.NoError(t, err)

			var respData map[string]interface{}
			assert.NoError(t, json.Unmarshal(respBody, &respData))

			assert.Equal(t, tt.expectedResponse, respData)
		})
	}
}
