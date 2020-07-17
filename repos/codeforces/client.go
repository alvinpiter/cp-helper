package codeforces

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/alvinpiter/cp-helper/entities"
)

var codeforcesAPI = "https://codeforces.com/api/"

func NewRepository() *Repository {
	return &Repository{
		Client: http.Client{},
	}
}

func (r *Repository) GetSubmissions(handle string) ([]entities.Submission, error) {
	url := fmt.Sprintf("%suser.status?handle=%s", codeforcesAPI, handle)

	resp, err := r.Client.Get(url)
	if err != nil {
		return nil, err
	}

	respObj := UserStatus{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&respObj)
	if err != nil {
		return nil, err
	}

	if respObj.Status != "OK" {
		return nil, errors.New(respObj.Comment)
	}

	submissions := []entities.Submission{}
	for _, sub := range respObj.Result {
		submissions = append(submissions, ToGeneralSubmission(sub))
	}

	return submissions, nil
}
