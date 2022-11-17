package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func LogRequest(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request path: %s", r.URL.Path)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

type RespError struct {
	Code    int
	Message string
}

type Response struct {
	Message string `json:"message"`
}

func ResponseJson(w http.ResponseWriter, httpCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(data)
}

func (r RespError) Error() string {
	return fmt.Sprintf("%d: %s", r.Code, r.Message)
}

func InternalServerError(msg string) error {
	return RespError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func JsonResponseError(w http.ResponseWriter, err error) {
	if _, ok := err.(RespError); !ok {
		err = InternalServerError(err.Error())
	}
	error, _ := err.(RespError)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(error.Code)
	res := Response{
		Message: error.Message,
	}
	json.NewEncoder(w).Encode(res)
}
