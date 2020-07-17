package main

import (
	"log"
	"net/http"

	"github.com/alvinpiter/cp-helper/api"
)

func main() {
	app := api.New()
	http.HandleFunc("/healthz", app.HealthzHandler)
	http.HandleFunc("/codeforces-problem-tags", app.CodeforcesProblemTagsHandler)
	http.HandleFunc("/compare", app.CompareHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
