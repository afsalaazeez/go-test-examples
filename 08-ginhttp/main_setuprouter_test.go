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
			name:           "Scenario 1: Different HTTP method used",
			method:         http.MethodPost,
			path:           "/ping",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Scenario 2: Invalid/nonexistent route used",
			method:         http.MethodGet,
			path:           "/invalid",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Scenario 3: Unexpected status code checked",
			method:         http.MethodGet,
			path:           "/ping",
			expectedStatus: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			router := setupRouter()

			req, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code, "Response status code should match the expected one")

			t.Log(tc.name)
		})
	}
}

