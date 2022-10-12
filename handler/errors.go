package handler

import (
	"encoding/json"
	"fmt"
	"github.com/quangpham789/golang-assessment/utils"
	"net/http"
)

var (
	errNameCannotBeBlank  = RespError{Code: http.StatusBadRequest, Message: "name cannot be empty"}
	errEmailCannotBeBlank = RespError{Code: http.StatusBadRequest, Message: "email cannot be empty"}
	errInvalidEmail       = RespError{Code: http.StatusBadRequest, Message: "invalid email address"}
	errDataIsEmpty        = RespError{Code: http.StatusBadRequest, Message: "list of friend request is empty"}
)

type RespError struct {
	Code    int
	Message string
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
	res := utils.Response{
		Message: error.Message,
	}
	json.NewEncoder(w).Encode(res)
}
