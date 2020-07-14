package entities_test

import (
	"testing"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/stretchr/testify/assert"
)

func TestAtCoderProblem(t *testing.T) {
	problem := &entities.AtCoderProblem{
		ID:         "abc001_1",
		ContestID:  "abc001",
		Title:      "Problem 1",
		Difficulty: 1234.85,
	}

	assert.Equal(t, "abc001_1", problem.GetID())
	assert.Equal(t, "Problem 1", problem.GetName())
	assert.Equal(t, 1234, problem.GetRating())
	assert.Equal(t, 0, len(problem.GetTags()))
	assert.Equal(t, "https://atcoder.jp/contests/abc001/tasks/abc001_1", problem.GetURL())

	problem.Difficulty = -1
	assert.Equal(t, 0, problem.GetRating())
}

func TestAtCoderSubmission(t *testing.T) {
	problem := &entities.AtCoderProblem{
		ID:         "abc001_1",
		ContestID:  "abc001",
		Title:      "Problem 1",
		Difficulty: 1234.85,
	}

	acSubmission := &entities.AtCoderSubmission{
		ProblemID: problem.ID,
		Problem:   problem,
		Result:    "AC",
	}

	assert.Equal(t, problem, acSubmission.GetProblem())
	assert.Equal(t, true, acSubmission.IsAccepted())

	tleSubmission := &entities.AtCoderSubmission{
		ProblemID: problem.ID,
		Problem:   problem,
		Result:    "TLE",
	}

	assert.Equal(t, problem, tleSubmission.GetProblem())
	assert.Equal(t, false, tleSubmission.IsAccepted())
}
