package main

import (
	"sync"
	"testing"
)

var (
	once               sync.Once
	candidateVotesStore map[string]int
)

func getCandidatesVote() map[string]int {
	once.Do(func() {
		candidateVotesStore = make(map[string]int)
	})
	return candidateVotesStore
}

func TestGetCandidatesVote_5ba0d06cb0(t *testing.T) {
	// Test case 1: Check initial state of the map
	votes := getCandidatesVote()
	if len(votes) != 0 {
		t.Error("Expected the vote store to be empty initially")
	}

	// Test case 2: Add some votes and check if the map is updated correctly
	candidateVotesStore["Candidate1"] = 5
	candidateVotesStore["Candidate2"] = 3
	votes = getCandidatesVote()
	if len(votes) != 2 {
		t.Error("Expected the vote store to have 2 candidates")
	}
	if votes["Candidate1"] != 5 {
		t.Error("Expected Candidate1 to have 5 votes")
	}
	if votes["Candidate2"] != 3 {
		t.Error("Expected Candidate2 to have 3 votes")
	}
}
