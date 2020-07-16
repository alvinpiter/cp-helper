package handlers

import (
	"errors"
	"net/http"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/util"
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

func CompareCodeforces(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		jsonError(w, errors.New("Invalid HTTP method"))
		return
	}

	params, err := util.NormalizeRequestBody("codeforces", req.Body)
	if err != nil {
		jsonError(w, err)
		return
	}

	handle := params.Handle
	rivalHandle := params.RivalHandle
	diffs, err := service.Compare("codeforces", handle, rivalHandle)
	if err != nil {
		jsonError(w, err)
		return
	}

	diffs = service.ApplyProblemFilter(diffs, params.Filter)

	jsonResponse(w, http.StatusOK, entities.NewCompareResponse(diffs))
}
