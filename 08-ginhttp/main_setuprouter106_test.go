package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
)

func TestsetupRouter(t *testing.T) {

	r := setupRouter()

	tests := []struct {
		method     string
		url        string
		statusCode int
		body       string
	}{

		{http.MethodGet, "/ping", http.StatusOK, "pong"},

		{http.MethodPost, "/ping", http.StatusMethodNotAllowed, ""},

		{http.MethodGet, "/nonexistent", http.StatusNotFound, ""},
	}

	for _, test := range tests {
		t.Logf("Running scenario where method is %s and url is %s", test.method, test.url)

		req, err := http.NewRequest(test.method, test.url, nil)
		if err != nil {
			t.Fatalf("Failed to make a %s request %v", test.method, err)
		}

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != test.statusCode {
			t.Fatalf("Expected to get status %d but instead got %d", test.statusCode, rec.Code)
		}

		if rec.Body.String() != test.body {
			t.Fatalf("Expected to get body %q but instead got %q", test.body, rec.Body.String())
		}
	}
}
