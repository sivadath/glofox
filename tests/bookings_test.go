package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sivadath/glofox/internal/errors"
	"github.com/sivadath/glofox/mocks"
	"github.com/sivadath/glofox/models"
	"github.com/sivadath/glofox/routes"
)

func TestCreateBooking(t *testing.T) {
	tests := []struct {
		name             string
		setupMock        func(db *mocks.MockStorage)
		requestBody      map[string]interface{}
		expectedResponse map[string]interface{}
		expectedStatus   int
	}{
		{
			name: "Successful Booking",
			setupMock: func(db *mocks.MockStorage) {
				db.EXPECT().
					GetClassesByDate(gomock.Any(), time.Date(2025, 2, 22, 0, 0, 0, 0, time.UTC)).
					Return([]models.Class{{ID: "1234"}}, nil)
				db.EXPECT().
					AddBooking(gomock.Any(), gomock.Any()).
					Return(models.Booking{Name: "Harry Potter", ID: "2345", Date: "2025-02-22", ClassID: "1234"}, nil)
			},
			requestBody: map[string]interface{}{
				"date": "2025-02-22",
				"name": "Harry Potter",
			},
			expectedResponse: map[string]interface{}{
				"date":     "2025-02-22",
				"name":     "Harry Potter",
				"id":       "2345",
				"class_id": "1234",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Classes schema issue",
			setupMock: func(db *mocks.MockStorage) {
				db.EXPECT().
					GetClassesByDate(gomock.Any(), time.Date(2025, 2, 22, 0, 0, 0, 0, time.UTC)).
					Return(nil, fmt.Errorf("connection failed"))
			},
			requestBody: map[string]interface{}{
				"date": "2025-02-22",
				"name": "Harry Potter",
			},
			expectedResponse: map[string]interface{}{
				"error": "connection failed",
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "Bookings schema issue",
			setupMock: func(db *mocks.MockStorage) {
				db.EXPECT().
					GetClassesByDate(gomock.Any(), time.Date(2025, 2, 22, 0, 0, 0, 0, time.UTC)).
					Return([]models.Class{{ID: "1234"}}, nil)
				db.EXPECT().
					AddBooking(gomock.Any(), gomock.Any()).
					Return(models.Booking{}, fmt.Errorf("connection failed"))
			},
			requestBody: map[string]interface{}{
				"date": "2025-02-22",
				"name": "Harry Potter",
			},
			expectedResponse: map[string]interface{}{
				"error": "connection failed",
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "No classes available",
			setupMock: func(db *mocks.MockStorage) {
				db.EXPECT().
					GetClassesByDate(gomock.Any(), time.Date(2025, 2, 22, 0, 0, 0, 0, time.UTC)).
					Return(nil, nil)
			},
			requestBody: map[string]interface{}{
				"date": "2025-02-22",
				"name": "Harry Potter",
			},
			expectedResponse: map[string]interface{}{
				"error": errors.ErrNoClassesFound.Message,
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockDB := mocks.NewMockStorage(ctrl)
			tt.setupMock(mockDB)

			// Setting up router
			r := setUpRouter()
			routes.RegisterBookingRoutes(r, mockDB)

			// Sending request and getting response
			resp, body := sendRequest(t, r, tt.requestBody)

			// Validating responses
			assert.Equal(t, tt.expectedStatus, resp.Code)

			var respData map[string]interface{}
			assert.NoError(t, json.Unmarshal(body, &respData))
			assert.JSONEq(t, toJSON(tt.expectedResponse), toJSON(respData))
		})
	}
}

// Helper to send HTTP request and return response
func sendRequest(t *testing.T, r http.Handler, requestBody map[string]interface{}) (*httptest.ResponseRecorder, []byte) {
	reqBody, err := json.Marshal(requestBody)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, routes.Version+routes.EndPointBookings, bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	respBody, err := io.ReadAll(w.Body)
	assert.NoError(t, err)

	return w, respBody
}

// Helper to convert struct to JSON string for comparison
func toJSON(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}
