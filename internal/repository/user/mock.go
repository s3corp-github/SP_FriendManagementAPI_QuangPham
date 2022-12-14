// Code generated by mockery v2.14.0. DO NOT EDIT.

package user

import (
	context "context"

	models "github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// UsersRepoMock is an autogenerated mock type for the UsersRepoMock type
type UsersRepoMock struct {
	mock.Mock
}

// CheckEmailIsExist provides a mock function with given fields: ctx, email
func (_m *UsersRepoMock) CheckEmailIsExist(ctx context.Context, email string) (bool, error) {
	ret := _m.Called(ctx, email)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *UsersRepoMock) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	ret := _m.Called(ctx, user)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, models.User) models.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUser provides a mock function with given fields: ctx
func (_m *UsersRepoMock) GetUsers(ctx context.Context) (models.UserSlice, error) {
	ret := _m.Called(ctx)

	var r0 models.UserSlice
	if rf, ok := ret.Get(0).(func(context.Context) models.UserSlice); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.UserSlice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListEmailByIDs provides a mock function with given fields: ctx, ids
func (_m *UsersRepoMock) GetEmailsByIDs(ctx context.Context, ids []int) ([]string, error) {
	ret := _m.Called(ctx, ids)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, []int) []string); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []int) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *UsersRepoMock) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	ret := _m.Called(ctx, email)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: ctx, id
func (_m *UsersRepoMock) GetUserByID(ctx context.Context, id int) (models.User, error) {
	ret := _m.Called(ctx, id)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, int) models.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserIDsByEmail provides a mock function with given fields: ctx, email
func (_m *UsersRepoMock) GetUserIDsByEmail(ctx context.Context, email []string) ([]int, error) {
	ret := _m.Called(ctx, email)

	var r0 []int
	if rf, ok := ret.Get(0).(func(context.Context, []string) []int); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUsersRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsersRepo creates a new instance of UsersRepoMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsersRepo(t mockConstructorTestingTNewUsersRepo) *UsersRepoMock {
	mock := &UsersRepoMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
