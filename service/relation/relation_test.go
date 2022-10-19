package relation

import (
	"context"
	"github.com/friendsofgo/errors"
	models "github.com/quangpham789/golang-assessment/models"
	mockrepo "github.com/quangpham789/golang-assessment/repository/mocks"
	"github.com/quangpham789/golang-assessment/utils"
	mocksutils "github.com/quangpham789/golang-assessment/utils/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"testing"
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
		addresseeID    int
		relationType   int
		isExistMock    bool
		expResult      bool
		expErr         error
	}{
		"success": {
			requesterID:  1,
			addresseeID:  3,
			relationType: 1,
			isExistMock:  true,
			expResult:    true,
		},
		"want failed is relations exists": {
			requesterID:  1,
			addresseeID:  3,
			relationType: 1,
			isExistMock:  true,
			expResult:    true,
			mockIsRelation: mockIsRelation{
				wantCall: true,
				expErr:   errors.New("something went wrong"),
			},
		},
		"case requesterID is 0": {
			addresseeID:  3,
			relationType: 1,
			isExistMock:  true,
			expErr:       errors.New("requesterId cannot be null"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			//mockUserRepo := new(mockrepo.UserRepo)
			mockRelationRepo := new(mockrepo.RelationsRepo)
			mockRelationRepo.On("IsRelationExist", mock.Anything, tc.requesterID, tc.addresseeID, tc.relationType).
				Return(tc.isExistMock, tc.expErr)

			//relationService := RelationsService{relationsRepository: mockRelationRepo, userRepository: mockUserRepo}
			//res, err := relationService.isRelationExist(ctx, tc.requesterID, tc.addresseeID, tc.relationType)
			res, err := isRelationExist(ctx, mockRelationRepo, tc.requesterID, tc.addresseeID, tc.relationType)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)

			}
			// Mock
			//defer func() {
			//	isRelationExistFunc = isRelationExist
			//}()
			//
			//if tc.mockIsRelation.wantCall {
			//	isRelationExistFunc = func(ctx context.Context, repo repository.RelationsRepo, requesterID int, addresseeID int, relationType int) (bool, error) {
			//		return tc.mockIsRelation.expResult, tc.mockIsRelation.expErr
			//	}
			//}

		})
	}

}

func TestService_CreateFriendRelation(t *testing.T) {
	tcs := map[string]struct {
		input                   CreateRelationsInput
		userRepoExpResultMock1  models.User
		userRepoExpResultMock2  models.User
		createRelationInputMock models.Relation
		repoExpResultMock       bool
		isExistMock             bool
		expResult               CreateRelationsResponse
		expErr                  error
	}{
		"success": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				AddresseeEmail: "lisa@example.com",
			},
			repoExpResultMock: true,
			isExistMock:       false,
			userRepoExpResultMock1: models.User{
				ID:       3,
				Email:    "common@example.com",
				IsActive: null.BoolFrom(true),
			},
			userRepoExpResultMock2: models.User{
				ID:       4,
				Email:    "lisa@example.com",
				IsActive: null.BoolFrom(true),
			},
			createRelationInputMock: models.Relation{
				RequesterID:    3,
				AddresseeID:    4,
				RequesterEmail: "common@example.com",
				AddresseeEmail: "lisa@example.com",
				RelationType:   utils.FriendRelation,
			},
			expResult: CreateRelationsResponse{
				Success: true,
			},
		},
		"requester user not found": {
			input: CreateRelationsInput{
				RequesterEmail: "",
				AddresseeEmail: "lisa@example.com",
			},
			repoExpResultMock:      false,
			isExistMock:            false,
			userRepoExpResultMock1: models.User{},
			userRepoExpResultMock2: models.User{
				ID:       4,
				Email:    "lisa@example.com",
				IsActive: null.BoolFrom(true),
			},
			createRelationInputMock: models.Relation{
				RequesterID:    0,
				AddresseeID:    4,
				RequesterEmail: "",
				AddresseeEmail: "lisa@example.com",
				RelationType:   utils.FriendRelation,
			},
			expResult: CreateRelationsResponse{
				Success: false,
			},
			expErr: errors.New("Email cannot be empty"),
		},
		"addressee user not found": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				AddresseeEmail: "",
			},
			repoExpResultMock: false,
			isExistMock:       false,
			userRepoExpResultMock1: models.User{
				ID:       3,
				Email:    "common@example.com",
				IsActive: null.BoolFrom(true),
			},
			userRepoExpResultMock2: models.User{},
			createRelationInputMock: models.Relation{
				RequesterID:    3,
				RequesterEmail: "common@example.com",
				AddresseeEmail: "",
				RelationType:   utils.FriendRelation,
			},
			expResult: CreateRelationsResponse{
				Success: false,
			},
			expErr: errors.New("Email cannot be empty"),
		},
		"error create relation because friend relation is existed ": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				AddresseeEmail: "lisa@example.com",
			},
			repoExpResultMock: false,
			isExistMock:       true,
			userRepoExpResultMock1: models.User{
				ID:       3,
				Email:    "common@example.com",
				IsActive: null.BoolFrom(true),
			},
			userRepoExpResultMock2: models.User{
				ID:       4,
				Email:    "lisa@example.com",
				IsActive: null.BoolFrom(true),
			},
			createRelationInputMock: models.Relation{
				RequesterID:    3,
				AddresseeID:    4,
				RequesterEmail: "common@example.com",
				AddresseeEmail: "lisa@example.com",
				RelationType:   utils.FriendRelation,
			},
			expResult: CreateRelationsResponse{
				Success: false,
			},
			expErr: errors.New("Email cannot be empty"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUserRepo := new(mockrepo.UserRepo)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.RequesterEmail).
				Return(tc.userRepoExpResultMock1, tc.expErr)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.AddresseeEmail).
				Return(tc.userRepoExpResultMock2, tc.expErr)

			mockRelationRepo := new(mockrepo.RelationsRepo)
			mockRelationRepo.On("IsRelationExist", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.isExistMock, tc.expErr)
			mockRelationRepo.On("CreateRelation", mock.Anything, tc.createRelationInputMock).
				Return(tc.repoExpResultMock, tc.expErr)

			relationService := RelationsService{relationsRepository: mockRelationRepo, userRepository: mockUserRepo}
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

func TestService_GetAllFriendsOfUser(t *testing.T) {
	tcs := map[string]struct {
		input                GetAllFriendsInput
		expUserByEmailMock   models.User
		expRelationIDsOfUser []int
		expListEmailByIDs    []string
		expResult            FriendListResponse
		expErr               error
	}{
		"success": {
			input: GetAllFriendsInput{
				Email: "common@example.com",
			},
			expUserByEmailMock: models.User{
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
			expUserByEmailMock:   models.User{},
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
			mockUserRepo := new(mockrepo.UserRepo)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.Email).
				Return(tc.expUserByEmailMock, tc.expErr)
			mockRelationRepo := new(mockrepo.RelationsRepo)
			mockRelationRepo.On("GetRelationIDsOfUser", mock.Anything, tc.expUserByEmailMock.ID, 1).
				Return(tc.expRelationIDsOfUser, tc.expErr)
			mockUserRepo.On("GetListEmailByIDs", mock.Anything, mock.Anything).
				Return(tc.expListEmailByIDs, tc.expErr)

			relationService := RelationsService{relationsRepository: mockRelationRepo, userRepository: mockUserRepo}
			res, err := relationService.GetAllFriendsOfUser(ctx, tc.input)
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
		expResult           FriendListResponse
		expErr              error
	}{
		"success": {
			input: CommonFriendsInput{
				FirstEmail:  "john@example.com",
				SecondEmail: "andy@example.com",
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
			expResult: FriendListResponse{
				Success: true,
				Friends: []string{"common@example.com"},
				Count:   1,
			},
		},
		"case first email is empty": {
			input: CommonFriendsInput{
				SecondEmail: "andy@example.com",
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
			expResult: FriendListResponse{
				Success: false,
				Friends: []string{},
				Count:   0,
			},
			expErr: errors.New("Email cannot be empty"),
		},
		"case second email is empty": {
			input: CommonFriendsInput{
				FirstEmail: "john@example.com",
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
			mockUserRepo := new(mockrepo.UserRepo)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.FirstEmail).
				Return(tc.expUserByEmailMock1, tc.expErr)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.SecondEmail).
				Return(tc.expUserByEmailMock2, tc.expErr)

			mockRelationRepo := new(mockrepo.RelationsRepo)
			mockRelationRepo.On("GetRelationIDsOfUser", mock.Anything, tc.expUserByEmailMock1.ID, 1).
				Return(tc.expIdsFirstEmail, tc.expErr)
			mockRelationRepo.On("GetRelationIDsOfUser", mock.Anything, tc.expUserByEmailMock2.ID, 1).
				Return(tc.expIdsSecondEmail, tc.expErr)

			mockUtils := new(mocksutils.UtilsInf)
			mockUtils.On("Intersection", tc.expIdsFirstEmail, tc.expIdsSecondEmail).
				Return(tc.expIntersectionIDs, tc.expErr)

			mockUserRepo.On("GetListEmailByIDs", mock.Anything, mock.Anything).
				Return(tc.expListEmail, tc.expErr)

			relationService := RelationsService{relationsRepository: mockRelationRepo, userRepository: mockUserRepo, utils: mockUtils}
			res, err := relationService.GetCommonFriendList(ctx, tc.input)
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
		createRelationInputMock models.Relation
		repoExpResultMock       bool
		isExistMock             bool
		expResult               CreateRelationsResponse
		expErr                  error
	}{
		"success": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				AddresseeEmail: "lisa@example.com",
			},
			repoExpResultMock: true,
			isExistMock:       false,
			userRepoExpResultMock1: models.User{
				ID:       3,
				Email:    "common@example.com",
				IsActive: null.BoolFrom(true),
			},
			userRepoExpResultMock2: models.User{
				ID:       4,
				Email:    "lisa@example.com",
				IsActive: null.BoolFrom(true),
			},
			createRelationInputMock: models.Relation{
				RequesterID:    3,
				AddresseeID:    4,
				RequesterEmail: "common@example.com",
				AddresseeEmail: "lisa@example.com",
				RelationType:   utils.Subscribe,
			},
			expResult: CreateRelationsResponse{
				Success: true,
			},
		},
		"requester user not found": {
			input: CreateRelationsInput{
				RequesterEmail: "",
				AddresseeEmail: "lisa@example.com",
			},
			repoExpResultMock:      false,
			isExistMock:            false,
			userRepoExpResultMock1: models.User{},
			userRepoExpResultMock2: models.User{
				ID:       4,
				Email:    "lisa@example.com",
				IsActive: null.BoolFrom(true),
			},
			createRelationInputMock: models.Relation{
				RequesterID:    0,
				AddresseeID:    4,
				RequesterEmail: "",
				AddresseeEmail: "lisa@example.com",
				RelationType:   utils.Subscribe,
			},
			expResult: CreateRelationsResponse{
				Success: false,
			},
			expErr: errors.New("Email cannot be empty"),
		},
		"addressee user not found": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				AddresseeEmail: "",
			},
			repoExpResultMock: false,
			isExistMock:       false,
			userRepoExpResultMock1: models.User{
				ID:       3,
				Email:    "common@example.com",
				IsActive: null.BoolFrom(true),
			},
			userRepoExpResultMock2: models.User{},
			createRelationInputMock: models.Relation{
				RequesterID:    3,
				RequesterEmail: "common@example.com",
				AddresseeEmail: "",
				RelationType:   utils.FriendRelation,
			},
			expResult: CreateRelationsResponse{
				Success: false,
			},
			expErr: errors.New("Email cannot be empty"),
		},
		"error create relation because friend relation is existed ": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				AddresseeEmail: "lisa@example.com",
			},
			repoExpResultMock: false,
			isExistMock:       true,
			userRepoExpResultMock1: models.User{
				ID:       3,
				Email:    "common@example.com",
				IsActive: null.BoolFrom(true),
			},
			userRepoExpResultMock2: models.User{
				ID:       4,
				Email:    "lisa@example.com",
				IsActive: null.BoolFrom(true),
			},
			createRelationInputMock: models.Relation{
				RequesterID:    3,
				AddresseeID:    4,
				RequesterEmail: "common@example.com",
				AddresseeEmail: "lisa@example.com",
				RelationType:   utils.Subscribe,
			},
			expResult: CreateRelationsResponse{
				Success: false,
			},
			expErr: errors.New("Email cannot be empty"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUserRepo := new(mockrepo.UserRepo)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.RequesterEmail).
				Return(tc.userRepoExpResultMock1, tc.expErr)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.AddresseeEmail).
				Return(tc.userRepoExpResultMock2, tc.expErr)

			mockRelationRepo := new(mockrepo.RelationsRepo)
			mockRelationRepo.On("IsRelationExist", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.isExistMock, tc.expErr)
			mockRelationRepo.On("CreateRelation", mock.Anything, tc.createRelationInputMock).
				Return(tc.repoExpResultMock, tc.expErr)

			relationService := RelationsService{relationsRepository: mockRelationRepo, userRepository: mockUserRepo}
			res, err := relationService.CreateSubscriptionRelation(ctx, tc.input)
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
		userRepoExpResultMock1  models.User
		userRepoExpResultMock2  models.User
		createRelationInputMock models.Relation
		repoExpResultMock       bool
		isExistMock             bool
		expResult               CreateRelationsResponse
		expErr                  error
	}{
		"success": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				AddresseeEmail: "lisa@example.com",
			},
			repoExpResultMock: true,
			isExistMock:       false,
			userRepoExpResultMock1: models.User{
				ID:       3,
				Email:    "common@example.com",
				IsActive: null.BoolFrom(true),
			},
			userRepoExpResultMock2: models.User{
				ID:       4,
				Email:    "lisa@example.com",
				IsActive: null.BoolFrom(true),
			},
			createRelationInputMock: models.Relation{
				RequesterID:    3,
				AddresseeID:    4,
				RequesterEmail: "common@example.com",
				AddresseeEmail: "lisa@example.com",
				RelationType:   utils.Block,
			},
			expResult: CreateRelationsResponse{
				Success: true,
			},
		},
		"requester user not found": {
			input: CreateRelationsInput{
				RequesterEmail: "",
				AddresseeEmail: "lisa@example.com",
			},
			repoExpResultMock:      false,
			isExistMock:            false,
			userRepoExpResultMock1: models.User{},
			userRepoExpResultMock2: models.User{
				ID:       4,
				Email:    "lisa@example.com",
				IsActive: null.BoolFrom(true),
			},
			createRelationInputMock: models.Relation{
				RequesterID:    0,
				AddresseeID:    4,
				RequesterEmail: "",
				AddresseeEmail: "lisa@example.com",
				RelationType:   utils.Block,
			},
			expResult: CreateRelationsResponse{
				Success: false,
			},
			expErr: errors.New("Email cannot be empty"),
		},
		"addressee user not found": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				AddresseeEmail: "",
			},
			repoExpResultMock: false,
			isExistMock:       false,
			userRepoExpResultMock1: models.User{
				ID:       3,
				Email:    "common@example.com",
				IsActive: null.BoolFrom(true),
			},
			userRepoExpResultMock2: models.User{},
			createRelationInputMock: models.Relation{
				RequesterID:    3,
				RequesterEmail: "common@example.com",
				AddresseeEmail: "",
				RelationType:   utils.Block,
			},
			expResult: CreateRelationsResponse{
				Success: false,
			},
			expErr: errors.New("Email cannot be empty"),
		},
		"error create relation because friend relation is existed ": {
			input: CreateRelationsInput{
				RequesterEmail: "common@example.com",
				AddresseeEmail: "lisa@example.com",
			},
			repoExpResultMock: false,
			isExistMock:       true,
			userRepoExpResultMock1: models.User{
				ID:       3,
				Email:    "common@example.com",
				IsActive: null.BoolFrom(true),
			},
			userRepoExpResultMock2: models.User{
				ID:       4,
				Email:    "lisa@example.com",
				IsActive: null.BoolFrom(true),
			},
			createRelationInputMock: models.Relation{
				RequesterID:    3,
				AddresseeID:    4,
				RequesterEmail: "common@example.com",
				AddresseeEmail: "lisa@example.com",
				RelationType:   utils.Block,
			},
			expResult: CreateRelationsResponse{
				Success: false,
			},
			expErr: errors.New("Email cannot be empty"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUserRepo := new(mockrepo.UserRepo)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.RequesterEmail).
				Return(tc.userRepoExpResultMock1, tc.expErr)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.AddresseeEmail).
				Return(tc.userRepoExpResultMock2, tc.expErr)

			mockRelationRepo := new(mockrepo.RelationsRepo)
			mockRelationRepo.On("IsRelationExist", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.isExistMock, tc.expErr)
			mockRelationRepo.On("CreateRelation", mock.Anything, tc.createRelationInputMock).
				Return(tc.repoExpResultMock, tc.expErr)
			mockRelationRepo.On("DeleteRelation", mock.Anything, tc.userRepoExpResultMock1.ID, tc.userRepoExpResultMock2.ID, utils.Subscribe).
				Return(nil, tc.expErr)

			relationService := RelationsService{relationsRepository: mockRelationRepo, userRepository: mockUserRepo}
			res, err := relationService.CreateBlockRelation(ctx, tc.input)
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
		expResult          EmailReceiveResponse
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
			expResult: EmailReceiveResponse{
				Success:    true,
				Recipients: []string{"common@example.com", "lisa@example.com"},
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockUserRepo := new(mockrepo.UserRepo)
			mockUserRepo.On("GetUserByEmail", mock.Anything, tc.input.Sender).
				Return(tc.expUserByEmailMock, tc.expErr)

			mockRelationRepo := new(mockrepo.RelationsRepo)
			mockRelationRepo.On("GetRelationIDsOfUser", mock.Anything, tc.expUserByEmailMock.ID, 1).
				Return(tc.expIdsEmail, tc.expErr)
			mockRelationRepo.On("GetRequesterIDRelation", mock.Anything, mock.Anything, mock.Anything).
				Return(tc.expRequesterIdsRelation, tc.expErr)

			mockUtils := new(mocksutils.UtilsInf)
			mockUtils.On("FindEmailFromText", tc.input.Text).
				Return(tc.expListEmail, tc.expErr)

			mockUserRepo.On("GetUserIDsByEmail", mock.Anything, tc.expListEmail).
				Return(tc.expUserIDByEmail, tc.expErr)
			//
			//mockUserRepo.On("GetRequesterIDRelation", mock.Anything, tc.expUserByEmailMock.ID, 3).
			//	Return(tc.expUserIDBlock, tc.expErr)

			mockUtils.On("UniqueSlice", mock.Anything).
				Return(tc.expReceiveIDs, tc.expErr)

			mockUtils.On("GetReceiverID", mock.Anything, mock.Anything).
				Return(tc.expReceiveIDs, tc.expErr)

			mockUserRepo.On("GetListEmailByIDs", mock.Anything, tc.expReceiveIDs).
				Return(tc.expListEmailResult, tc.expErr)

			relationService := RelationsService{relationsRepository: mockRelationRepo, userRepository: mockUserRepo, utils: mockUtils}
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
