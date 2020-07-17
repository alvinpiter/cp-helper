package api

import "github.com/alvinpiter/cp-helper/entities"

type ErrorResponse struct {
	Message string `json:"message"`
}

type CodeforcesProblemTagsResponse struct {
	ProblemTags []string `json:"problem_tags"`
}

type CompareResponse struct {
	Problems []entities.Problem `json:"problems"`
}
