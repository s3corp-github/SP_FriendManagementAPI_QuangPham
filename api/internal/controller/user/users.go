package user

import (
	"context"
	"database/sql"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/repository"
	models "github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/repository/orm/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/repository/user"
	"github.com/volatiletech/null/v8"
)

// UserService type contain user repository
type UserService struct {
	userRepository repository.UserRepo
}

// UserServ define function of user
type UserServ interface {
	CreateUser(ctx context.Context, input CreateUserInput) (UserResponse, error)
	GetListUser(ctx context.Context) ([]UserResponse, error)
}

// CreateUserInput param using create user
type CreateUserInput struct {
	Email    string
	Phone    string
	IsActive bool
}

// UserResponse response api create user
type UserResponse struct {
	ID       int
	Email    string
	Phone    string
	IsActive bool
}

// NewUserService create new user service
func NewUserService(db *sql.DB) UserServ {
	return UserService{
		userRepository: user.NewUserRepository(db),
	}
}

// CreateUser creates new user
func (serv UserService) CreateUser(ctx context.Context, input CreateUserInput) (UserResponse, error) {
	user, err := serv.userRepository.CreateUser(ctx, models.User{
		Email:    input.Email,
		Phone:    null.StringFrom(input.Phone),
		IsActive: null.BoolFrom(input.IsActive),
	})
	if err != nil {
		return UserResponse{}, err
	}

	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Phone:    user.Phone.String,
		IsActive: user.IsActive.Bool,
	}, nil
}

// GetListUser return list of user
func (serv UserService) GetListUser(ctx context.Context) ([]UserResponse, error) {
	var userResult []UserResponse
	user, err := serv.userRepository.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range user {
		var item = UserResponse{
			ID:       u.ID,
			Email:    u.Email,
			Phone:    u.Phone.String,
			IsActive: u.IsActive.Bool,
		}
		userResult = append(userResult, item)
	}

	return userResult, nil
}
