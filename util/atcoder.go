package util

import (
	"github.com/alvinpiter/cp-helper/entities"
)

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

With this method, we intend to merge those responses and returns a nap of AtCoderProblems.
The map's key is the problem ID.
*/

func MergeAtCoderProblemResponse(problems []*entities.AtCoderProblem, problemDifficulty map[string]entities.AtCoderProblemDifficulty) map[string]entities.Problem {
	result := make(map[string]entities.Problem)

	for _, problem := range problems {
		difficulty, exist := problemDifficulty[problem.GetID()]
		if exist {
			problem.Difficulty = difficulty.Difficulty
		}

		result[problem.GetID()] = problem
	}

	return result
}
