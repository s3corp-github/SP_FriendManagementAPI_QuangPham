package rest

import (
	"encoding/json"
	"fmt"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/pkg/utils"
	"net/http"
)

var (
	ErrNameCannotBeBlank               = RespError{Code: http.StatusBadRequest, Message: "Name cannot be empty"}
	ErrEmailCannotBeBlank              = RespError{Code: http.StatusBadRequest, Message: "Email cannot be empty"}
	ErrInvalidEmail                    = RespError{Code: http.StatusBadRequest, Message: "Invalid email address"}
	ErrDataIsEmpty                     = RespError{Code: http.StatusBadRequest, Message: "List of friend request is empty"}
	ErrRequesterEmailAndAddresseeEmail = RespError{Code: http.StatusBadRequest, Message: "Requester email and target email must not be the same"}
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
