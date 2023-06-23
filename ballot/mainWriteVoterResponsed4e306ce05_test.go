package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Status struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(status)
	if err != nil {
		log.Println("error marshaling response to vote request. error: ", err)
	}
	w.Write(resp)
}

func TestWriteVoterResponse(t *testing.T) {
	t.Run("ValidStatus", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		status := Status{
			Message: "Success",
			Code:    200,
		}

		writeVoterResponse(recorder, status)

		result := recorder.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, result.StatusCode)
		}

		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			t.Fatal("Error reading response body:", err)
		}

		var response Status
		err = json.Unmarshal(body, &response)
		if err != nil {
			t.Fatal("Error unmarshaling response body:", err)
		}

		if response.Message != status.Message || response.Code != status.Code {
			t.Errorf("Expected response %+v, got %+v", status, response)
		}
	})

	t.Run("EmptyStatus", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		status := Status{}

		writeVoterResponse(recorder, status)

		result := recorder.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, result.StatusCode)
		}

		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			t.Fatal("Error reading response body:", err)
		}

		var response Status
		err = json.Unmarshal(body, &response)
		if err != nil {
			t.Fatal("Error unmarshaling response body:", err)
		}

		if response.Message != status.Message || response.Code != status.Code {
			t.Errorf("Expected response %+v, got %+v", status, response)
		}
	})
}