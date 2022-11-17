package rest

import (
	"net/http"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
)

var (
	ErrInvalidName             = utils.RespError{Code: http.StatusBadRequest, Message: "Name cannot be empty"}
	ErrInvalidEmail            = utils.RespError{Code: http.StatusBadRequest, Message: "Invalid email address"}
	ErrInvalidBodyRequest      = utils.RespError{Code: http.StatusBadRequest, Message: "Invalid body request"}
	ErrRequesterAndTargetEmail = utils.RespError{Code: http.StatusBadRequest, Message: "Requester email and target email must not be the same"}
)
