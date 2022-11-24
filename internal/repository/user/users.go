package user

import (
	"context"
	"database/sql"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func NewUserRepository(db *sql.DB) repository.UsersRepo {
	return userRepository{
		db: db,
	}
}

type userRepository struct {
	db *sql.DB
}

// GetEmailsByIDs function get user emails by IDs
func (repo userRepository) GetEmailsByIDs(ctx context.Context, ids []int) ([]string, error) {
	var emails []string
	convertedIDs := make([]interface{}, len(ids))
	for index, num := range ids {
		convertedIDs[index] = num
	}

	userResult, err := models.Users(qm.WhereIn(models.UserColumns.ID+" IN ?", convertedIDs...)).All(ctx, repo.db)
	if err != nil {
		return nil, err
	}

	for _, i := range userResult {
		emails = append(emails, i.Email)
	}

	return emails, nil
}

// GetUserByID function get user by ID
func (repo userRepository) GetUserByID(ctx context.Context, id int) (models.User, error) {
	var userResult models.User

	if err := models.Users(models.UserWhere.ID.EQ(id)).Bind(ctx, repo.db, &userResult); err != nil {
		return models.User{}, err
	}

	return userResult, nil
}

// CreateUser creates users
func (repo userRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	if err := user.Insert(ctx, repo.db, boil.Infer()); err != nil {
		return models.User{}, err
	}

	return user, nil
}

// GetUserByEmail get users by email
func (repo userRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var userResult models.User
	if err := models.Users(models.UserWhere.Email.EQ(email)).Bind(ctx, repo.db, &userResult); err != nil {
		return models.User{}, err
	}

	return userResult, nil
}

// GetUserIDsByEmail function get userIDS by emails
func (repo userRepository) GetUserIDsByEmail(ctx context.Context, emails []string) ([]int, error) {
	emailParams := make([]interface{}, len(emails))
	for i, email := range emails {
		emailParams[i] = email
	}

	users, err := models.Users(
		qm.Select(models.UserColumns.ID, models.UserColumns.Email),
		qm.WhereIn(models.UserColumns.Email+" in ?", emailParams...),
	).All(ctx, repo.db)
	if err != nil {
		return nil, err
	}

	result := make([]int, len(users))
	for idx, u := range users {
		result[idx] = u.ID
	}

	return result, nil
}

// GetUsers get list users
func (repo userRepository) GetUsers(ctx context.Context) (models.UserSlice, error) {
	return models.Users(
		qm.Select(models.UserColumns.Name, models.UserColumns.Email),
	).All(ctx, repo.db)
}

// CheckEmailIsExist check if email is existing in db
func (repo userRepository) CheckEmailIsExist(ctx context.Context, email string) (bool, error) {
	return models.Users(
		models.UserWhere.Email.EQ(email),
	).Exists(ctx, repo.db)
}
