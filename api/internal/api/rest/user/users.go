package user

import (
	"database/sql"
	"encoding/json"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/api/rest"
	"net/http"
	"net/mail"
	"strings"

	userService "github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/controller/user"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/pkg/utils"
)

// UsersHandler create users handler contain user service
type UsersHandler struct {
	userService userService.UserServ
}

// NewUserHandler create user handler contain UsersHandler
func NewUserHandler(db *sql.DB) UsersHandler {
	return UsersHandler{
		userService: userService.NewUserService(db),
	}
}

// UsersRequest request to create new user
type UsersRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
}

// CreateUser end point to create new user
func (user UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Convert body request to struct of Handler
	userReq := UsersRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		rest.JsonResponseError(w, err)
		return
	}

	input, err := validateUserInput(userReq)
	if err != nil {
		rest.JsonResponseError(w, err)
		return
	}

	result, err := user.userService.CreateUser(r.Context(), input)
	if err != nil {
		rest.JsonResponseError(w, err)
		return
	}
	utils.ResponseJson(w, http.StatusCreated, result)
}

// validateUserInput function validate user request
func validateUserInput(user UsersRequest) (userService.CreateUserInput, error) {
	email := strings.TrimSpace(user.Email)
	if email == "" {
		return userService.CreateUserInput{}, rest.ErrEmailCannotBeBlank
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return userService.CreateUserInput{}, rest.ErrInvalidEmail
	}

	return userService.CreateUserInput{
		Email:    email,
		Phone:    user.Phone,
		IsActive: user.IsActive,
	}, nil
}
