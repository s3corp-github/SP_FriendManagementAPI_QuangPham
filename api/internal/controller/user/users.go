package user

import (
	"context"
	"database/sql"
	"github.com/quangpham789/golang-assessment/api/internal/repository"
	models "github.com/quangpham789/golang-assessment/api/internal/repository/orm/models"
	"github.com/quangpham789/golang-assessment/api/internal/repository/user"
	"github.com/volatiletech/null/v8"
)

type UserService struct {
	userRepository repository.UserRepo
}

type UserServ interface {
	CreateUser(ctx context.Context, input CreateUserInput) (UserResponse, error)
}

type CreateUserInput struct {
	Email    string
	Phone    string
	IsActive bool
}

type UserResponse struct {
	ID       int
	Email    string
	Phone    string
	IsActive bool
}

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
