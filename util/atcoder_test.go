package util_test

import (
	"testing"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/util"
	"github.com/stretchr/testify/assert"
)

func TestMergeAtCoderProblemResponse(t *testing.T) {
	problem1 := &entities.AtCoderProblem{ID: "abc1"}
	problem2 := &entities.AtCoderProblem{ID: "abc2"}
	problem3 := &entities.AtCoderProblem{ID: "abc3"}

	problems := []*entities.AtCoderProblem{problem1, problem2, problem3}

	problemDifficulty := map[string]entities.AtCoderProblemDifficulty{
		"abc2": entities.AtCoderProblemDifficulty{
			Difficulty: 1234.12,
		},
		"abc1": entities.AtCoderProblemDifficulty{
			Difficulty: 4321.12,
		},
	}

	result := util.MergeAtCoderProblemResponse(problems, problemDifficulty)

	assert.Equal(t, 3, len(result))
	assert.Equal(t, "abc1", result[0].GetID())
	assert.Equal(t, 4321, result[0].GetRating())
	assert.Equal(t, "abc2", result[1].GetID())
	assert.Equal(t, 1234, result[1].GetRating())
	assert.Equal(t, "abc3", result[2].GetID())
	assert.Equal(t, 0, result[2].GetRating())

}
