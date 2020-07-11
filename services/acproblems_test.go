package services_test

import (
	"testing"

	"github.com/alvinpiter/cp-helper/services"
	"github.com/stretchr/testify/assert"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/entities/mocks"
)

func TestGetAcceptedProblemsSuccess(t *testing.T) {
	oj := "codeforces"
	handle := "alvinpiter"

	problem1 := &entities.CodeforcesProblem{
		ContestID: 1,
		Index:     "A",
	}

	problem2 := &entities.CodeforcesProblem{
		ContestID: 1,
		Index:     "B",
	}

	submission1 := &entities.CodeforcesSubmission{
		Problem: problem1,
		Verdict: "OK",
	}

	submission2 := &entities.CodeforcesSubmission{
		Problem: problem1,
		Verdict: "OK",
	}

	submission3 := &entities.CodeforcesSubmission{
		Problem: problem2,
		Verdict: "WRONG_ANSWER",
	}

	submissions := []*entities.CodeforcesSubmission{submission1, submission2, submission3}

	cfRepo := new(mocks.CodeforcesRepository)
	cfRepo.On("GetSubmissions", handle).Return(submissions, nil)

	svc := services.NewService(cfRepo)

	result, _ := svc.GetAcceptedProblems(oj, handle)

	assert.Equal(t, 1, len(result))
	assert.Equal(t, "1A", result[0].GetID())
}

func TestGetAcceptedProblemsWithInvalidOj(t *testing.T) {
	oj := "uva"
	handle := "alvinpiter"

	cfRepo := new(mocks.CodeforcesRepository)

	svc := services.NewService(cfRepo)

	_, err := svc.GetAcceptedProblems(oj, handle)

	assert.True(t, err != nil)
}
