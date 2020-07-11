package entities_test

import (
	"testing"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/stretchr/testify/assert"
)

func TestCodeforcesProblemsetProblem(t *testing.T) {
	problem := &entities.CodeforcesProblem{
		ContestID: 1234,
		Index:     "A",
		Name:      "First Problem",
		Rating:    2100,
		Tags:      []string{"implementation", "graph"},
	}

	assert.Equal(t, problem.GetID(), "1234A")
	assert.Equal(t, problem.GetName(), "First Problem")
	assert.Equal(t, problem.GetRating(), 2100)
	assert.Equal(t, problem.GetTags()[0], "implementation")
	assert.Equal(t, problem.GetTags()[1], "graph")
	assert.Equal(t, problem.GetURL(), "https://codeforces.com/contest/1234/problem/A")
}

func TestCodeforcesGymProblem(t *testing.T) {
	problem := &entities.CodeforcesProblem{
		ContestID: 102640,
		Index:     "A",
		Name:      "First Problem",
		Tags:      []string{},
	}

	assert.Equal(t, problem.GetID(), "102640A")
	assert.Equal(t, problem.GetName(), "First Problem")
	assert.Empty(t, problem.GetTags())
	assert.Equal(t, problem.GetURL(), "https://codeforces.com/gym/102640/problem/A")
}
