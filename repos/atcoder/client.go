/*
TODO:
* Fix dependency to API, i.e if API down, we shouldn't
*/

package atcoder

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alvinpiter/cp-helper/entities"
)

var atcoderAPI = "https://kenkoooo.com/atcoder/"

func NewRepository() *Repository {
	return &Repository{
		Client:   http.Client{},
		Problems: make(map[string]Problem),
	}
}

func (r *Repository) populateProblemsCache() error {
	urlProblemDetail := fmt.Sprintf("%sresources/merged-problems.json", atcoderAPI)
	resp1, err := r.Client.Get(urlProblemDetail)
	if err != nil {
		return err
	}

	problems := []Problem{}
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

	problemDifficulty := make(map[string]ProblemDifficulty)
	dec2 := json.NewDecoder(resp2.Body)
	err = dec2.Decode(&problemDifficulty)
	if err != nil {
		return err
	}

	r.Problems = mergeProblemResponse(problems, problemDifficulty)
	return nil
}

func (r *Repository) GetSubmissions(handle string) ([]entities.Submission, error) {
	if len(r.Problems) == 0 {
		err := r.populateProblemsCache()
		if err != nil {
			return nil, err
		}
	}

	url := fmt.Sprintf("%satcoder-api/results?user=%s", atcoderAPI, handle)

	resp, err := r.Client.Get(url)
	if err != nil {
		return nil, err
	}

	submissions := []Submission{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&submissions)
	if err != nil {
		return nil, err
	}

	subs := []entities.Submission{}
	for _, sub := range submissions {
		sub.Problem = r.Problems[sub.ProblemID]
		subs = append(subs, ToGeneralSubmission(sub))
	}

	return subs, nil
}
