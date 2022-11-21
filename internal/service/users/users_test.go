package users

import (
	"context"
	"testing"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestService_CreateUser(t *testing.T) {
	tcs := map[string]struct {
		input          CreateUserInput
		expResult      UserResponse
		expResultCheck bool
		expErr         error
	}{
		"success": {
			input: CreateUserInput{
				Email: "john1@gmail.com",
			},
			expResultCheck: false,
			expResult: UserResponse{
				ID:    15,
				Email: "john1@gmail.com",
			},
		},
	}

	tcMockUserRepo := map[string]struct {
		result         models.User
		expResultCheck bool
		err            error
	}{
		"success": {
			expResultCheck: false,
			result: models.User{
				ID:    15,
				Email: "john1@gmail.com",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(user.UsersRepoMock)
			mockRepo.On("CheckEmailIsExist", mock.Anything, mock.Anything).
				Return(tcMockUserRepo[desc].expResultCheck, tcMockUserRepo[desc].err)
			mockRepo.On("CreateUser", mock.Anything, mock.Anything).
				Return(tcMockUserRepo[desc].result, tcMockUserRepo[desc].err)

			userService := UserService{mockRepo}
			res, err := userService.CreateUser(ctx, tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}

func TestService_GetAllUser(t *testing.T) {
	tcs := map[string]struct {
		expResult     []UserEmailResponse
		expResultRepo models.UserSlice
		expErr        error
	}{
		"success": {
			expResult: []UserEmailResponse{
				{
					Email: "andy@example.com",
					Name:  "andy",
				},
			},
			expResultRepo: models.UserSlice{
				{
					Email: "andy@example.com",
					Name:  "andy",
				},
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(user.UsersRepoMock)
			mockRepo.On("GetUsers", mock.Anything, mock.Anything).
				Return(tc.expResultRepo, tc.expErr)

			userService := UserService{mockRepo}
			res, err := userService.GetUsers(ctx)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}
