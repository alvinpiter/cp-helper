package services_test

import (
	"testing"

	"github.com/alvinpiter/cp-helper/services"
	"github.com/stretchr/testify/assert"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/entities/mocks"
)

func TestGetAcceptedProblemsCodeforces(t *testing.T) {
	handle := "alvinpiter"

	problem1 := entities.Problem{ID: "1A"}
	problem2 := entities.Problem{ID: "1B"}

	submission1 := entities.Submission{
		Problem:    problem1,
		IsAccepted: true,
	}

	submission2 := entities.Submission{
		Problem:    problem1,
		IsAccepted: true,
	}

	submission3 := entities.Submission{
		Problem:    problem2,
		IsAccepted: false,
	}

	submissions := []entities.Submission{submission1, submission2, submission3}

	cfRepo := new(mocks.Repository)
	atcRepo := new(mocks.Repository)
	cfRepo.On("GetSubmissions", handle).Return(submissions, nil)
	atcRepo.On("GetSubmissions", handle).Return(submissions, nil)

	svc := services.NewService()
	svc.CodeforcesRepo = cfRepo
	svc.AtCoderRepo = atcRepo

	result1, _ := svc.GetAcceptedProblems("codeforces", handle)
	assert.Equal(t, 1, len(result1))
	assert.Equal(t, "1A", result1[0].ID)

	result2, _ := svc.GetAcceptedProblems("atcoder", handle)
	assert.Equal(t, 1, len(result2))
	assert.Equal(t, "1A", result2[0].ID)
}

func TestGetAcceptedProblemsWithInvalidOj(t *testing.T) {
	handle := "alvinpiter"

	svc := services.NewService()

	_, err := svc.GetAcceptedProblems("uva", handle)

	assert.True(t, err != nil)
}
