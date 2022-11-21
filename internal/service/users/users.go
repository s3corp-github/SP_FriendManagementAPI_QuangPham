package users

import (
	"context"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
)

// UserEmail param using create users
type UserEmail struct {
	Name  string
	Email string
}

// UserResponse response handler create users
type UserResponse struct {
	ID    int
	Email string
}

// CreateUser creates new users
func (serv UserService) CreateUser(ctx context.Context, input UserEmail) (UserResponse, error) {
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
func (serv UserService) GetUsers(ctx context.Context) ([]UserEmail, error) {
	var userResult []UserEmail
	users, err := serv.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		user := UserEmail{
			Name:  u.Name,
			Email: u.Email,
		}
		userResult = append(userResult, user)
	}

	return userResult, nil
}
