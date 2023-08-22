package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServeRoot_da512bfd0e(t *testing.T) {
	t.Run("GET request: success", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `{"Results":[{"CandidateID":"candidate1","Votes":10},{"CandidateID":"candidate2","Votes":5}],"TotalVotes":15}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("POST request: success", func(t *testing.T) {
		vote := Vote{
			VoterID:     "voter1",
			CandidateID: "candidate1",
		}
		voteJSON, _ := json.Marshal(vote)

		req, err := http.NewRequest("POST", "/", strings.NewReader(string(voteJSON)))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}

		expected := `{"Code":201,"Message":"Vote saved sucessfully"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("POST request: invalid vote data", func(t *testing.T) {
		vote := Vote{
			VoterID:     "voter1",
			CandidateID: "",
		}
		voteJSON, _ := json.Marshal(vote)

		req, err := http.NewRequest("POST", "/", strings.NewReader(string(voteJSON)))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := `{"Code":400,"Message":"Vote is not valid. Vote can not be saved"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("Invalid method: failure", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

		expected := `{"Code":405,"Message":"Bad Request. Vote can not be saved"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
