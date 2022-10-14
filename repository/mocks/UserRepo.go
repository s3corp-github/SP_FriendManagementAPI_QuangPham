// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dbmodels "github.com/quangpham789/golang-assessment/models"
	mock "github.com/stretchr/testify/mock"
)

// UserRepo is an autogenerated mock type for the UserRepo type
type UserRepo struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *UserRepo) CreateUser(ctx context.Context, user dbmodels.User) (dbmodels.User, error) {
	ret := _m.Called(ctx, user)

	var r0 dbmodels.User
	if rf, ok := ret.Get(0).(func(context.Context, dbmodels.User) dbmodels.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(dbmodels.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dbmodels.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListEmailByIDs provides a mock function with given fields: ctx, ids
func (_m *UserRepo) GetListEmailByIDs(ctx context.Context, ids []int) ([]string, error) {
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
func (_m *UserRepo) GetUserByEmail(ctx context.Context, email string) (dbmodels.User, error) {
	ret := _m.Called(ctx, email)

	var r0 dbmodels.User
	if rf, ok := ret.Get(0).(func(context.Context, string) dbmodels.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(dbmodels.User)
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
func (_m *UserRepo) GetUserByID(ctx context.Context, id int) (dbmodels.User, error) {
	ret := _m.Called(ctx, id)

	var r0 dbmodels.User
	if rf, ok := ret.Get(0).(func(context.Context, int) dbmodels.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(dbmodels.User)
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
func (_m *UserRepo) GetUserIDsByEmail(ctx context.Context, email []string) ([]int, error) {
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

type mockConstructorTestingTNewUserRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepo creates a new instance of UserRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepo(t mockConstructorTestingTNewUserRepo) *UserRepo {
	mock := &UserRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
