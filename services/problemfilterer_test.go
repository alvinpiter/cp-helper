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
		Tags:      []string{"implementation", "math"},
	}

	problem4 := &entities.CodeforcesProblem{
		ContestID: 1,
		Index:     "D",
		Rating:    2000,
		Tags:      []string{"implementation"},
	}

	problems := []entities.Problem{problem1, problem2, problem3, problem4}

	filter1 := map[string]interface{}{
		"rating": map[string]int{
			"minimum": 2000,
			"maximum": 2100,
		},
		"id": map[string]interface{}{
			"mode":   "exclude",
			"values": []string{"1B"},
		},
		"tag": map[string]interface{}{
			"mode":   "or",
			"values": []string{"implementation", "math"},
		},
	}

	filter2 := map[string]interface{}{
		"rating": map[string]int{
			"minimum": 2000,
			"maximum": 2100,
		},
		"id": map[string]interface{}{
			"mode":   "exclude",
			"values": []string{"1B"},
		},
		"tag": map[string]interface{}{
			"mode":   "and",
			"values": []string{"implementation", "math"},
		},
	}

	cfRepo := new(mocks.CodeforcesRepository)
	svc := services.NewService(cfRepo)

	result1 := svc.ApplyProblemFilter(problems, filter1)
	assert.Equal(t, 2, len(result1))
	assert.Equal(t, "1C", result1[0].GetID())
	assert.Equal(t, "1D", result1[1].GetID())

	result2 := svc.ApplyProblemFilter(problems, filter2)
	assert.Equal(t, 1, len(result2))
	assert.Equal(t, "1C", result2[0].GetID())
}
