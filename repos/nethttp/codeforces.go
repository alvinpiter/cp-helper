package nethttp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/alvinpiter/cp-helper/entities"
)

var codeforcesAPI = "https://codeforces.com/api/"

/*
CodeforcesRepository is a struct that represents an object that
communicates with codeforces API.
*/
type CodeforcesRepository struct {
	Client http.Client
}

func NewCodeforcesRepository() *CodeforcesRepository {
	return &CodeforcesRepository{
		Client: http.Client{},
	}
}

func (r *CodeforcesRepository) GetSubmissions(handle string) ([]entities.Submission, error) {
	type userStatusResponse struct {
		Status  string                           `json:"status"`
		Comment string                           `json:"comment"`
		Result  []*entities.CodeforcesSubmission `json:"result"`
	}

	url := fmt.Sprintf("%suser.status?handle=%s", codeforcesAPI, handle)

	resp, err := r.Client.Get(url)
	if err != nil {
		return nil, err
	}

	respObj := userStatusResponse{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&respObj)
	if err != nil {
		return nil, err
	}

	if respObj.Status != "OK" {
		return nil, errors.New(respObj.Comment)
	}

	submissions := []entities.Submission{}
	for _, s := range respObj.Result {
		submissions = append(submissions, s)
	}

	return submissions, nil
}

func (r *CodeforcesRepository) GetProblems() ([]entities.Problem, error) {
	type problemsetProblemsResultResponse struct {
		Problems []*entities.CodeforcesProblem `json:"problems"`
	}

	type problemsetProblemsResponse struct {
		Status  string                           `json:"status"`
		Comment string                           `json:"comment"`
		Result  problemsetProblemsResultResponse `json:"result"`
	}

	url := fmt.Sprintf("%sproblemset.problems", codeforcesAPI)

	resp, err := r.Client.Get(url)
	if err != nil {
		return nil, err
	}

	respObj := problemsetProblemsResponse{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&respObj)
	if err != nil {
		return nil, err
	}

	if respObj.Status != "OK" {
		return nil, errors.New(respObj.Comment)
	}

	problems := []entities.Problem{}
	for _, p := range respObj.Result.Problems {
		problems = append(problems, p)
	}

	return problems, nil
}
