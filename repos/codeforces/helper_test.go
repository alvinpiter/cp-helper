package codeforces_test

import (
	"testing"

	"github.com/alvinpiter/cp-helper/repos/codeforces"
	"github.com/stretchr/testify/assert"
)

func TestToGeneralProblem(t *testing.T) {
	problemsetProblem := codeforces.Problem{
		ContestID: 1234,
		Index:     "A",
		Name:      "Problem A",
		Rating:    1500,
		Tags:      []string{"implementation"},
	}

	p1 := codeforces.ToGeneralProblem(problemsetProblem)
	assert.Equal(t, "1234A", p1.ID)
	assert.Equal(t, "Problem A", p1.Name)
	assert.Equal(t, 1500, p1.Rating)
	assert.Equal(t, 1, len(p1.Tags))
	assert.Equal(t, "implementation", p1.Tags[0])
	assert.Equal(t, "https://codeforces.com/contest/1234/problem/A", p1.URL)

	gymProblem := codeforces.Problem{
		ContestID: 102644,
		Index:     "A",
		Name:      "Problem A",
	}

	p2 := codeforces.ToGeneralProblem(gymProblem)
	assert.Equal(t, 0, p2.Rating)
	assert.Equal(t, 0, len(p2.Tags))
	assert.Equal(t, "https://codeforces.com/gym/102644/problem/A", p2.URL)
}

func TestToGeneralSubmission(t *testing.T) {
	prob := codeforces.Problem{
		ContestID: 1234,
		Index:     "A",
	}

	acSubmission := codeforces.Submission{
		Problem: prob,
		Verdict: "OK",
	}

	sub1 := codeforces.ToGeneralSubmission(acSubmission)
	assert.Equal(t, "1234A", sub1.Problem.ID)
	assert.Equal(t, true, sub1.IsAccepted)

	waSubmission := codeforces.Submission{
		Problem: prob,
		Verdict: "WRONG_ANSWER",
	}

	sub2 := codeforces.ToGeneralSubmission(waSubmission)
	assert.Equal(t, false, sub2.IsAccepted)
}
