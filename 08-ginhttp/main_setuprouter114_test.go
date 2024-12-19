package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestsetupRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	scenarios := []struct {
		desc     string
		method   string
		path     string
		status   int
		response string
	}{
		{
			"Valid GET Request to /ping Endpoint",
			"GET",
			"/ping",
			200,
			"pong",
		},
		{
			"Invalid HTTP Method to /ping Endpoint",
			"POST",
			"/ping",
			405,
			"Method Not Allowed",
		},
		{
			"Request to Non-existent Endpoint",
			"GET",
			"/nonexistent",
			404,
			"404 page not found",
		},
		{
			"Empty Request Path",
			"GET",
			"",
			404,
			"404 page not found",
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			t.Log("Scenario:", s.desc)

			req, err := http.NewRequest(s.method, s.path, nil)
			if err != nil {
				t.Fatal("Failed to create request:", err)
			}

			rec := httptest.NewRecorder()
			router := setupRouter()
			router.ServeHTTP(rec, req)

			assert.Equal(t, s.status, rec.Code, "Expected status code does not match actual status code")
			assert.Equal(t, s.response, rec.Body.String(), "Expected response does not match actual response")
		})
	}
}
