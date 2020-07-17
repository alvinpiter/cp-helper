package codeforces

import (
	"net/http"
)

type Repository struct {
	Client http.Client
}

type Problem struct {
	ContestID int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Rating    int      `json:"rating"`
	Tags      []string `json:"tags"`
}

type Submission struct {
	Problem Problem `json:"problem"`
	Verdict string  `json:"verdict"`
}

type UserStatus struct {
	Status  string       `json:"status"`
	Comment string       `json:"comment"`
	Result  []Submission `json:"result"`
}
