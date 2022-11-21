// Code generated by mockery v2.15.0. DO NOT EDIT.

package friend

import (
	context "context"

	models "github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// FriendsRepoMock is an autogenerated mock type for the FriendsRepoMock type
type FriendsRepoMock struct {
	mock.Mock
}

// CreateUserFriend provides a mock function with given fields: ctx, input
func (_m *FriendsRepoMock) CreateUserFriend(ctx context.Context, input models.UserFriend) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UserFriend) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRelation provides a mock function with given fields: ctx, RequesterID, targetID, relationType
func (_m *FriendsRepoMock) DeleteRelation(ctx context.Context, requesterID int, targetID int, relationType int) error {
	ret := _m.Called(ctx, requesterID, targetID, relationType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) error); ok {
		r0 = rf(ctx, requesterID, targetID, relationType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRelationIDs provides a mock function with given fields: ctx, RequesterID, relationType
func (_m *FriendsRepoMock) GetRelationIDs(ctx context.Context, requesterID int, relationType int) ([]int, error) {
	ret := _m.Called(ctx, requesterID, relationType)

	var r0 []int
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []int); ok {
		r0 = rf(ctx, requesterID, relationType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, requesterID, relationType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRequesterIDFriends provides a mock function with given fields: ctx, RequesterID, relationType
func (_m *FriendsRepoMock) GetRequesterIDFriends(ctx context.Context, requesterID int, relationType int) ([]int, error) {
	ret := _m.Called(ctx, requesterID, relationType)

	var r0 []int
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []int); ok {
		r0 = rf(ctx, requesterID, relationType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, requesterID, relationType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsRelationExist provides a mock function with given fields: ctx, RequesterID, targetID, relationType
func (_m *FriendsRepoMock) IsRelationExist(ctx context.Context, requesterID int, targetID int, relationType int) (bool, error) {
	ret := _m.Called(ctx, requesterID, targetID, relationType)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) bool); ok {
		r0 = rf(ctx, requesterID, targetID, relationType)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, int) error); ok {
		r1 = rf(ctx, requesterID, targetID, relationType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFriendsRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewFriendsRepo creates a new instance of FriendsRepoMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFriendsRepo(t mockConstructorTestingTNewFriendsRepo) *FriendsRepoMock {
	mock := &FriendsRepoMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}