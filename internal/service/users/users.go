package users

import (
	"context"
	"database/sql"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository"
	models "github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository/orm/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository/user"

	"github.com/volatiletech/null/v8"
)

// UsersService type contain users repository
type UsersService struct {
	userRepository repository.UsersRepo
}

// UsersServ define function of users
type UsersServ interface {
	CreateUser(ctx context.Context, input CreateUserInput) (UsersResponse, error)
	GetListUser(ctx context.Context) (UsersEmailResponse, error)
}

// CreateUserInput param using create users
type CreateUserInput struct {
	Email    string
	Phone    string
	IsActive bool
}

// UsersResponse response handler create users
type UsersResponse struct {
	ID       int
	Email    string
	Phone    string
	IsActive bool
}

// UsersEmailResponse type list email of users
type UsersEmailResponse struct {
	Email []string
}

// NewUserService create new users service
func NewUserService(db *sql.DB) UsersServ {
	return UsersService{
		userRepository: user.NewUserRepository(db),
	}
}

// CreateUser creates new users
func (serv UsersService) CreateUser(ctx context.Context, input CreateUserInput) (UsersResponse, error) {
	user, err := serv.userRepository.CreateUser(ctx, models.User{
		Email:    input.Email,
		Phone:    null.StringFrom(input.Phone),
		IsActive: null.BoolFrom(input.IsActive),
	})
	if err != nil {
		return UsersResponse{}, err
	}

	return UsersResponse{
		ID:       user.ID,
		Email:    user.Email,
		Phone:    user.Phone.String,
		IsActive: user.IsActive.Bool,
	}, nil
}

// GetListUser return list of users
func (serv UsersService) GetListUser(ctx context.Context) (UsersEmailResponse, error) {
	var userResult UsersEmailResponse
	user, err := serv.userRepository.GetAllUser(ctx)
	if err != nil {
		return UsersEmailResponse{}, err
	}

	userResult.Email = user

	return userResult, nil
}
