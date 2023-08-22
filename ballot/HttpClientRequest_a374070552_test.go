package main

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestHttpClientRequest_Success(t *testing.T) {
	operation := "GET"
	hostAddr := "www.example.com"
	command := "/api/resource"
	params := strings.NewReader("")

	statusCode, body, err := httpClientRequest(operation, hostAddr, command, params)
	if err != nil {
		t.Error("Expected no error, but got:", err)
	}

	if statusCode != http.StatusOK {
		t.Error("Expected status code", http.StatusOK, "but got:", statusCode)
	}

	expectedBody := []byte("success")
	if string(body) != string(expectedBody) {
		t.Error("Expected body", string(expectedBody), "but got:", string(body))
	}
}

func TestHttpClientRequest_Failure(t *testing.T) {
	operation := "POST"
	hostAddr := "www.example.com"
	command := "/api/resource"
	params := strings.NewReader("")

	statusCode, body, err := httpClientRequest(operation, hostAddr, command, params)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	expectedStatusCode := http.StatusBadRequest
	if statusCode != expectedStatusCode {
		t.Error("Expected status code", expectedStatusCode, "but got:", statusCode)
	}

	expectedErrMsg := "Failed to create HTTP request."
	if !strings.Contains(err.Error(), expectedErrMsg) {
		t.Error("Expected error message to contain", expectedErrMsg, "but got:", err.Error())
	}
}
