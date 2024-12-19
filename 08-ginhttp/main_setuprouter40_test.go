package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestsetupRouter(t *testing.T) {

	tests := []struct {
		name     string
		method   string
		endpoint string
		status   int
		body     string
	}{
		{
			name:     "Valid GET Request to /ping Endpoint",
			method:   "GET",
			endpoint: "/ping",
			status:   http.StatusOK,
			body:     "pong",
		},
		{
			name:     "Invalid Method Request to /ping Endpoint",
			method:   "POST",
			endpoint: "/ping",
			status:   http.StatusMethodNotAllowed,
		},
		{
			name:     "Request to Non-existent Endpoint",
			method:   "GET",
			endpoint: "/nonexistent",
			status:   http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			router := setupRouter()

			req, _ := http.NewRequest(tt.method, tt.endpoint, nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.status, resp.Code)

			if resp.Code == http.StatusOK {
				assert.Equal(t, tt.body, resp.Body.String())
			}
		})
	}
}
