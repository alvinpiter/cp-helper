package atcoder

import (
	"fmt"

	"github.com/alvinpiter/cp-helper/entities"
)

var problemURLFormat = "https://atcoder.jp/contests/%s/tasks/%s"

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

/*
AtCoder doesn't provide a single endpoint that returns problem detail and
its difficulty, hence we need to hit two endpoints and merge the responses.

The endpoint for problem detail is https://kenkoooo.com/atcoder/resources/merged-problems.json,
and it has response like:
[
	{
		"id": ...,
		"contest_id": ...,
		"title": ...
	}
]

The endpoint for problem difficulty is https://kenkoooo.com/atcoder/resources/problem-models.json,
and it has response like:
{
	<problem_id>: {
		difficulty: ...
	}
}

With this method, we intend to merge those responses and returns a map of Problem.
The map's key is the problem ID.
*/
func mergeProblemResponse(problems []Problem, problemDifficulty map[string]ProblemDifficulty) map[string]Problem {
	result := make(map[string]Problem)

	for _, problem := range problems {
		difficulty, exist := problemDifficulty[problem.ID]
		if exist {
			problem.Difficulty = difficulty.Difficulty
		}

		result[problem.ID] = problem
	}

	return result
}
