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
		desc           string
		method         string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			desc:           "Valid GET Request to /ping Endpoint",
			method:         "GET",
			path:           "/ping",
			expectedStatus: 200,
			expectedBody:   "pong",
		},
		{
			desc:           "Invalid HTTP Method to /ping Endpoint",
			method:         "POST",
			path:           "/ping",
			expectedStatus: 405,
			expectedBody:   "",
		},
		{
			desc:           "Request to Non-existent Endpoint",
			method:         "GET",
			path:           "/nonexistent",
			expectedStatus: 404,
			expectedBody:   "",
		},
		{
			desc:           "Empty Request Path",
			method:         "GET",
			path:           "",
			expectedStatus: 404,
			expectedBody:   "",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			router := setupRouter()

			req, err := http.NewRequest(tC.method, tC.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tC.expectedStatus, rr.Code, "unexpected HTTP status code")
			if tC.expectedBody != "" {
				assert.Equal(t, tC.expectedBody, rr.Body.String(), "unexpected response body")
			}
		})
	}
}
