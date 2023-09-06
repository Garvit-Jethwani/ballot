package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func httpClientRequest(method, url, path string, body io.ReadCloser) (int, string, error) {
    client := &http.Client{}
    req, err := http.NewRequest(method, url+path, body)
    if err != nil {
        return 0, "", err
    }

    resp, err := client.Do(req)
    if err != nil {
        return 0, "", err
    }
    defer resp.Body.Close()

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return 0, "", err
    }

    return resp.StatusCode, string(respBody), nil
}

func TestHttpClientRequest_a374070552(t *testing.T) {

	// Test case 1: Successful request
	t.Run("Successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		statusCode, _, err := httpClientRequest("GET", server.URL, "/test", nil)
		if err != nil {
			t.Error("Expected no error, got ", err)
		}

		if statusCode != http.StatusOK {
			t.Error("Expected status code 200, got ", statusCode)
		}
	})

	// Test case 2: Server returns error
	t.Run("Server returns error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		statusCode, _, err := httpClientRequest("GET", server.URL, "/test", nil)
		if err == nil {
			t.Error("Expected error, got none")
		}

		if statusCode != http.StatusInternalServerError {
			t.Error("Expected status code 500, got ", statusCode)
		}
	})

	// Test case 3: Invalid operation
	t.Run("Invalid operation", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		_, _, err := httpClientRequest("INVALID", server.URL, "/test", nil)
		if err == nil {
			t.Error("Expected error, got none")
		}
	})

	// Test case 4: Invalid host address
	t.Run("Invalid host address", func(t *testing.T) {
		_, _, err := httpClientRequest("GET", "invalid_host_address", "/test", nil)
		if err == nil {
			t.Error("Expected error, got none")
		}
	})

	// Test case 5: Invalid command
	t.Run("Invalid command", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		_, _, err := httpClientRequest("GET", server.URL, "invalid_command", nil)
		if err == nil {
			t.Error("Expected error, got none")
		}
	})

	// Test case 6: Invalid parameters
	t.Run("Invalid parameters", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		_, _, err := httpClientRequest("GET", server.URL, "/test", ioutil.NopCloser(bytes.NewReader([]byte("invalid_parameters"))))
		if err == nil {
			t.Error("Expected error, got none")
		}
	})
}
