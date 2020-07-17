package atcoder_test

import (
	"testing"

	"github.com/alvinpiter/cp-helper/repos/atcoder"
	"github.com/stretchr/testify/assert"
)

func TestToGeneralProblem(t *testing.T) {
	problem1 := atcoder.Problem{
		ContestID:  "abc123",
		ID:         "abc123_a",
		Title:      "Problem A",
		Difficulty: -123.12,
	}

	p1 := atcoder.ToGeneralProblem(problem1)
	assert.Equal(t, "abc123_a", p1.ID)
	assert.Equal(t, "Problem A", p1.Name)
	assert.Equal(t, 0, p1.Rating)
	assert.Equal(t, 0, len(p1.Tags))
	assert.Equal(t, "https://atcoder.jp/contests/abc123/tasks/abc123_a", p1.URL)

	problem2 := atcoder.Problem{
		ContestID:  "abc123",
		ID:         "abc123_b",
		Title:      "Problem B",
		Difficulty: 123.12,
	}

	p2 := atcoder.ToGeneralProblem(problem2)
	assert.Equal(t, 123, p2.Rating)
}

func TestToGeneralSubmission(t *testing.T) {
	prob := atcoder.Problem{
		ContestID: "abc123",
		ID:        "abc123_a",
	}

	acSubmission := atcoder.Submission{
		Problem: prob,
		Result:  "AC",
	}

	sub1 := atcoder.ToGeneralSubmission(acSubmission)
	assert.Equal(t, "abc123_a", sub1.Problem.ID)
	assert.Equal(t, true, sub1.IsAccepted)

	tleSubmission := atcoder.Submission{
		Problem: prob,
		Result:  "TLE",
	}

	sub2 := atcoder.ToGeneralSubmission(tleSubmission)
	assert.Equal(t, false, sub2.IsAccepted)
}
