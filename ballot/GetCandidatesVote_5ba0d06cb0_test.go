package main

import (
	"sync"
	"testing"
)

var once sync.Once
var candidateVotesStore map[string]int

func getCandidatesVote() map[string]int {
	once.Do(func() {
		candidateVotesStore = make(map[string]int)
	})
	return candidateVotesStore
}

func TestGetCandidatesVote(t *testing.T) {
	// Test case 1: Check if the candidateVotesStore is initialized as an empty map
	votes := getCandidatesVote()
	if len(votes) != 0 {
		t.Error("Expected an empty map, but got a non-empty map")
	}

	// Test case 2: Check if the candidateVotesStore is initialized only once
	votes1 := getCandidatesVote()
	votes2 := getCandidatesVote()
	if &votes1 != &votes2 {
		t.Error("Expected the same instance of candidateVotesStore, but got different instances")
	}
}
