package user

import (
	"context"
	models "github.com/quangpham789/golang-assessment/api/internal/repository/orm/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m MockUserRepo) GetUserByID(ctx context.Context, id int) (models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m MockUserRepo) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(models.User), args.Error(1)
}

func (m MockUserRepo) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m MockUserRepo) GetListEmailByIDs(ctx context.Context, ids []int) ([]string, error) {
	//args := m.Called(ctx, ids)
	//return args.Get(0).(models.UserSlice), args.Error(1)
	panic("To implement")
}
