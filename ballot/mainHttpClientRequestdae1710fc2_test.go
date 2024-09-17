package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func httpClientRequest(operation, hostAddr, command string, params io.Reader) (int, []byte, error) {

	url := "http://" + hostAddr + command
	if strings.Contains(hostAddr, "http://") {
		url = hostAddr + command
	}

	req, err := http.NewRequest(operation, url, params)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("Failed to create HTTP request." + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	defer resp.Body.Close()

	body, ioErr := ioutil.ReadAll(resp.Body)
	if hBit := resp.StatusCode / 100; hBit != 2 && hBit != 3 {
		if ioErr != nil {
			ioErr = fmt.Errorf("status code error %d", resp.StatusCode)
		}
	}
	return resp.StatusCode, body, ioErr
}

func TestHttpClientRequest(t *testing.T) {
	t.Run("SuccessRequest", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Success"))
		}))
		defer server.Close()

		statusCode, body, err := httpClientRequest("GET", server.URL, "/", nil)
		if err != nil {
			t.Error("Expected no error, got:", err)
		}

		if statusCode != http.StatusOK {
			t.Error("Expected status code 200, got:", statusCode)
		}

		if string(body) != "Success" {
			t.Error("Expected body 'Success', got:", string(body))
		}
	})

	t.Run("BadRequest", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
		}))
		defer server.Close()

		statusCode, body, err := httpClientRequest("GET", server.URL, "/", nil)
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if statusCode != http.StatusBadRequest {
			t.Error("Expected status code 400, got:", statusCode)
		}

		if string(body) != "Bad Request" {
			t.Error("Expected body 'Bad Request', got:", string(body))
		}
	})

	t.Run("InvalidURL", func(t *testing.T) {
		statusCode, body, err := httpClientRequest("GET", ":", "/", nil)
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if statusCode != http.StatusBadRequest {
			t.Error("Expected status code 400, got:", statusCode)
		}

		if body != nil {
			t.Error("Expected body to be nil, got:", string(body))
		}
	})

	t.Run("InvalidOperation", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Success"))
		}))
		defer server.Close()

		statusCode, body, err := httpClientRequest("INVALID", server.URL, "/", nil)
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if statusCode != http.StatusBadRequest {
			t.Error("Expected status code 400, got:", statusCode)
		}

		if body != nil {
			t.Error("Expected body to be nil, got:", string(body))
		}
	})

	t.Run("PostRequest", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Error("Expected method POST, got:", r.Method)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Post Success"))
		}))
		defer server.Close()

		data := bytes.NewBuffer([]byte(`{"key": "value"}`))
		statusCode, body, err := httpClientRequest("POST", server.URL, "/", data)
		if err != nil {
			t.Error("Expected no error, got:", err)
		}

		if statusCode != http.StatusOK {
			t.Error("Expected status code 200, got:", statusCode)
		}

		if string(body) != "Post Success" {
			t.Error("Expected body 'Post Success', got:", string(body))
		}
	})
}