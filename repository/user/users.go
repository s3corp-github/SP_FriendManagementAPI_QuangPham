package user

import (
	"context"
	"database/sql"
	models "github.com/quangpham789/golang-assessment/models"
	"github.com/quangpham789/golang-assessment/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewUserRepository(db *sql.DB) repository.UserRepo {
	return userRepository{
		connection: db,
	}
}

type userRepository struct {
	connection *sql.DB
}

func (repo userRepository) GetUserByID(ctx context.Context, id int) (models.User, error) {
	var userResult models.User
	err := models.Users(models.UserWhere.ID.EQ(id)).Bind(ctx, repo.connection, &userResult)
	if err != nil {
		return models.User{}, err
	}
	return userResult, nil
}

// CreateUser creates user
func (repo userRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	if err := user.Insert(ctx, repo.connection, boil.Infer()); err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetUserByEmail get user by email
func (repo userRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var userResult models.User
	if err := models.Users(models.UserWhere.Email.EQ(email)).Bind(ctx, repo.connection, &userResult); err != nil {
		return models.User{}, err
	}
	return userResult, nil
}
