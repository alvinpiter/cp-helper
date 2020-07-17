package codeforces

import (
	"fmt"

	"github.com/alvinpiter/cp-helper/entities"
)

var (
	gymProblemURLFormat        = "https://codeforces.com/gym/%d/problem/%s"
	problemsetProblemURLFormat = "https://codeforces.com/contest/%d/problem/%s"
)

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

func ToGeneralProblem(p Problem) entities.Problem {
	var problemURL string
	if p.ContestID >= 100000 {
		problemURL = fmt.Sprintf(gymProblemURLFormat, p.ContestID, p.Index)
	} else {
		problemURL = fmt.Sprintf(problemsetProblemURLFormat, p.ContestID, p.Index)
	}

	return entities.Problem{
		ID:     fmt.Sprintf("%d%s", p.ContestID, p.Index),
		Name:   p.Name,
		Rating: p.Rating,
		Tags:   p.Tags,
		URL:    problemURL,
	}
}

func ToGeneralSubmission(s Submission) entities.Submission {
	accepted := false
	if s.Verdict == "OK" {
		accepted = true
	}

	return entities.Submission{
		Problem:    ToGeneralProblem(s.Problem),
		IsAccepted: accepted,
	}
}
