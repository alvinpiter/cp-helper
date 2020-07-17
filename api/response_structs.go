package api

type ErrorResponse struct {
	Message string `json:"message"`
}

type CodeforcesProblemTagsResponse struct {
	ProblemTags []string `json:"problem_tags"`
}
