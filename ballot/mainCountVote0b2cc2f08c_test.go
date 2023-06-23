// Modified test case to avoid redeclaration errors

package main

import (
	"math/rand"
	"sort"
	"testing"
)

type TestCandidateVotes struct {
	CandidateID string
	Votes       int
}

type TestResultBoard struct {
	Results    []TestCandidateVotes
	TotalVotes int
}

func testGetCandidatesVote() map[string]int {
	return map[string]int{
		"A": rand.Intn(100),
		"B": rand.Intn(100),
		"C": rand.Intn(100),
		"D": rand.Intn(100),
	}
}

func testCountVote() (res TestResultBoard, err error) {
	votes := testGetCandidatesVote()
	for candidateID, votes := range votes {
		res.Results = append(res.Results, TestCandidateVotes{candidateID, votes})
		res.TotalVotes += votes
	}

	sort.Slice(res.Results, func(i, j int) bool {
		return res.Results[i].Votes > res.Results[j].Votes
	})
	return res, err
}

func TestCountVote(t *testing.T) {
	res, err := testCountVote()
	if err != nil {
		t.Error("Error occurred while counting votes:", err)
	}

	if len(res.Results) != 4 {
		t.Error("Expected 4 candidates, got", len(res.Results))
	}

	totalVotes := 0
	for _, candidate := range res.Results {
		totalVotes += candidate.Votes
	}

	if totalVotes != res.TotalVotes {
		t.Error("Total votes mismatch. Expected", totalVotes, "got", res.TotalVotes)
	}

	for i := 0; i < len(res.Results)-1; i++ {
		if res.Results[i].Votes < res.Results[i+1].Votes {
			t.Error("Results are not sorted in descending order by votes")
		}
	}
}

func TestCountVoteWithCustomCandidates(t *testing.T) {
	originalGetCandidatesVote := testGetCandidatesVote
	defer func() { testGetCandidatesVote = originalGetCandidatesVote }()
	testGetCandidatesVote = func() map[string]int {
		return map[string]int{
			"E": 20,
			"F": 40,
			"G": 60,
			"H": 80,
		}
	}

	res, err := testCountVote()
	if err != nil {
		t.Error("Error occurred while counting votes:", err)
	}

	if len(res.Results) != 4 {
		t.Error("Expected 4 candidates, got", len(res.Results))
	}

	if res.Results[0].CandidateID != "H" || res.Results[0].Votes != 80 {
		t.Error("Expected candidate H with 80 votes, got", res.Results[0].CandidateID, res.Results[0].Votes)
	}

	if res.Results[1].CandidateID != "G" || res.Results[1].Votes != 60 {
		t.Error("Expected candidate G with 60 votes, got", res.Results[1].CandidateID, res.Results[1].Votes)
	}

	if res.Results[2].CandidateID != "F" || res.Results[2].Votes != 40 {
		t.Error("Expected candidate F with 40 votes, got", res.Results[2].CandidateID, res.Results[2].Votes)
	}

	if res.Results[3].CandidateID != "E" || res.Results[3].Votes != 20 {
		t.Error("Expected candidate E with 20 votes, got", res.Results[3].CandidateID, res.Results[3].Votes)
	}
}