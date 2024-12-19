package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {

	tests := []struct {
		name           string
		method         string
		endpoint       string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Test for successful GET request to /ping endpoint",
			method:         "GET",
			endpoint:       "/ping",
			expectedStatus: 200,
			expectedBody:   "pong",
		},
		{
			name:           "Test for unsuccessful request to /ping endpoint with a method other than GET",
			method:         "POST",
			endpoint:       "/ping",
			expectedStatus: 405,
		},
		{
			name:           "Test for request to non-existent endpoint",
			method:         "GET",
			endpoint:       "/fake",
			expectedStatus: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, err := http.NewRequest(tt.method, tt.endpoint, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			router := setupRouter()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code, "status code should match")

			if tt.expectedBody != "" {

				assert.Equal(t, tt.expectedBody, rr.Body.String(), "response body should match")
			}
		})
	}
}
