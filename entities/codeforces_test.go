package entities_test

import (
	"testing"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/stretchr/testify/assert"
)

func TestCodeforcesFilterParameter(t *testing.T) {
	tagsFilter := &entities.TagsFilterParameter{
		Mode:   "and",
		Values: []string{"implementation"},
	}

	low := 1000
	high := 2000
	ratingFilter := &entities.RatingFilterParameter{
		Minimum: &low,
		Maximum: &high,
	}

	cfFilterParameter := &entities.CodeforcesFilterParameter{
		TagsFilter:   tagsFilter,
		RatingFilter: ratingFilter,
	}

	assert.Equal(t, tagsFilter, cfFilterParameter.GetTagsFilterParameter())
	assert.Equal(t, ratingFilter, cfFilterParameter.GetRatingFilterParameter())
}
