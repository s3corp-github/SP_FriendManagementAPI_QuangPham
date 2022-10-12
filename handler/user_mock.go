package handler

import (
	"context"
	"github.com/quangpham789/golang-assessment/service"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, user service.CreateUserInput) (service.UserResponse, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(service.UserResponse), args.Error(1)
}
