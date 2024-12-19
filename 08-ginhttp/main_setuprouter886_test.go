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
		name           string
		method         string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid GET request to /ping endpoint",
			method:         "GET",
			path:           "/ping",
			expectedStatus: http.StatusOK,
			expectedBody:   "pong",
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
			path:           "/nonexistent",
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

			r := setupRouter()
			req, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			assert.Equal(t, tc.expectedStatus, resp.Code)
			if resp.Code == http.StatusOK {
				assert.Equal(t, tc.expectedBody, resp.Body.String())
			}
		})
	}
}
