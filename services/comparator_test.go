package services_test

import (
	"testing"

	"github.com/alvinpiter/cp-helper/services"
	"github.com/stretchr/testify/assert"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/entities/mocks"
)

func TestCompare(t *testing.T) {
	handle1 := "handle1"
	handle2 := "handle2"

	problem1 := &entities.CodeforcesProblem{
		ContestID: 1,
		Index:     "A",
	}

	problem2 := &entities.CodeforcesProblem{
		ContestID: 1,
		Index:     "B",
	}

	problem3 := &entities.CodeforcesProblem{
		ContestID: 1,
		Index:     "C",
	}

	submissions1 := []entities.Submission{
		&entities.CodeforcesSubmission{
			Problem: problem1,
			Verdict: "OK",
		},
		&entities.CodeforcesSubmission{
			Problem: problem2,
			Verdict: "OK",
		},
	}

	submissions2 := []entities.Submission{
		&entities.CodeforcesSubmission{
			Problem: problem2,
			Verdict: "OK",
		},
		&entities.CodeforcesSubmission{
			Problem: problem3,
			Verdict: "OK",
		},
	}

	cfRepo := new(mocks.Repository)
	cfRepo.On("GetSubmissions", handle1).Return(submissions1, nil)
	cfRepo.On("GetSubmissions", handle2).Return(submissions2, nil)

	svc := services.NewService()
	svc.CodeforcesRepo = cfRepo

	result, _ := svc.Compare("codeforces", handle1, handle2)

	assert.Equal(t, 1, len(result))
	assert.Equal(t, "1C", result[0].GetID())
}
