package entities_test

import (
	"testing"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/stretchr/testify/assert"
)

func TestAtCoderFilterParameter(t *testing.T) {
	low := 1000
	high := 2000

	ratingFilter := &entities.RatingFilterParameter{
		Minimum: &low,
		Maximum: &high,
	}

	atcFilterParameter := &entities.AtCoderFilterParameter{
		RatingFilter: ratingFilter,
	}

	//TODO: Assert tag filter is nil
	assert.Equal(t, ratingFilter, atcFilterParameter.GetRatingFilterParameter())
}
