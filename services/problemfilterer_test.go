package services_test

import (
	"testing"

	"github.com/alvinpiter/cp-helper/services"
	"github.com/stretchr/testify/assert"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/entities/mocks"
)

func TestApplyProblemFilter(t *testing.T) {
	problem1 := &entities.CodeforcesProblem{
		ContestID: 1,
		Index:     "A",
		Rating:    1900,
		Tags:      []string{"implementation", "math"},
	}

	problem2 := &entities.CodeforcesProblem{
		ContestID: 1,
		Index:     "B",
		Rating:    2000,
		Tags:      []string{"implementation", "math"},
	}

	problem3 := &entities.CodeforcesProblem{
		ContestID: 1,
		Index:     "C",
		Rating:    2100,
		Tags:      []string{"implementation"},
	}

	problems := []entities.Problem{problem1, problem2, problem3}

	low := 1900
	high := 2000

	filter1 := &entities.CodeforcesFilterParameter{
		TagsFilter: &entities.TagsFilterParameter{
			Mode:   "and",
			Values: []string{"implementation", "math"},
		},
		RatingFilter: &entities.RatingFilterParameter{
			Maximum: &low,
		},
	}

	filter2 := &entities.CodeforcesFilterParameter{
		TagsFilter: &entities.TagsFilterParameter{
			Mode:   "and",
			Values: []string{"implementation", "math"},
		},
		RatingFilter: &entities.RatingFilterParameter{
			Minimum: &high,
		},
	}

	filter3 := &entities.CodeforcesFilterParameter{
		TagsFilter: &entities.TagsFilterParameter{
			Mode:   "or",
			Values: []string{"implementation", "math"},
		},
	}

	filter4 := &entities.CodeforcesFilterParameter{
		RatingFilter: &entities.RatingFilterParameter{
			Minimum: &low,
			Maximum: &high,
		},
	}

	cfRepo := new(mocks.Repository)
	svc := services.NewService()
	svc.CodeforcesRepo = cfRepo

	result1 := svc.ApplyProblemFilter(problems, filter1)
	assert.Equal(t, 1, len(result1))
	assert.Equal(t, "1A", result1[0].GetID())

	result2 := svc.ApplyProblemFilter(problems, filter2)
	assert.Equal(t, 1, len(result2))
	assert.Equal(t, "1B", result2[0].GetID())

	result3 := svc.ApplyProblemFilter(problems, filter3)
	assert.Equal(t, 3, len(result3))
	assert.Equal(t, "1A", result3[0].GetID())
	assert.Equal(t, "1B", result3[1].GetID())
	assert.Equal(t, "1C", result3[2].GetID())

	result4 := svc.ApplyProblemFilter(problems, filter4)
	assert.Equal(t, 2, len(result4))
	assert.Equal(t, "1A", result4[0].GetID())
	assert.Equal(t, "1B", result4[1].GetID())
}
