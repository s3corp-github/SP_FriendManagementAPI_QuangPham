package user

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/pkg/utils"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/repository"
	models "github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/repository/orm/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func NewUserRepository(db *sql.DB) repository.UserRepo {
	return userRepository{
		connection: db,
	}
}

type userRepository struct {
	connection *sql.DB
}

func (repo userRepository) GetListEmailByIDs(ctx context.Context, ids []int) ([]string, error) {
	var emails []string
	convertedIDs := make([]interface{}, len(ids))
	for index, num := range ids {
		convertedIDs[index] = num
	}
	userResult, err := models.Users(qm.WhereIn(models.UserColumns.ID+" IN ?", convertedIDs...)).All(ctx, repo.connection)
	if err != nil {
		return nil, err
	}
	for _, i := range userResult {
		emails = append(emails, i.Email)
	}
	return emails, nil

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
	if email == "" {
		return models.User{}, errors.New(utils.ErrMessageCanNotBeBlank)
	}
	var userResult models.User
	if err := models.Users(models.UserWhere.Email.EQ(email)).Bind(ctx, repo.connection, &userResult); err != nil {
		return models.User{}, err
	}
	return userResult, nil
}

func (repo userRepository) GetUserIDsByEmail(ctx context.Context, emails []string) ([]int, error) {

	if len(emails) == 0 {
		return nil, errors.New(utils.ErrMessageSlideCannotEmpty)
	}

	args := make([]interface{}, len(emails))
	for i, email := range emails {
		args[i] = email
	}

	slice, err := models.Users(
		qm.Select(models.UserColumns.ID, models.UserColumns.Email),
		qm.WhereIn(models.UserColumns.Email+" in ?", args...),
	).All(ctx, repo.connection)

	if err != nil {
		return nil, err
	}

	result := make([]int, len(slice))
	for idx, u := range slice {
		result[idx] = u.ID
	}

	return result, nil
}

func (repo userRepository) GetAllUser(ctx context.Context) ([]string, error) {

	var emailUser []string
	userResult, err := models.Users(
		qm.Select(models.UserColumns.Email),
	).All(ctx, repo.connection)

	if err != nil {
		return nil, err
	}
	for _, i := range userResult {
		emailUser = append(emailUser, i.Email)
	}
	return emailUser, nil
}
