package atcoder

import (
	"net/http"
)

type Repository struct {
	Client   http.Client
	Problems map[string]Problem
}

type Problem struct {
	ID         string `json:"id"`
	ContestID  string `json:"contest_id"`
	Title      string `json:"title"`
	Difficulty float64
}

type Submission struct {
	ProblemID string `json:"problem_id"`
	Problem   Problem
	Result    string `json:"result"`
}

type ProblemDifficulty struct {
	Difficulty float64 `json:"difficulty"`
}
