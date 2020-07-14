package services_test

import (
	"testing"

	"github.com/alvinpiter/cp-helper/services"
	"github.com/stretchr/testify/assert"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/entities/mocks"
)

func TestGetAcceptedProblemsCodeforces(t *testing.T) {
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

	submissions := []entities.Submission{submission1, submission2, submission3}

	cfRepo := new(mocks.Repository)
	cfRepo.On("GetSubmissions", handle).Return(submissions, nil)

	svc := services.NewService()
	svc.CodeforcesRepo = cfRepo

	result, _ := svc.GetAcceptedProblems(oj, handle)

	assert.Equal(t, 1, len(result))
	assert.Equal(t, "1A", result[0].GetID())
}

func TestGetAcceptedProblemsAtCoder(t *testing.T) {
	oj := "atcoder"
	handle := "alvinpiter"

	problem1 := &entities.AtCoderProblem{ID: "atc1"}
	problem2 := &entities.AtCoderProblem{ID: "atc2"}

	submission1 := &entities.AtCoderSubmission{
		Problem: problem1,
		Result:  "AC",
	}

	submission2 := &entities.AtCoderSubmission{
		Problem: problem1,
		Result:  "AC",
	}

	submission3 := &entities.AtCoderSubmission{
		Problem: problem2,
		Result:  "TLE",
	}

	submissions := []entities.Submission{submission1, submission2, submission3}

	atcRepo := new(mocks.Repository)
	atcRepo.On("GetSubmissions", handle).Return(submissions, nil)

	svc := services.NewService()
	svc.AtCoderRepo = atcRepo

	result, _ := svc.GetAcceptedProblems(oj, handle)

	assert.Equal(t, 1, len(result))
	assert.Equal(t, "atc1", result[0].GetID())
}

func TestGetAcceptedProblemsWithInvalidOj(t *testing.T) {
	oj := "uva"
	handle := "alvinpiter"

	cfRepo := new(mocks.Repository)

	svc := services.NewService()
	svc.CodeforcesRepo = cfRepo

	_, err := svc.GetAcceptedProblems(oj, handle)

	assert.True(t, err != nil)
}
