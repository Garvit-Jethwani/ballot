```go
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"sync"
	"testing"
)

type Vote struct {
	VoterID     string
	CandidateID string
}

type Status struct {
	Code    int
	Message string
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement countVote() and saveVote(vote Vote) functions
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodGet:
		defer r.Body.Close()
		log.Println("result request received")

		voteData, err := countVote()
		out, err := json.Marshal(voteData)
		if err != nil {
			log.Println("error marshaling response to result request. error: ", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(out)

	case http.MethodPost:
		log.Println("vote received")
		vote := Vote{}
		status := Status{}
		defer writeVoterResponse(w, status)
		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&vote)
		if err != nil || vote.CandidateID == "" {
			log.Printf("error parsing vote data. error: %v\n", err)
			status.Code = http.StatusBadRequest
			status.Message = "Vote is not valid. Vote can not be saved"
			return
		}
		log.Printf("Voting done by voter: %s to candidate: %s\n", vote.VoterID, vote.CandidateID)
		err = saveVote(vote)
		if err != nil {
			log.Println(err)
			status.Code = http.StatusBadRequest
			status.Message = "Vote is not valid. Vote can not be saved"
			return
		}
		status.Code = http.StatusCreated
		status.Message = "Vote saved successfully"
		return

	default:
		status := Status{}
		status.Code = http.StatusMethodNotAllowed
		status.Message = "Bad Request. Vote can not be saved"
		return
	}
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status.Code)
	json.NewEncoder(w).Encode(status)
}

func TestServeRoot(t *testing.T) {
	t.Run("Test GET method", func(t *testing.T) {
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
	})

	t.Run("Test POST method with valid vote", func(t *testing.T) {
		vote := Vote{
			VoterID:     "Voter1",
			CandidateID: "Candidate1",
		}
		voteJSON, _ := json.Marshal(vote)
		req, err := http.NewRequest("POST", "/", bytes.NewBuffer(voteJSON))
		req.Header.Set("Content-Type", "application/json")
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
	})

	t.Run("Test POST method with invalid vote", func(t *testing.T) {
		vote := Vote{
			VoterID:     "",
			CandidateID: "",
		}
		voteJSON, _ := json.Marshal(vote)
		req, err := http.NewRequest("POST", "/", bytes.NewBuffer(voteJSON))
		req.Header.Set("Content-Type", "application/json")
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
	})

	t.Run("Test unsupported method", func(t *testing.T) {
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
	})
}