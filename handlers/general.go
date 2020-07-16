package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alvinpiter/cp-helper/services"
)

var service *services.Service

func init() {
	service = services.NewService()
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	jsonData, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}
