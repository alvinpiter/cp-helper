package atcoder

import (
	"fmt"

	"github.com/alvinpiter/cp-helper/entities"
)

var problemURLFormat = "https://atcoder.jp/contests/%s/tasks/%s"

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

func ToGeneralProblem(p Problem) entities.Problem {
	rating := int(p.Difficulty)
	if rating < 0 {
		rating = 0
	}

	return entities.Problem{
		ID:     p.ID,
		Name:   p.Title,
		Rating: rating,
		Tags:   []string{},
		URL:    fmt.Sprintf(problemURLFormat, p.ContestID, p.ID),
	}
}

func ToGeneralSubmission(s Submission) entities.Submission {
	accepted := false
	if s.Result == "AC" {
		accepted = true
	}

	return entities.Submission{
		Problem:    ToGeneralProblem(s.Problem),
		IsAccepted: accepted,
	}
}
