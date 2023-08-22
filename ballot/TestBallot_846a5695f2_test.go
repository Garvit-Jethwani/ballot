package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"testing"
)

func TestBallot(t *testing.T) {
	port := ""

	err := Ballot(port)
	if err != nil {
		t.Errorf("Test case 1 failed: %s", err.Error())
	}

	err = Ballot(port)
	if err == nil {
		t.Error("Test case 2 failed: expected an error but got nil")
	}
}

func Ballot(port string) error {
	_, result, err := httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		log.Printf("Failed to get ballot count resp:%s error:%+v", string(result), err)
		return err
	}

	var initalRespData Response
	if err = json.Unmarshal(result, &initalRespData); err != nil {
		log.Printf("Failed to unmarshal get ballot response. %+v", err)
		return err
	}

	var ballotvotereq Vote
	ballotvotereq.CandidateID = fmt.Sprint(rand.Intn(10))
	ballotvotereq.VoterID = fmt.Sprint(rand.Intn(10))
	reqBuff, err := json.Marshal(ballotvotereq)
	if err != nil {
		log.Printf("Failed to marshall post ballot request %+v", err)
		return err
	}

	_, result, err = httpClientRequest(http.MethodPost, net.JoinHostPort("", port), "/", bytes.NewReader(reqBuff))
	if err != nil {
		log.Printf("Failed to get ballot count resp:%s error:%+v", string(result), err)
		return err
	}

	var postballotResp Status
	if err = json.Unmarshal(result, &postballotResp); err != nil {
		log.Printf("Failed to unmarshal post ballot response. %+v", err)
		return err
	}
	if postballotResp.Code != 201 {
		return errors.New("post ballot resp status code")
	}

	_, result, err = httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		log.Printf("Failed to get final ballot count resp:%s error:%+v", string(result), err)
		return err
	}

	var finalRespData Response
	if err = json.Unmarshal(result, &finalRespData); err != nil {
		log.Printf("Failed to unmarshal get final ballot response. %+v", err)
		return err
	}
	if finalRespData.TotalVotes-initalRespData.TotalVotes != 1 {
		return errors.New("ballot vote count error")
	}
	return nil
}
