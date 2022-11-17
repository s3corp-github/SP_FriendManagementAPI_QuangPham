package friends

import (
	"context"
	"testing"

	"github.com/friendsofgo/errors"
	models2 "github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository/friend"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
)

type mockIsRelation struct {
	wantCall  bool
	expResult bool
	expErr    error
}

func TestService_IsRelationExist(t *testing.T) {

	tcs := map[string]struct {
		mockIsRelation mockIsRelation
		requesterID    int
		TargetID       int
		relationType   int
		isExistMock    bool
		expResult      bool
		expErr         error
	}{
		"success": {
			requesterID:  1,
			TargetID:     3,
			relationType: 1,
			isExistMock:  true,
			expResult:    true,
		},
		"case requesterID is 0": {
			TargetID:     3,
			relationType: 1,
			isExistMock:  true,
			expErr:       errors.New("requesterId cannot be null"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockRelationRepo := new(friend.FriendsRepoMock)
			mockRelationRepo.On("IsRelationExist", mock.Anything, tc.requesterID, tc.TargetID, tc.relationType).
				Return(tc.isExistMock, tc.expErr)

			res, err := isRelationExist(ctx, mockRelationRepo, tc.requesterID, tc.TargetID, tc.relationType)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)

			}
		})
	}

}

func TestService_CreateFriendRelation(t *testing.T) {
	tcs := map[string]struct {
		input                   CreateRelationsInput
		userRepoExpResultMock1  models2.User
		userRepoExpResultMock2  models2.User
		createRelationInputMock models2.UserFriend
		repoExpResultMock       bool
		isExistMock             bool
		expResult               bool
		expErr                  error
	}{
		"success": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "lisa@example.com",
			},
			repoExpResultMock: true,
			isExistMock:       false,
			userRepoExpResultMock1: models2.User{
				ID:    3,
				Email: "common@example.com",
			},
			userRepoExpResultMock2: models2.User{
				ID:    4,
				Email: "lisa@example.com",
			},
			createRelationInputMock: models2.UserFriend{
				RequesterID:  3,
				TargetID:     4,
				RelationType: null.IntFrom(utils.FriendRelation),
			},
			expResult: true,
		},
		"case requester users not found": {
			input: CreateRelationsInput{
				RequesterEmail: "",
				TargetEmail:    "lisa@example.com",
			},
			repoExpResultMock:      false,
			isExistMock:            false,
			userRepoExpResultMock1: models2.User{},
			userRepoExpResultMock2: models2.User{
				ID:    4,
				Email: "lisa@example.com",
			},
			createRelationInputMock: models2.UserFriend{
				RequesterID:  0,
				TargetID:     4,
				RelationType: null.IntFrom(utils.FriendRelation),
			},
			expResult: false,
			expErr:    errors.New("Error get requester email from request"),
		},
		"case addressee users not found": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "",
			},
			repoExpResultMock: false,
			isExistMock:       false,
			userRepoExpResultMock1: models2.User{
				ID:    3,
				Email: "common@example.com",
			},
			userRepoExpResultMock2: models2.User{},
			createRelationInputMock: models2.UserFriend{
				RequesterID:  3,
				RelationType: null.IntFrom(utils.FriendRelation),
			},
			expResult: false,
			expErr:    errors.New("Error get requester email from request"),
		},
		"case error create friends because friend friends is existed ": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "lisa@example.com",
			},
			repoExpResultMock: false,
			isExistMock:       true,
			userRepoExpResultMock1: models2.User{
				ID:    3,
				Email: "common@example.com",
			},
			userRepoExpResultMock2: models2.User{
				ID:    4,
				Email: "lisa@example.com",
			},
			createRelationInputMock: models2.UserFriend{
				RequesterID:  3,
				TargetID:     4,
				RelationType: null.IntFrom(utils.FriendRelation),
			},
			expResult: false,
			expErr:    errors.New("Error get requester email from request"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUserRepo := new(user.UsersRepoMock)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.RequesterEmail).
				Return(tc.userRepoExpResultMock1, tc.expErr)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.TargetEmail).
				Return(tc.userRepoExpResultMock2, tc.expErr)

			mockFriendsRepo := new(friend.FriendsRepoMock)
			mockFriendsRepo.On("IsRelationExist", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.isExistMock, tc.expErr)
			mockFriendsRepo.On("CreateRelation", mock.Anything, tc.createRelationInputMock).
				Return(tc.repoExpResultMock, tc.expErr)

			friendsService := FriendService{userRepo: mockUserRepo, friendRepo: mockFriendsRepo}
			err := friendsService.CreateFriend(ctx, tc.input)
			var res bool
			if err == nil {
				res = true
			}

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)

			}
		})
	}

}

func TestService_GetAllFriendsOfUser(t *testing.T) {
	tcs := map[string]struct {
		input                GetAllFriendsInput
		expUserByEmailMock   models2.User
		expRelationIDsOfUser []int
		expListEmailByIDs    []string
		expResult            FriendListResponse
		expErr               error
	}{
		"success": {
			input: GetAllFriendsInput{
				Email: "common@example.com",
			},
			expUserByEmailMock: models2.User{
				ID:    3,
				Email: "common@example.com",
			},
			expRelationIDsOfUser: []int{1},
			expListEmailByIDs:    []string{"john@example.com", "andy@example.com"},
			expResult: FriendListResponse{
				Success: true,
				Friends: []string{"john@example.com", "andy@example.com"},
				Count:   2,
			},
		},
		"case email is empty": {
			input: GetAllFriendsInput{
				Email: "",
			},
			expUserByEmailMock:   models2.User{},
			expRelationIDsOfUser: []int{},
			expListEmailByIDs:    []string{},
			expResult: FriendListResponse{
				Success: false,
				Friends: []string{},
				Count:   0,
			},
			expErr: errors.New("Email cannot be empty"),
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
		expUserByEmailMock1 models2.User
		expUserByEmailMock2 models2.User
		expIdsFirstEmail    []int
		expIdsSecondEmail   []int
		expIntersectionIDs  []int
		expListEmail        []string
		expResult           FriendListResponse
		expErr              error
	}{
		"success": {
			input: CommonFriendsInput{
				RequesterEmail: "john@example.com",
				TargetEmail:    "andy@example.com",
			},
			expUserByEmailMock1: models2.User{
				ID:    1,
				Email: "john@example.com",
			},
			expUserByEmailMock2: models2.User{
				ID:    2,
				Email: "andy@example.com",
			},
			expIdsFirstEmail:   []int{3},
			expIdsSecondEmail:  []int{3},
			expIntersectionIDs: []int{3},
			expListEmail:       []string{"common@example.com"},
			expResult: FriendListResponse{
				Success: true,
				Friends: []string{"common@example.com"},
				Count:   1,
			},
		},
		"case first email is empty": {
			input: CommonFriendsInput{
				TargetEmail: "andy@example.com",
			},
			expUserByEmailMock1: models2.User{},
			expUserByEmailMock2: models2.User{
				ID:    2,
				Email: "andy@example.com",
			},
			expIdsFirstEmail:   []int{},
			expIdsSecondEmail:  []int{},
			expIntersectionIDs: []int{},
			expListEmail:       []string{},
			expResult: FriendListResponse{
				Success: false,
				Friends: []string{},
				Count:   0,
			},
			expErr: errors.New("Email cannot be empty"),
		},
		"case second email is empty": {
			input: CommonFriendsInput{
				RequesterEmail: "john@example.com",
			},
			expUserByEmailMock1: models2.User{
				ID:    1,
				Email: "john@example.com",
			},
			expUserByEmailMock2: models2.User{},
			expIdsFirstEmail:    []int{},
			expIdsSecondEmail:   []int{},
			expIntersectionIDs:  []int{},
			expListEmail:        []string{},
			expResult: FriendListResponse{
				Success: false,
				Friends: []string{},
				Count:   0,
			},
			expErr: errors.New("Email cannot be empty"),
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

			//mockUtils := new(utilsmock.UtilsInf)
			//mockUtils.On("Intersection", tc.expIdsFirstEmail, tc.expIdsSecondEmail).
			//	Return(tc.expIntersectionIDs, tc.expErr)

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
		userRepoExpResultMock1  models2.User
		userRepoExpResultMock2  models2.User
		createRelationInputMock models2.UserFriend
		repoExpResultMock       bool
		isExistMock             bool
		expResult               bool
		expErr                  error
	}{
		"success": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "lisa@example.com",
			},
			repoExpResultMock: true,
			isExistMock:       false,
			userRepoExpResultMock1: models2.User{
				ID:    3,
				Email: "common@example.com",
			},
			userRepoExpResultMock2: models2.User{
				ID:    4,
				Email: "lisa@example.com",
			},
			createRelationInputMock: models2.UserFriend{
				RequesterID:  3,
				TargetID:     4,
				RelationType: null.IntFrom(utils.Subscribe),
			},
			expResult: true,
		},
		"case requester users not found": {
			input: CreateRelationsInput{
				RequesterEmail: "",
				TargetEmail:    "lisa@example.com",
			},
			repoExpResultMock:      false,
			isExistMock:            false,
			userRepoExpResultMock1: models2.User{},
			userRepoExpResultMock2: models2.User{
				ID:    4,
				Email: "lisa@example.com",
			},
			createRelationInputMock: models2.UserFriend{
				RequesterID:  0,
				TargetID:     4,
				RelationType: null.IntFrom(utils.Subscribe),
			},
			expResult: false,
			expErr:    errors.New("Email cannot be empty"),
		},
		"case addressee users not found": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "",
			},
			repoExpResultMock: false,
			isExistMock:       false,
			userRepoExpResultMock1: models2.User{
				ID:    3,
				Email: "common@example.com",
			},
			userRepoExpResultMock2: models2.User{},
			createRelationInputMock: models2.UserFriend{
				RequesterID:  3,
				RelationType: null.IntFrom(utils.FriendRelation),
			},
			expResult: false,
			expErr:    errors.New("Email cannot be empty"),
		},
		"case error create friends because friend friends is existed ": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "lisa@example.com",
			},
			repoExpResultMock: false,
			isExistMock:       true,
			userRepoExpResultMock1: models2.User{
				ID:    3,
				Email: "common@example.com",
			},
			userRepoExpResultMock2: models2.User{
				ID:    4,
				Email: "lisa@example.com",
			},
			createRelationInputMock: models2.UserFriend{
				RequesterID:  3,
				TargetID:     4,
				RelationType: null.IntFrom(utils.Subscribe),
			},
			expResult: false,
			expErr:    errors.New("Email cannot be empty"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUserRepo := new(user.UsersRepoMock)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.RequesterEmail).
				Return(tc.userRepoExpResultMock1, tc.expErr)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.TargetEmail).
				Return(tc.userRepoExpResultMock2, tc.expErr)

			mockRelationRepo := new(friend.FriendsRepoMock)
			mockRelationRepo.On("IsRelationExist", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.isExistMock, tc.expErr)
			mockRelationRepo.On("CreateRelation", mock.Anything, tc.createRelationInputMock).
				Return(tc.repoExpResultMock, tc.expErr)

			relationService := FriendService{friendRepo: mockRelationRepo, userRepo: mockUserRepo}

			var res bool
			err := relationService.CreateSubscription(ctx, tc.input)
			if err == nil {
				res = true
			}
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)

			}
		})
	}

}

func TestService_CreateBlockRelation(t *testing.T) {
	tcs := map[string]struct {
		input                   CreateRelationsInput
		userRepoExpResultMock1  models2.User
		userRepoExpResultMock2  models2.User
		createRelationInputMock models2.UserFriend
		repoExpResultMock       bool
		isExistMock             bool
		expResult               bool
		expErr                  error
	}{
		"success": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "lisa@example.com",
			},
			repoExpResultMock: true,
			isExistMock:       false,
			userRepoExpResultMock1: models2.User{
				ID:    3,
				Email: "common@example.com",
			},
			userRepoExpResultMock2: models2.User{
				ID:    4,
				Email: "lisa@example.com",
			},
			createRelationInputMock: models2.UserFriend{
				RequesterID:  3,
				TargetID:     4,
				RelationType: null.IntFrom(utils.Blocked),
			},
			expResult: true,
		},
		"case requester users not found": {
			input: CreateRelationsInput{
				RequesterEmail: "",
				TargetEmail:    "lisa@example.com",
			},
			repoExpResultMock:      false,
			isExistMock:            false,
			userRepoExpResultMock1: models2.User{},
			userRepoExpResultMock2: models2.User{
				ID:    4,
				Email: "lisa@example.com",
			},
			createRelationInputMock: models2.UserFriend{
				RequesterID:  0,
				TargetID:     4,
				RelationType: null.IntFrom(utils.Blocked),
			},
			expResult: false,
			expErr:    errors.New("Email cannot be empty"),
		},
		"case addressee users not found": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "",
			},
			repoExpResultMock: false,
			isExistMock:       false,
			userRepoExpResultMock1: models2.User{
				ID:    3,
				Email: "common@example.com",
			},
			userRepoExpResultMock2: models2.User{},
			createRelationInputMock: models2.UserFriend{
				RequesterID:  3,
				RelationType: null.IntFrom(utils.Blocked),
			},
			expResult: false,
			expErr:    errors.New("Email cannot be empty"),
		},
		"case error create friends because friend friends is existed ": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				TargetEmail:    "lisa@example.com",
			},
			repoExpResultMock: false,
			isExistMock:       true,
			userRepoExpResultMock1: models2.User{
				ID:    3,
				Email: "common@example.com",
			},
			userRepoExpResultMock2: models2.User{
				ID:    4,
				Email: "lisa@example.com",
			},
			createRelationInputMock: models2.UserFriend{
				RequesterID:  3,
				TargetID:     4,
				RelationType: null.IntFrom(utils.Blocked),
			},
			expResult: false,
			expErr:    errors.New("Email cannot be empty"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUsersRepo := new(user.UsersRepoMock)
			mockUsersRepo.On("GetUserByEmail", mock.Anything, tc.input.RequesterEmail).
				Return(tc.userRepoExpResultMock1, tc.expErr)
			mockUsersRepo.On("GetUserByEmail", mock.Anything, tc.input.TargetEmail).
				Return(tc.userRepoExpResultMock2, tc.expErr)

			mockFriendRepo := new(friend.FriendsRepoMock)
			mockFriendRepo.On("IsRelationExist", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.isExistMock, tc.expErr)
			mockFriendRepo.On("CreateRelation", mock.Anything, tc.createRelationInputMock).
				Return(tc.repoExpResultMock, tc.expErr)
			mockFriendRepo.On("DeleteRelation", mock.Anything, tc.userRepoExpResultMock1.ID, tc.userRepoExpResultMock2.ID, utils.Subscribe).
				Return(nil, tc.expErr)

			relationService := FriendService{friendRepo: mockFriendRepo, userRepo: mockUsersRepo}

			var res bool
			err := relationService.CreateBlock(ctx, tc.input)
			if err == nil {
				res = true
			}

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)

			}
		})
	}

}

func TestService_GetEmailReceive(t *testing.T) {
	tcs := map[string]struct {
		input                   EmailReceiveInput
		expUserByEmailMock      models2.User
		expIdsEmail             []int
		expRequesterIdsRelation []int
		expEmailFromText        []string
		expUserIDByEmail        []int
		expUserIDBlock          []int
		expReceiveIDs           []int
		expUniqueSlice          []int

		expListEmail       []string
		expListEmailResult []string
		expResult          EmailReceiveResponse
		expErr             error
	}{
		"success": {
			input: EmailReceiveInput{
				Sender: "andy@example.com",
				Text:   "Hello World! lisa@example.com",
			},
			expUserByEmailMock: models2.User{
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
			expResult: EmailReceiveResponse{
				Success:    true,
				Recipients: []string{"common@example.com", "lisa@example.com"},
			},
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
