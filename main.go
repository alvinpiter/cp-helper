package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alvinpiter/cp-helper/api"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}

	return ":8000"
}

func main() {
	app := api.New()
	http.HandleFunc("/healthz", app.HealthzHandler)
	http.HandleFunc("/codeforces-problem-tags", app.CodeforcesProblemTagsHandler)
	http.HandleFunc("/compare", app.CompareHandler)

	port := getPort()
	fmt.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
