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
		description string
		method      string
		path        string
		statusCode  int
		body        string
	}{
		{
			description: "Valid Request to /ping Endpoint",
			method:      "GET",
			path:        "/ping",
			statusCode:  http.StatusOK,
			body:        "pong",
		},
		{
			description: "Invalid HTTP Method to /ping Endpoint",
			method:      "POST",
			path:        "/ping",
			statusCode:  http.StatusMethodNotAllowed,
		},
		{
			description: "Request to Non-existent Endpoint",
			method:      "GET",
			path:        "/nonexistent",
			statusCode:  http.StatusNotFound,
		},
		{
			description: "Empty Request Path",
			method:      "GET",
			path:        "",
			statusCode:  http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			req, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			router := setupRouter()

			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.statusCode, rr.Code, "Status code should match expected value")

			if tc.body != "" {
				assert.Equal(t, tc.body, rr.Body.String(), "Response body should match expected value")
			}
		})
	}
}
