package user

import (
	"context"
	"testing"

	models "github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/repository/orm/models"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
)

func TestService_CreateUser(t *testing.T) {
	tcs := map[string]struct {
		input     CreateUserInput
		expResult UsersResponse
		expErr    error
	}{
		"success": {
			input: CreateUserInput{
				Email:    "nhutquang23@gmail.com",
				Phone:    "123456",
				IsActive: true,
			},
			expResult: UsersResponse{
				ID:       15,
				Email:    "nhutquang23@gmail.com",
				Phone:    "123456",
				IsActive: true,
			},
		},
	}

	tcMockUserRepo := map[string]struct {
		result models.User
		err    error
	}{
		"success": {
			result: models.User{
				ID:       15,
				Email:    "nhutquang23@gmail.com",
				Phone:    null.StringFrom("123456"),
				IsActive: null.BoolFrom(true),
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(mocks.UserRepo)
			mockRepo.On("CreateUser", mock.Anything, mock.Anything).
				Return(tcMockUserRepo[desc].result, tcMockUserRepo[desc].err)

			userService := UsersService{mockRepo}
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
		expResult     UsersEmailResponse
		expResultRepo []string
		expErr        error
	}{
		"success": {
			expResult: UsersEmailResponse{
				Email: []string{
					"andy@example.com",
					"john@example.com",
					"common@example.com",
					"lisa@example.com",
				},
			},
			expResultRepo: []string{
				"andy@example.com",
				"john@example.com",
				"common@example.com",
				"lisa@example.com",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(mocks.UserRepo)
			mockRepo.On("GetAllUser", mock.Anything, mock.Anything).
				Return(tc.expResultRepo, tc.expErr)

			userService := UsersService{mockRepo}
			res, err := userService.GetListUser(ctx)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}
