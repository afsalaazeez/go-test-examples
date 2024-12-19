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
		name       string
		method     string
		path       string
		statusCode int
		body       string
	}{
		{
			name:       "Valid Request to /ping Endpoint",
			method:     "GET",
			path:       "/ping",
			statusCode: 200,
			body:       "pong",
		},
		{
			name:       "Invalid HTTP Method to /ping Endpoint",
			method:     "POST",
			path:       "/ping",
			statusCode: 405,
			body:       "",
		},
		{
			name:       "Request to Non-existent Endpoint",
			method:     "GET",
			path:       "/nonexistent",
			statusCode: 404,
			body:       "",
		},
		{
			name:       "Empty Request Path",
			method:     "GET",
			path:       "",
			statusCode: 404,
			body:       "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, err := http.NewRequest(tt.method, tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			r := setupRouter()
			r.ServeHTTP(rr, req)

			assert.Equal(t, tt.statusCode, rr.Code)

			if tt.statusCode == 200 {
				assert.Equal(t, tt.body, rr.Body.String())
			}
		})
	}
}
