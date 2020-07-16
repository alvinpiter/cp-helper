package main

import (
	"log"
	"net/http"

	"github.com/alvinpiter/cp-helper/handlers"
)

func main() {
	http.HandleFunc("/codeforces-problem-tags", handlers.GetCodeforcesProblemTags)
	http.HandleFunc("/compare-codeforces", handlers.CompareCodeforces)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
