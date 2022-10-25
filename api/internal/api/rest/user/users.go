package user

import (
	"database/sql"
	"encoding/json"
	"github.com/quangpham789/golang-assessment/api/internal/api/rest/errors"
	userService "github.com/quangpham789/golang-assessment/api/internal/controller/user"
	"github.com/quangpham789/golang-assessment/api/internal/pkg/utils"
	"net/http"
	"net/mail"
	"strings"
)

type UserHandler struct {
	userService userService.UserServ
}

func NewUserHandler(db *sql.DB) UserHandler {
	return UserHandler{
		userService: userService.NewUserService(db),
	}
}

type UserRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
}

func (user UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Convert body request to struct of Handler
	userReq := UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Validate body user request
	input, err := validateUserInput(userReq)
	if err != nil {
		errors.JsonResponseError(w, err)
		return
	}

	// Call service
	result, err := user.userService.CreateUser(r.Context(), input)
	if err != nil {
		errors.JsonResponseError(w, err)
		return
	}
	utils.JsonResponse(w, http.StatusCreated, result)
}

func validateUserInput(user UserRequest) (userService.CreateUserInput, error) {
	email := strings.TrimSpace(user.Email)
	if email == "" {
		return userService.CreateUserInput{}, errors.ErrEmailCannotBeBlank
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return userService.CreateUserInput{}, errors.ErrInvalidEmail
	}

	return userService.CreateUserInput{
		Email:    email,
		Phone:    user.Phone,
		IsActive: user.IsActive,
	}, nil
}
