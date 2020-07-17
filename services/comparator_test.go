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

	problem1 := entities.Problem{ID: "1A"}
	problem2 := entities.Problem{ID: "1B"}
	problem3 := entities.Problem{ID: "1C"}

	submissions1 := []entities.Submission{
		entities.Submission{
			Problem:    problem1,
			IsAccepted: true,
		},
		entities.Submission{
			Problem:    problem2,
			IsAccepted: true,
		},
	}

	submissions2 := []entities.Submission{
		entities.Submission{
			Problem:    problem2,
			IsAccepted: true,
		},
		entities.Submission{
			Problem:    problem3,
			IsAccepted: true,
		},
		entities.Submission{
			Problem:    problem3,
			IsAccepted: true,
		},
	}

	cfRepo := new(mocks.Repository)
	cfRepo.On("GetSubmissions", handle1).Return(submissions1, nil)
	cfRepo.On("GetSubmissions", handle2).Return(submissions2, nil)

	svc := services.New()
	svc.CodeforcesRepo = cfRepo

	result, _ := svc.Compare("codeforces", handle1, handle2)

	assert.Equal(t, 1, len(result))
	assert.Equal(t, "1C", result[0].ID)

	_, err := svc.Compare("uva", handle1, handle2)
	assert.NotNil(t, err)
}
