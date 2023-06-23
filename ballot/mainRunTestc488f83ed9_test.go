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

type Status struct {
	Message string
	Code    int
}

func TestRunTest(t *testing.T) {
	t.Run("TestCasesPassed", func(t *testing.T) {
		// Mock TestBallot to return nil error
		originalTestBallot := TestBallot
		defer func() { TestBallot = originalTestBallot }()
		TestBallot = func() error {
			return nil
		}

		req, err := http.NewRequest("GET", "/runTest", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(runTest)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `{"Message":"Test Cases passed","Code":200}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("TestCasesFailed", func(t *testing.T) {
		// Mock TestBallot to return an error
		originalTestBallot := TestBallot
		defer func() { TestBallot = originalTestBallot }()
		TestBallot = func() error {
			return errors.New("test error")
		}

		req, err := http.NewRequest("GET", "/runTest", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(runTest)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := `{"Message":"Test Cases Failed with error : test error","Code":400}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}