package users

import (
	"context"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
)

// CreateUserInput param using create users
type CreateUserInput struct {
	Name  string
	Email string
}

// UserResponse response handler create users
type UserResponse struct {
	ID    int
	Email string
}

// UserEmailResponse type list email of users
type UserEmailResponse struct {
	Email string
	Name  string
}

// CreateUser creates new users
func (serv UserService) CreateUser(ctx context.Context, input CreateUserInput) (UserResponse, error) {
	checkEmail, err := serv.userRepository.CheckEmailIsExist(ctx, input.Email)
	if err != nil {
		return UserResponse{}, err
	}
	if checkEmail {
		return UserResponse{}, ErrEmailExist
	}

	user, err := serv.userRepository.CreateUser(ctx, models.User{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		return UserResponse{}, ErrCreateUser
	}

	return UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

// GetUsers return list of users
func (serv UserService) GetUsers(ctx context.Context) ([]UserEmailResponse, error) {
	var userResult []UserEmailResponse
	users, err := serv.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		user := UserEmailResponse{
			Name:  u.Name,
			Email: u.Email,
		}
		userResult = append(userResult, user)
	}

	return userResult, nil
}
