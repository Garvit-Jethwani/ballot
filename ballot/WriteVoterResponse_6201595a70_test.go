package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteVoterResponse_Success(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	status := Status{Code: 200, Message: "Success"}
	writeVoterResponse(w, status)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	expectedResp := `{"code": 200, "message": "Success"}`
	if w.Body.String() != expectedResp {
		t.Errorf("Expected response body %q, but got %q", expectedResp, w.Body.String())
	}
}

func TestWriteVoterResponse_Error(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	status := Status{Code: 500, Message: "Internal Server Error"}
	writeVoterResponse(w, status)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	expectedResp := `{"code": 500, "message": "Internal Server Error"}`
	if w.Body.String() != expectedResp {
		t.Errorf("Expected response body %q, but got %q", expectedResp, w.Body.String())
	}
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(status)
	if err != nil {
		log.Println("error marshaling response to vote request. error: ", err)
	}
	w.Write(resp)
}
