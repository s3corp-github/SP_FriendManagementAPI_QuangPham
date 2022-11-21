package users

import (
	"context"
	"database/sql"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository/user"
)

// IService define function of users
type IService interface {
	CreateUser(ctx context.Context, input UserEmail) (UserResponse, error)
	GetUsers(ctx context.Context) ([]UserEmail, error)
}

// UserService type contain users repository
type UserService struct {
	userRepository repository.UsersRepo
}

// NewUserService create new users service
func NewUserService(db *sql.DB) IService {
	return UserService{
		userRepository: user.NewUserRepository(db),
	}
}
