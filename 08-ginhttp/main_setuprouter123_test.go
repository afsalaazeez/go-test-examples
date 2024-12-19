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
			method:         "GET",
			path:           "/ping",
			expectedStatus: 200,
			expectedBody:   "pong",
		},
		{
			name:           "Invalid HTTP Method to /ping Endpoint",
			method:         "POST",
			path:           "/ping",
			expectedStatus: 405,
		},
		{
			name:           "Request to Non-existent Endpoint",
			method:         "GET",
			path:           "/nonexistent",
			expectedStatus: 404,
		},
		{
			name:           "Empty Request Path",
			method:         "GET",
			path:           "",
			expectedStatus: 404,
		},
	}

	router := setupRouter()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code, "Response status code does not match")

			if tc.expectedBody != "" {
				assert.Equal(t, tc.expectedBody, rr.Body.String(), "Response body does not match")
			}

			t.Log("Test case passed: ", tc.name)
		})
	}
}
