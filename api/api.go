package api

import (
	"net/http"

	"github.com/alvinpiter/cp-helper/entities"
	"github.com/alvinpiter/cp-helper/services"
)

type API struct {
	Service entities.Service
}

func New() *API {
	svc := services.New()
	return &API{Service: svc}
}

func (a *API) HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (a *API) CodeforcesProblemTagsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		writeJSON(w, http.StatusNotFound, nil)
		return
	}

	tags := []string{
		"constructive algorithms", "divide and conquer", "dfs and similar", "data structures", "binary search", "2-sat", "meet-in-the-middle", "schedules", "interactive", "implementation", "shortest paths", "fft", "games", "strings", "combinatorics", "bitmasks", "matrices", "number theory", "brute force", "dsu", "graph matchings", "*special", "geometry", "graphs", "trees", "two pointers", "dp", "probabilities", "hashing", "greedy", "string suffix structures", "expression parsing", "math", "sortings", "ternary search", "flows", "chinese remainder theorem",
	}

	resp := CodeforcesProblemTagsResponse{
		ProblemTags: tags,
	}

	writeJSON(w, http.StatusOK, resp)
}

func (a *API) CompareHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		return
	}

	if r.Method != "POST" {
		writeJSON(w, http.StatusNotFound, nil)
		return
	}

	body, err := NormalizeCompareRequest(r)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, ErrorResponse{Message: err.Error()})
		return
	}

	diffs, err := a.Service.CompareWithFilter(
		*body.OnlineJudge,
		*body.Handle,
		*body.RivalHandle,
		body.Filter,
	)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, ErrorResponse{Message: err.Error()})
		return
	}

	resp := CompareResponse{Problems: diffs}

	writeJSON(w, http.StatusOK, resp)
}
