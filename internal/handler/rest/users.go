package rest

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"strings"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/service/users"
)

// UsersRequest request to create new users
type UsersRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
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

// GetUsers end point to get list users
func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := h.userService.GetUsers(r.Context())
	if err != nil {
		utils.JsonResponseError(w, err)
	}

	utils.ResponseJson(w, http.StatusOK, result)
}

// validateUserInput validate users request and convert to service input
func validateUserInput(input UsersRequest) (users.CreateUserInput, error) {
	email := strings.TrimSpace(input.Email)
	if _, err := mail.ParseAddress(email); err != nil {
		return users.CreateUserInput{}, ErrInvalidEmail
	}

	return users.CreateUserInput{
		Name:  input.Name,
		Email: email,
	}, nil
}
