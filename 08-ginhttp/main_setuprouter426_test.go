package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestsetupRouter(t *testing.T) {

	testCases := []struct {
		name       string
		method     string
		path       string
		statusCode int
		body       string
	}{
		{
			name:       "Scenario 1: Test successful GET request to /ping endpoint",
			method:     "GET",
			path:       "/ping",
			statusCode: http.StatusOK,
			body:       "pong",
		},
		{
			name:       "Scenario 2: Test unsuccessful POST request to /ping endpoint",
			method:     "POST",
			path:       "/ping",
			statusCode: http.StatusMethodNotAllowed,
		},
		{
			name:       "Scenario 3: Test request to nonexistent endpoint",
			method:     "GET",
			path:       "/nonexistent",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "Scenario 4: Test request with invalid HTTP method",
			method:     "INVALID",
			path:       "/ping",
			statusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			router := setupRouter()

			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.statusCode, rr.Code, "Unexpected status code")

			if tc.body != "" {
				assert.Equal(t, tc.body, rr.Body.String(), "Unexpected body")
			}
		})
	}
}
