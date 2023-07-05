package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"testing"
)

type Response struct {
	TotalVotes int
}

type Vote struct {
	CandidateID string
	VoterID     string
}

type Status struct {
	Code int
}

var httpClientRequest = func(method, url, path string, body io.Reader) (int, []byte, error) {
	// TODO: Implement the actual http client request here.
	return 0, nil, nil
}

var port = "8080" // TODO: Change this to the actual port number.

func TestBallot(t *testing.T) {
	_, result, err := httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		t.Fatalf("Failed to get ballot count resp:%s error:%+v", string(result), err)
	}
	log.Println("get ballot resp:", string(result))
	var initalRespData Response
	if err = json.Unmarshal(result, &initalRespData); err != nil {
		t.Fatalf("Failed to unmarshal get ballot response. %+v", err)
	}

	var ballotvotereq Vote
	ballotvotereq.CandidateID = fmt.Sprint(rand.Intn(10))
	ballotvotereq.VoterID = fmt.Sprint(rand.Intn(10))
	reqBuff, err := json.Marshal(ballotvotereq)
	if err != nil {
		t.Fatalf("Failed to marshall post ballot request %+v", err)
	}

	_, result, err = httpClientRequest(http.MethodPost, net.JoinHostPort("", port), "/", bytes.NewReader(reqBuff))
	if err != nil {
		t.Fatalf("Failed to get ballot count resp:%s error:%+v", string(result), err)
	}
	log.Println("post ballot resp:", string(result))
	var postballotResp Status
	if err = json.Unmarshal(result, &postballotResp); err != nil {
		t.Fatalf("Failed to unmarshal post ballot response. %+v", err)
	}
	if postballotResp.Code != 201 {
		t.Error("post ballot resp status code")
	}

	_, result, err = httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		t.Fatalf("Failed to get final ballot count resp:%s error:%+v", string(result), err)
	}
	log.Println("get final ballot resp:", string(result))
	var finalRespData Response
	if err = json.Unmarshal(result, &finalRespData); err != nil {
		t.Fatalf("Failed to unmarshal get final ballot response. %+v", err)
	}
	if finalRespData.TotalVotes-initalRespData.TotalVotes != 1 {
		t.Error("ballot vote count error")
	}
}

func TestTestBallot_9bbe2d8b8d_Error(t *testing.T) {
	// Save current function and restore at the end:
	oldHttpClientRequest := httpClientRequest
	defer func() { httpClientRequest = oldHttpClientRequest }()

	// Mock to simulate error:
	httpClientRequest = func(method, url, path string, body io.Reader) (int, []byte, error) {
		return 0, nil, errors.New("Mock error")
	}

	TestBallot(t)
}
