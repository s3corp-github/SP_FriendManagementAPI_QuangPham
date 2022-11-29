package friends

import (
	"context"
	"testing"

	"github.com/friendsofgo/errors"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository/friend"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestService_CreateFriendRelation(t *testing.T) {
	tcs := map[string]struct {
		input                   CreateRelationsInput
		userRepoExpResultMock1  models.User
		userRepoExpResultMock2  models.User
		createRelationInputMock models.UserFriend
		checkFriendMock         bool
		checkBlockMock          bool
		expErr                  error
		expErrEmail             error
	}{
		"success": {
			input: CreateRelationsInput{
				RequesterEmail: "john@example.com",
				TargetEmail:    "andy@example.com",
			},
			checkFriendMock: false,
			checkBlockMock:  false,
			userRepoExpResultMock1: models.User{
				ID:    1,
				Email: "john@example.com",
			},
			userRepoExpResultMock2: models.User{
				ID:    2,
				Email: "andy@example.com",
			},
			createRelationInputMock: models.UserFriend{
				RequesterID:  1,
				TargetID:     2,
				RelationType: utils.FriendRelation,
			},
			expErr: nil,
		},
		"case error create friends because friends is existed ": {
			input: CreateRelationsInput{
				RequesterEmail: "john@example.com",
				TargetEmail:    "andy@example.com",
			},
			checkFriendMock: true,
			userRepoExpResultMock1: models.User{
				ID:    1,
				Email: "john@example.com",
			},
			userRepoExpResultMock2: models.User{
				ID:    2,
				Email: "andy@example.com",
			},
			createRelationInputMock: models.UserFriend{
				RequesterID:  1,
				TargetID:     2,
				RelationType: utils.FriendRelation,
			},
			expErrEmail: nil,
			expErr:      ErrRelationIsExists,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUserRepo := new(user.UsersRepoMock)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.RequesterEmail).
				Return(tc.userRepoExpResultMock1, tc.expErrEmail)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.TargetEmail).
				Return(tc.userRepoExpResultMock2, tc.expErrEmail)

			mockFriendsRepo := new(friend.FriendsRepoMock)
			mockFriendsRepo.On("CheckFriendRelationExist", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.checkFriendMock, tc.expErr)
			mockFriendsRepo.On("CheckBlockRelationExist", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.checkBlockMock, tc.expErr)
			mockFriendsRepo.On("CreateUserFriend", mock.Anything, tc.createRelationInputMock).
				Return(tc.expErr)

			friendsService := FriendService{userRepo: mockUserRepo, friendRepo: mockFriendsRepo}
			err := friendsService.CreateFriend(ctx, tc.input)

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestService_GetAllFriendsOfUser(t *testing.T) {
	tcs := map[string]struct {
		input                GetAllFriendsInput
		expUserByEmailMock   models.User
		expRelationIDsOfUser []int
		expListEmailByIDs    []string
		expResult            []string
		expErr               error
	}{
		"success": {
			input: GetAllFriendsInput{
				Email: "john@example.com",
			},
			expUserByEmailMock: models.User{
				ID:    1,
				Email: "john@example.com",
			},
			expRelationIDsOfUser: []int{1},
			expListEmailByIDs:    []string{"andy@example.com"},
			expResult:            []string{"andy@example.com"},
		},
		"case email is empty": {
			input: GetAllFriendsInput{
				Email: "",
			},
			expUserByEmailMock:   models.User{},
			expRelationIDsOfUser: []int{},
			expListEmailByIDs:    []string{},
			expResult:            []string{},
			expErr:               errors.New("Email cannot be empty"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUsersRepo := new(user.UsersRepoMock)
			mockUsersRepo.On("GetUserByEmail", mock.Anything, tc.input.Email).
				Return(tc.expUserByEmailMock, tc.expErr)

			mockFriendsRepo := new(friend.FriendsRepoMock)
			mockFriendsRepo.On("GetRelationIDs", mock.Anything, tc.expUserByEmailMock.ID, 1).
				Return(tc.expRelationIDsOfUser, tc.expErr)
			mockUsersRepo.On("GetEmailsByIDs", mock.Anything, mock.Anything).
				Return(tc.expListEmailByIDs, tc.expErr)

			relationService := FriendService{friendRepo: mockFriendsRepo, userRepo: mockUsersRepo}
			res, err := relationService.GetFriends(ctx, tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}
func TestService_GetCommonFriendList(t *testing.T) {
	tcs := map[string]struct {
		input               CommonFriendsInput
		expUserByEmailMock1 models.User
		expUserByEmailMock2 models.User
		expIdsFirstEmail    []int
		expIdsSecondEmail   []int
		expIntersectionIDs  []int
		expListEmail        []string
		expResult           []string
		expErr              error
	}{
		"success": {
			input: CommonFriendsInput{
				RequesterEmail: "john@example.com",
				TargetEmail:    "andy@example.com",
			},
			expUserByEmailMock1: models.User{
				ID:    1,
				Email: "john@example.com",
			},
			expUserByEmailMock2: models.User{
				ID:    2,
				Email: "andy@example.com",
			},
			expIdsFirstEmail:   []int{3},
			expIdsSecondEmail:  []int{3},
			expIntersectionIDs: []int{3},
			expListEmail:       []string{"common@example.com"},
			expResult:          []string{"common@example.com"},
		},
		"case first email is empty": {
			input: CommonFriendsInput{
				TargetEmail: "andy@example.com",
			},
			expUserByEmailMock1: models.User{},
			expUserByEmailMock2: models.User{
				ID:    2,
				Email: "andy@example.com",
			},
			expIdsFirstEmail:   []int{},
			expIdsSecondEmail:  []int{},
			expIntersectionIDs: []int{},
			expListEmail:       []string{},
			expResult:          []string{},
			expErr:             errors.New("Email cannot be empty"),
		},
		"case second email is empty": {
			input: CommonFriendsInput{
				RequesterEmail: "john@example.com",
			},
			expUserByEmailMock1: models.User{
				ID:    1,
				Email: "john@example.com",
			},
			expUserByEmailMock2: models.User{},
			expIdsFirstEmail:    []int{},
			expIdsSecondEmail:   []int{},
			expIntersectionIDs:  []int{},
			expListEmail:        []string{},
			expResult:           []string{},
			expErr:              errors.New("Email cannot be empty"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUsersRepo := new(user.UsersRepoMock)
			mockUsersRepo.On("GetUserByEmail", mock.Anything, tc.input.RequesterEmail).
				Return(tc.expUserByEmailMock1, tc.expErr)
			mockUsersRepo.On("GetUserByEmail", mock.Anything, tc.input.TargetEmail).
				Return(tc.expUserByEmailMock2, tc.expErr)

			mockFriendRepo := new(friend.FriendsRepoMock)
			mockFriendRepo.On("GetRelationIDs", mock.Anything, tc.expUserByEmailMock1.ID, 1).
				Return(tc.expIdsFirstEmail, tc.expErr)
			mockFriendRepo.On("GetRelationIDs", mock.Anything, tc.expUserByEmailMock2.ID, 1).
				Return(tc.expIdsSecondEmail, tc.expErr)

			mockUsersRepo.On("GetEmailsByIDs", mock.Anything, mock.Anything).
				Return(tc.expListEmail, tc.expErr)

			relationService := FriendService{friendRepo: mockFriendRepo, userRepo: mockUsersRepo}
			res, err := relationService.GetCommonFriends(ctx, tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}
func TestService_CreateSubscriptionRelation(t *testing.T) {
	tcs := map[string]struct {
		input                   CreateRelationsInput
		userRepoExpResultMock1  models.User
		userRepoExpResultMock2  models.User
		createRelationInputMock models.UserFriend
		isExistMock             bool
		expErr                  error
		expErrEmail             error
	}{
		"success": {
			input: CreateRelationsInput{
				RequesterEmail: "john@example.com",
				TargetEmail:    "andy@example.com",
			},
			isExistMock: false,
			userRepoExpResultMock1: models.User{
				ID:    1,
				Email: "john@example.com",
			},
			userRepoExpResultMock2: models.User{
				ID:    2,
				Email: "andy@example.com",
			},
			createRelationInputMock: models.UserFriend{
				RequesterID:  1,
				TargetID:     2,
				RelationType: utils.Subscribe,
			},
			expErr: nil,
		},
		"case requester users not found": {
			input: CreateRelationsInput{
				RequesterEmail: "",
				TargetEmail:    "lisa@example.com",
			},
			isExistMock:            false,
			userRepoExpResultMock1: models.User{},
			userRepoExpResultMock2: models.User{
				ID:    4,
				Email: "lisa@example.com",
			},
			createRelationInputMock: models.UserFriend{
				RequesterID:  0,
				TargetID:     4,
				RelationType: utils.Subscribe,
			},
			expErr: errors.New("Email cannot be empty"),
		},
		"case target users not found": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "",
			},
			isExistMock: false,
			userRepoExpResultMock1: models.User{
				ID:    3,
				Email: "common@example.com",
			},
			userRepoExpResultMock2: models.User{},
			createRelationInputMock: models.UserFriend{
				RequesterID:  3,
				RelationType: utils.FriendRelation,
			},
			expErr: errors.New("Email cannot be empty"),
		},
		"case error create friends because friend friends is existed ": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "lisa@example.com",
			},
			isExistMock: true,
			userRepoExpResultMock1: models.User{
				ID:    3,
				Email: "common@example.com",
			},
			userRepoExpResultMock2: models.User{
				ID:    4,
				Email: "lisa@example.com",
			},
			createRelationInputMock: models.UserFriend{
				RequesterID:  3,
				TargetID:     4,
				RelationType: utils.Subscribe,
			},
			expErr: errors.New("Email cannot be empty"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUserRepo := new(user.UsersRepoMock)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.RequesterEmail).
				Return(tc.userRepoExpResultMock1, tc.expErrEmail)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.TargetEmail).
				Return(tc.userRepoExpResultMock2, tc.expErrEmail)

			mockRelationRepo := new(friend.FriendsRepoMock)
			mockRelationRepo.On("CheckSubscriptionRelationExist", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.isExistMock, tc.expErr)
			mockRelationRepo.On("CreateUserFriend", mock.Anything, tc.createRelationInputMock).
				Return(tc.expErr)

			relationService := FriendService{friendRepo: mockRelationRepo, userRepo: mockUserRepo}

			err := relationService.CreateSubscription(ctx, tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestService_CreateBlockRelation(t *testing.T) {
	tcs := map[string]struct {
		input                   CreateRelationsInput
		userRepoExpResultMock1  models.User
		userRepoExpResultMock2  models.User
		createRelationInputMock models.UserFriend
		isExistMock             bool
		expErrEmail             error
		expErr                  error
	}{
		"success": {
			input: CreateRelationsInput{
				RequesterEmail: "john@example.com",
				TargetEmail:    "andy@example.com",
			},
			isExistMock: false,
			userRepoExpResultMock1: models.User{
				ID:    1,
				Email: "john@example.com",
			},
			userRepoExpResultMock2: models.User{
				ID:    2,
				Email: "andy@example.com",
			},
			createRelationInputMock: models.UserFriend{
				RequesterID:  1,
				TargetID:     2,
				RelationType: utils.Blocked,
			},
		},
		"case requester users not found": {
			input: CreateRelationsInput{
				RequesterEmail: "",
				TargetEmail:    "andy@example.com",
			},
			isExistMock:            false,
			userRepoExpResultMock1: models.User{},
			userRepoExpResultMock2: models.User{
				ID:    2,
				Email: "lisa@example.com",
			},
			createRelationInputMock: models.UserFriend{
				RequesterID:  0,
				TargetID:     2,
				RelationType: utils.Blocked,
			},
			expErr: errors.New("Email cannot be empty"),
		},
		"case target users not found": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "",
			},
			isExistMock: false,
			userRepoExpResultMock1: models.User{
				ID:    1,
				Email: "john@example.com",
			},
			userRepoExpResultMock2: models.User{},
			createRelationInputMock: models.UserFriend{
				RequesterID:  1,
				RelationType: utils.Blocked,
			},
			expErr: errors.New("Email cannot be empty"),
		},
		"case error create friends because friend friends is existed ": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "lisa@example.com",
			},
			isExistMock: true,
			userRepoExpResultMock1: models.User{
				ID:    1,
				Email: "john@example.com",
			},
			userRepoExpResultMock2: models.User{
				ID:    2,
				Email: "andy@example.com",
			},
			createRelationInputMock: models.UserFriend{
				RequesterID:  1,
				TargetID:     2,
				RelationType: utils.Blocked,
			},
			expErr: errors.New("Email cannot be empty"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUsersRepo := new(user.UsersRepoMock)
			mockUsersRepo.On("GetUserByEmail", mock.Anything, tc.input.RequesterEmail).
				Return(tc.userRepoExpResultMock1, tc.expErrEmail)
			mockUsersRepo.On("GetUserByEmail", mock.Anything, tc.input.TargetEmail).
				Return(tc.userRepoExpResultMock2, tc.expErrEmail)

			mockFriendRepo := new(friend.FriendsRepoMock)
			mockFriendRepo.On("CheckBlockRelationExist", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.isExistMock, tc.expErr)
			mockFriendRepo.On("CreateUserFriend", mock.Anything, tc.createRelationInputMock).
				Return(tc.expErr)
			mockFriendRepo.On("DeleteRelation", mock.Anything, tc.userRepoExpResultMock1.ID, tc.userRepoExpResultMock2.ID, utils.Subscribe).
				Return(nil, tc.expErr)

			relationService := FriendService{friendRepo: mockFriendRepo, userRepo: mockUsersRepo}

			err := relationService.CreateBlock(ctx, tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestService_GetEmailReceive(t *testing.T) {
	tcs := map[string]struct {
		input                   EmailReceiveInput
		expUserByEmailMock      models.User
		expIdsEmail             []int
		expRequesterIdsRelation []int
		expEmailFromText        []string
		expUserIDByEmail        []int
		expUserIDBlock          []int
		expReceiveIDs           []int
		expUniqueSlice          []int

		expListEmail       []string
		expListEmailResult []string
		expResult          []string
		expErr             error
	}{
		"success": {
			input: EmailReceiveInput{
				Sender: "andy@example.com",
				Text:   "Hello World! lisa@example.com",
			},
			expUserByEmailMock: models.User{
				ID:    1,
				Email: "andy@example.com",
			},
			expIdsEmail:             []int{3},
			expRequesterIdsRelation: []int{2},
			expListEmail:            []string{"lisa@example.com"},
			expUserIDByEmail:        []int{4},
			expUserIDBlock:          []int{2},
			expReceiveIDs:           []int{3, 4},
			expListEmailResult:      []string{"common@example.com", "lisa@example.com"},
			expResult:               []string{"common@example.com", "lisa@example.com"},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUsersRepo := new(user.UsersRepoMock)
			mockUsersRepo.On("GetUserByEmail", mock.Anything, tc.input.Sender).
				Return(tc.expUserByEmailMock, tc.expErr)

			mockFriendsRepo := new(friend.FriendsRepoMock)
			mockFriendsRepo.On("GetRelationIDs", mock.Anything, tc.expUserByEmailMock.ID, 1).
				Return(tc.expIdsEmail, tc.expErr)
			mockFriendsRepo.On("GetRequesterIDFriends", mock.Anything, mock.Anything, mock.Anything).
				Return(tc.expRequesterIdsRelation, tc.expErr)

			mockUsersRepo.On("GetUserIDsByEmail", mock.Anything, tc.expListEmail).
				Return(tc.expUserIDByEmail, tc.expErr)

			mockUsersRepo.On("GetEmailsByIDs", mock.Anything, tc.expReceiveIDs).
				Return(tc.expListEmailResult, tc.expErr)

			relationService := FriendService{friendRepo: mockFriendsRepo, userRepo: mockUsersRepo}
			res, err := relationService.GetEmailReceive(ctx, tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}
