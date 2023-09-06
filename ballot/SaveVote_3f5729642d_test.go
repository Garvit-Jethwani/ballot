package main

import (
	"testing"
)

type Vote struct {
	CandidateID string
}

var candidateVotesStore = make(map[string]int)

func saveVote(vote Vote) error {
	candidateVotesStore[vote.CandidateID]++
	return nil
}

func TestSaveVote_3f5729642d(t *testing.T) {
	// Test case 1: Check if vote count increases after a vote is saved
	vote := Vote{CandidateID: "John"}
	err := saveVote(vote)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if candidateVotesStore[vote.CandidateID] != 1 {
		t.Error("Expected 1, got ", candidateVotesStore[vote.CandidateID])
	}

	// Test case 2: Check if vote count increases correctly after multiple votes are saved
	vote2 := Vote{CandidateID: "John"}
	err = saveVote(vote2)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if candidateVotesStore[vote.CandidateID] != 2 {
		t.Error("Expected 2, got ", candidateVotesStore[vote.CandidateID])
	}

	// Test case 3: Check if votes are saved correctly for different candidates
	vote3 := Vote{CandidateID: "Jane"}
	err = saveVote(vote3)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if candidateVotesStore[vote3.CandidateID] != 1 {
		t.Error("Expected 1, got ", candidateVotesStore[vote3.CandidateID])
	}
	if candidateVotesStore[vote.CandidateID] != 2 {
		t.Error("Expected 2, got ", candidateVotesStore[vote.CandidateID])
	}
}
