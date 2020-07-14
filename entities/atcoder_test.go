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
}
