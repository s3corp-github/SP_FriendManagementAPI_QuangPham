package service

import (
	"context"
	models "github.com/quangpham789/golang-assessment/models"
	"github.com/quangpham789/golang-assessment/repository/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"testing"
)

func TestService_CreateUser(t *testing.T) {
	tcs := map[string]struct {
		input     CreateUserInput
		expResult UserResponse
		expErr    error
	}{
		"success": {
			input: CreateUserInput{
				Email:    "nhutquang23@gmail.com",
				Phone:    "123456",
				IsActive: true,
			},
			expResult: UserResponse{
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
			mockRepo := new(user.MockUserRepo)
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
