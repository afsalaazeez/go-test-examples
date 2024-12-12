package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

// TestSetupRouter is a unit test for the function setupRouter
func TestSetupRouter(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	router := setupRouter()

	// This is how you record HTTP responses
	recorder := httptest.NewRecorder()

	// Create a request to pass to our router
	req, _ := http.NewRequest("GET", "/ping", nil)

	// Create a test cases
	testCases := []struct {
		name       string
		statusCode int
	}{
		{"Scenario 1: Valid GET /ping", http.StatusOK},
	}

	// We don't have any form data to send, so we'll pass 'nil' as the third parameter
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			router.ServeHTTP(recorder, req)
			assert.Equal(t, tt.statusCode, recorder.Code)
			// Check the response body is what we expect.
			expected := `pong`
			actual := recorder.Body.String()
			assert.Equal(t, expected, actual)
		})
	}
}
