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
	}{
		{
			name:           "Valid GET request to /ping endpoint",
			method:         "GET",
			path:           "/ping",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid HTTP method to /ping endpoint",
			method:         "POST",
			path:           "/ping",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Request to non-existent endpoint",
			method:         "GET",
			path:           "/notExist",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Empty request path",
			method:         "GET",
			path:           "",
			expectedStatus: http.StatusNotFound,
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

			assert.Equal(t, tc.expectedStatus, rr.Code, "Unexpected status code")

			if tc.name == "Valid GET request to /ping endpoint" {
				assert.Equal(t, "pong", rr.Body.String(), "Unexpected response body")
			}
		})
	}
}
