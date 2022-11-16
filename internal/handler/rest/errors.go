package rest

import (
	"net/http"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
)

var (
	ErrNameInvalid                     = utils.RespError{Code: http.StatusBadRequest, Message: "Name cannot be empty"}
	ErrEmailCannotBeBlank              = utils.RespError{Code: http.StatusBadRequest, Message: "Email cannot be empty"}
	ErrInvalidEmail                    = utils.RespError{Code: http.StatusBadRequest, Message: "Invalid email address"}
	ErrDataIsEmpty                     = utils.RespError{Code: http.StatusBadRequest, Message: "List of friend request is empty"}
	ErrRequesterEmailAndAddresseeEmail = utils.RespError{Code: http.StatusBadRequest, Message: "Requester email and target email must not be the same"}
)
