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

	assert.Equal(t, "1234A", problem.GetID())
	assert.Equal(t, "First Problem", problem.GetName())
	assert.Equal(t, 2100, problem.GetRating())
	assert.Equal(t, "implementation", problem.GetTags()[0])
	assert.Equal(t, "graph", problem.GetTags()[1])
	assert.Equal(t, "https://codeforces.com/contest/1234/problem/A", problem.GetURL())
}

func TestCodeforcesGymProblem(t *testing.T) {
	problem := &entities.CodeforcesProblem{
		ContestID: 102640,
		Index:     "A",
		Name:      "First Problem",
		Tags:      []string{},
	}

	assert.Equal(t, "102640A", problem.GetID())
	assert.Equal(t, "First Problem", problem.GetName())
	assert.Empty(t, problem.GetTags())
	assert.Equal(t, "https://codeforces.com/gym/102640/problem/A", problem.GetURL())
}

func TestCodeforcesSubmission(t *testing.T) {
	problem := &entities.CodeforcesProblem{
		ContestID: 1234,
		Index:     "A",
	}

	acSubmission := &entities.CodeforcesSubmission{
		Problem: problem,
		Verdict: "OK",
	}

	assert.Equal(t, problem, acSubmission.GetProblem())
	assert.Equal(t, true, acSubmission.IsAccepted())

	waSubmission := &entities.CodeforcesSubmission{
		Problem: problem,
		Verdict: "WRONG_ANSWER",
	}

	assert.Equal(t, problem, waSubmission.GetProblem())
	assert.Equal(t, false, waSubmission.IsAccepted())
}

func TestCodeforcesFilterParameter(t *testing.T) {
	tagsFilter := &entities.TagsFilterParameter{
		Mode:   "and",
		Values: []string{"implementation"},
	}

	ratingFilter := &entities.RatingFilterParameter{
		Minimum: 1000,
		Maximum: 2000,
	}

	cfFilterParameter := &entities.CodeforcesFilterParameter{
		TagsFilter:   tagsFilter,
		RatingFilter: ratingFilter,
	}

	assert.Equal(t, tagsFilter, cfFilterParameter.GetTagsFilterParameter())
	assert.Equal(t, ratingFilter, cfFilterParameter.GetRatingFilterParameter())
}
