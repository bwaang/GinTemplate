// handlers.user_test.go

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
  "strings"
)

func TestHealthCheck(t *testing.T) {
  r := getRouter()

  // Define the route similar to its definition in the routes file
  r.GET("/file/v1/health", healthCheck)

  // Create a request to send to the above route
  req, _ := http.NewRequest("GET", "/file/v1/health", nil)

  testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Contains(string(p), `{"status":"Ok"}`)

    return statusOK && pageOK
  })
}

func TestVersionCheck(t *testing.T) {
  r := getRouter()

  // Define the route similar to its definition in the routes file
  r.GET("/file/v1/version", versionCheck)

  // Create a request to send to the above route
  req, _ := http.NewRequest("GET", "/file/v1/version", nil)

  testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Contains(string(p), `{"version":"v0.0.0"}`)

    return statusOK && pageOK
  })
}
