package rest

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"strings"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
	userServ "github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/service/users"
)

// UsersRequest request to create new users
type UsersRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
}

// CreateUser end point to create new users
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	userReq := UsersRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	input, err := validateUserInput(userReq)
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	result, err := h.userService.CreateUser(r.Context(), input)
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, result)
}

// GetListUser end point to get list users
func (h Handler) GetListUser(w http.ResponseWriter, r *http.Request) {
	result, err := h.userService.GetListUser(r.Context())
	if err != nil {
		utils.JsonResponseError(w, err)
		return
	}

	utils.ResponseJson(w, http.StatusOK, result)
}

// validateUserInput validate users request and convert to service input
func validateUserInput(user UsersRequest) (userServ.CreateUserInput, error) {
	email := strings.TrimSpace(user.Email)
	if email == "" {
		return userServ.CreateUserInput{}, ErrEmailCannotBeBlank
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return userServ.CreateUserInput{}, ErrInvalidEmail
	}

	return userServ.CreateUserInput{
		Email:    email,
		Phone:    user.Phone,
		IsActive: user.IsActive,
	}, nil
}
