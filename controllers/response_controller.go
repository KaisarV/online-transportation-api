package controllers

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, status int, response interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
