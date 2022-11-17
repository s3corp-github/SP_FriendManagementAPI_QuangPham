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
		input     CreateUserInput
		expResult UserResponse
		expErr    error
	}{
		"success": {
			input: CreateUserInput{
				Email: "nhutquang23@gmail.com",
			},
			expResult: UserResponse{
				ID:    15,
				Email: "nhutquang23@gmail.com",
			},
		},
	}

	tcMockUserRepo := map[string]struct {
		result models.User
		err    error
	}{
		"success": {
			result: models.User{
				ID:    15,
				Email: "nhutquang23@gmail.com",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(user.UsersRepoMock)
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
		expResultRepo []string
		expErr        error
	}{
		"success": {
			expResult: []UserEmailResponse{
				{
					Email: "andy@example.com",
					Name:  "andy",
				},
			},
			expResultRepo: []string{
				"andy@example.com",
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
