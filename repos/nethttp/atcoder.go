/*
TODO:
* Fix dependency to API, i.e if API down, we shouldn't
*/

package nethttp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/util"
)

var atcoderAPI = "https://kenkoooo.com/atcoder/"

type AtCoderRepository struct {
	Client   http.Client
	Problems map[string]entities.Problem
}

func NewAtCoderRespository() *AtCoderRepository {
	return &AtCoderRepository{
		Client:   http.Client{},
		Problems: make(map[string]entities.Problem),
	}
}

func (r *AtCoderRepository) cacheProblems() error {
	urlProblemDetail := fmt.Sprintf("%sresources/merged-problems.json", atcoderAPI)
	resp1, err := r.Client.Get(urlProblemDetail)
	if err != nil {
		return err
	}

	problems := []*entities.AtCoderProblem{}
	dec1 := json.NewDecoder(resp1.Body)
	err = dec1.Decode(&problems)
	if err != nil {
		return err
	}

	urlProblemDifficulty := fmt.Sprintf("%sresources/problem-models.json", atcoderAPI)
	resp2, err := r.Client.Get(urlProblemDifficulty)
	if err != nil {
		return err
	}

	problemDifficulty := make(map[string]entities.AtCoderProblemDifficulty)
	dec2 := json.NewDecoder(resp2.Body)
	err = dec2.Decode(&problemDifficulty)
	if err != nil {
		return err
	}

	r.Problems = util.MergeAtCoderProblemResponse(problems, problemDifficulty)
	return nil
}

func (r *AtCoderRepository) GetProblems() ([]entities.Problem, error) {
	if len(r.Problems) == 0 {
		err := r.cacheProblems()
		if err != nil {
			return nil, err
		}
	}

	result := []entities.Problem{}
	for _, problem := range r.Problems {
		result = append(result, problem)
	}

	return result, nil
}

func (r *AtCoderRepository) GetSubmissions(handle string) ([]entities.Submission, error) {
	if len(r.Problems) == 0 {
		err := r.cacheProblems()
		if err != nil {
			return nil, err
		}
	}

	url := fmt.Sprintf("%satcoder-api/results?user=%s", atcoderAPI, handle)

	resp, err := r.Client.Get(url)
	if err != nil {
		return nil, err
	}

	submissions := []*entities.AtCoderSubmission{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&submissions)
	if err != nil {
		return nil, err
	}

	subs := []entities.Submission{}
	for _, sub := range submissions {
		sub.Problem = r.Problems[sub.ProblemID].(*entities.AtCoderProblem)
		subs = append(subs, sub)
	}

	return subs, nil
}
