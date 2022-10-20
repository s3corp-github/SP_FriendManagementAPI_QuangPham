package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func JsonResponse(w http.ResponseWriter, httpCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(data)
}
