package user

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/api/rest"
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
		log.Println("CreateUser: error when decode request ", err)
		rest.JsonResponseError(w, err)
		return
	}

	input, err := validateUserInput(userReq)
	if err != nil {
		log.Println("CreateUser: error validate input ", err)
		rest.JsonResponseError(w, err)
		return
	}

	result, err := user.userService.CreateUser(r.Context(), input)
	if err != nil {
		log.Println("CreateUser error", err)
		rest.JsonResponseError(w, err)
		return
	}
	utils.ResponseJson(w, http.StatusCreated, result)
}

// GetListUser end point to get list user
func (user UsersHandler) GetListUser(w http.ResponseWriter, r *http.Request) {
	result, err := user.userService.GetListUser(r.Context())
	if err != nil {
		log.Println("GetListUser error", err)
		rest.JsonResponseError(w, err)
		return
	}
	utils.ResponseJson(w, http.StatusOK, result)
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
