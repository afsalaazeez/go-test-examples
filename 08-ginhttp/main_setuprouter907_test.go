package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {

	testCases := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid GET Request to /ping Endpoint",
			method:         http.MethodGet,
			path:           "/ping",
			expectedStatus: http.StatusOK,
			expectedBody:   "pong",
		},
		{
			name:           "Invalid Request to /ping Endpoint",
			method:         http.MethodPost,
			path:           "/ping",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Request to Non-existent Endpoint",
			method:         http.MethodGet,
			path:           "/nonexistent",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Empty Request Path",
			method:         http.MethodGet,
			path:           "",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			router := setupRouter()

			req, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, tc.expectedStatus, resp.Code, "Expected status code does not match actual status code")

			if tc.expectedBody != "" {
				assert.Equal(t, tc.expectedBody, resp.Body.String(), "Expected response body does not match actual response body")
			}
		})
	}
}
