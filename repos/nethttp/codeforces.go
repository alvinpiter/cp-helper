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

func (r *CodeforcesRepository) GetSubmissions(handle string) ([]*entities.CodeforcesSubmission, error) {
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

	return respObj.Result, nil
}

func (r *CodeforcesRepository) GetProblems() ([]*entities.CodeforcesProblem, error) {
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

	return respObj.Result.Problems, nil
}
