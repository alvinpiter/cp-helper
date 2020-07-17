package api

import "net/http"

type API struct {
	Mux *http.ServeMux
}

func New() *API {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", HealthzHandler)
	mux.HandleFunc("/codeforces-problem-tags", CodeforcesProblemTagsHandler)

	return &API{
		Mux: mux,
	}
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func CodeforcesProblemTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags := []string{
		"constructive algorithms", "divide and conquer", "dfs and similar", "data structures", "binary search", "2-sat", "meet-in-the-middle", "schedules", "interactive", "implementation", "shortest paths", "fft", "games", "strings", "combinatorics", "bitmasks", "matrices", "number theory", "brute force", "dsu", "graph matchings", "*special", "geometry", "graphs", "trees", "two pointers", "dp", "probabilities", "hashing", "greedy", "string suffix structures", "expression parsing", "math", "sortings", "ternary search", "flows", "chinese remainder theorem",
	}

	resp := CodeforcesProblemTagsResponse{
		ProblemTags: tags,
	}

	writeJSON(w, http.StatusOK, resp)
}
