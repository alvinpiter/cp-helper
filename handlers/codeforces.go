package handlers

import (
	"net/http"

	"github.com/alvinpiter/cp-helper/entities"
)

func GetCodeforcesProblemTags(w http.ResponseWriter, req *http.Request) {
	tags := []string{
		"constructive algorithms", "divide and conquer", "dfs and similar", "data structures", "binary search", "2-sat", "meet-in-the-middle", "schedules", "interactive", "implementation", "shortest paths", "fft", "games", "strings", "combinatorics", "bitmasks", "matrices", "number theory", "brute force", "dsu", "graph matchings", "*special", "geometry", "graphs", "trees", "two pointers", "dp", "probabilities", "hashing", "greedy", "string suffix structures", "expression parsing", "math", "sortings", "ternary search", "flows", "chinese remainder theorem",
	}

	resp := entities.CodeforcesProblemTagsResponse{
		ProblemTags: tags,
	}

	jsonResponse(w, http.StatusOK, resp)
}
