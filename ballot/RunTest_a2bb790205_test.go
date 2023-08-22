package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func runTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	log.Println("ballot endpoint tests running")
	status := Status{}
	err := TestBallot()
	if err != nil {
		status.Message = fmt.Sprintf("Test Cases Failed with error : %v", err)
		status.Code = http.StatusBadRequest
	}
	status.Message = "Test Cases passed"
	status.Code = http.StatusOK
	writeVoterResponse(w, status)
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	// TODO: Implement writeVoterResponse function
}

func TestBallot(t *testing.T) error {
	// TODO: Implement TestBallot function
	return nil
}

// TODO: Implement TestMethodName test cases
func TestRunTest_a2bb790205(t *testing.T) {
	// TODO: Implement test cases for runTest function
}
