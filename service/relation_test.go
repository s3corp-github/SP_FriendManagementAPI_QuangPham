package service

import (
	"context"
	"github.com/quangpham789/golang-assessment/repository/relation"
	"github.com/quangpham789/golang-assessment/repository/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_CreateFriendRelation(t *testing.T) {
	tcs := map[string]struct {
		input     CreateRelationsInput
		expResult CreateRelationsResponse
		expErr    error
	}{
		"success": {
			input: CreateRelationsInput{
				RequesterEmail: "common@gmail.com",
				AddresseeEmail: "john@gmail.com",
			},
			expResult: CreateRelationsResponse{
				Success: true,
			},
		},
	}

	//tcMockUserRepo := map[string]struct {
	//	result models.User
	//	err    error
	//}{
	//	"success": {
	//		result: models.User{
	//			ID:       15,
	//			Email:    "nhutquang23@gmail.com",
	//			Phone:    null.StringFrom("123456"),
	//			IsActive: null.BoolFrom(true),
	//		},
	//	},
	//}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUserRepo := new(user.MockUserRepo)
			mockUserRepo.On("GetUserByEmail", mock.Anything, mock.Anything).
				Return(tc.expResult, tc.expErr)

			mockRelationRepo := new(relation.MockRelationRepo)
			mockRelationRepo.On("CreateRelation", mock.Anything, mock.Anything).
				Return(tc.expResult, tc.expErr)

			relationService := RelationsService{mockRelationRepo, mockUserRepo}
			res, err := relationService.CreateFriendRelation(ctx, tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}
