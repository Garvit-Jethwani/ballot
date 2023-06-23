package main

import (
	"sync"
	"testing"
)

type BallotVote struct {
	CandidateID string
}

var ballotCandidateVotesStore = make(map[string]int)
var ballotCandidateVotesStoreLock = &sync.Mutex{}

func saveBallotVote(vote BallotVote) error {
	ballotCandidateVotesStoreLock.Lock()
	defer ballotCandidateVotesStoreLock.Unlock()

	ballotCandidateVotesStore[vote.CandidateID]++
	return nil
}

func TestSaveVote31e2883bdc(t *testing.T) {
	// Test case 1: Save a valid vote and check if the vote count is incremented
	t.Run("SaveValidVote", func(t *testing.T) {
		candidateID := "candidate1"
		initialVoteCount := ballotCandidateVotesStore[candidateID]
		vote := BallotVote{CandidateID: candidateID}
		err := saveBallotVote(vote)

		if err != nil {
			t.Error("Expected no error, got:", err)
		}

		if ballotCandidateVotesStore[candidateID] != initialVoteCount+1 {
			t.Errorf("Expected vote count to be %d, got %d", initialVoteCount+1, ballotCandidateVotesStore[candidateID])
		}
	})

	// Test case 2: Save multiple votes for different candidates and check if the vote counts are incremented correctly
	t.Run("SaveMultipleVotes", func(t *testing.T) {
		candidates := []string{"candidate2", "candidate3"}
		initialVoteCounts := make(map[string]int)

		for _, candidate := range candidates {
			initialVoteCounts[candidate] = ballotCandidateVotesStore[candidate]
			vote := BallotVote{CandidateID: candidate}
			err := saveBallotVote(vote)

			if err != nil {
				t.Error("Expected no error, got:", err)
			}
		}

		for _, candidate := range candidates {
			if ballotCandidateVotesStore[candidate] != initialVoteCounts[candidate]+1 {
				t.Errorf("Expected vote count for %s to be %d, got %d", candidate, initialVoteCounts[candidate]+1, ballotCandidateVotesStore[candidate])
			}
		}
	})
}