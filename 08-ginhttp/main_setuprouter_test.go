package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
)

func TestsetupRouter(t *testing.T) {
	tt := []struct {
		name           string
		httpMethod     string
		route          string
		expectedStatus int
	}{
		{"Valid Request to /ping Endpoint", "GET", "/ping", http.StatusOK},
		{"Invalid Request to /ping Endpoint", "POST", "/ping", http.StatusMethodNotAllowed},
		{"Request to Non-existent Endpoint", "GET", "/nonexistent", http.StatusNotFound},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := setupRouter()
			req, err := http.NewRequest(tc.httpMethod, tc.route, nil)
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, w.Code)
			}

			if tc.route == "/ping" && tc.httpMethod == "GET" && w.Body.String() != "pong" {
				t.Errorf("Expected body 'pong', got '%s'", w.Body.String())
			}
		})
	}
}
