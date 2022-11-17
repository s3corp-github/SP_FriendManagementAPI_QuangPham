package rest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/service/friends"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateFriendsRelation(t *testing.T) {
	type mockSrcCreateFriend struct {
		givenInput friends.CreateRelationsInput
		wantCall   bool
		expErr     error
	}
	tcs := map[string]struct {
		input               string
		expBody             string
		expCode             int
		mockSrcCreateFriend mockSrcCreateFriend
	}{
		"success": {
			input:   `{"friends":["andy@example.com", "common@example.com"]}`,
			expBody: "{\"success\":true}\n",
			mockSrcCreateFriend: mockSrcCreateFriend{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
			},
			expCode: http.StatusCreated,
		},
		"case input empty requester": {
			input:   `{"friends":["", "common@example.com"]}`,
			expBody: "{\"message\":\"Email cannot be empty\"}\n",
			mockSrcCreateFriend: mockSrcCreateFriend{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
				expErr:   ErrInvalidEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case input email invalid": {
			input:   `{"friends":["andy.com", "common@example.com"]}`,
			expBody: "{\"message\":\"Invalid email address\"}\n",
			mockSrcCreateFriend: mockSrcCreateFriend{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy.com",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
				expErr:   ErrInvalidEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case requester email and addressee email is same": {
			input:   `{"friends":["common@example.com", "common@example.com"]}`,
			expBody: "{\"message\":\"Requester email and target email must not be the same\"}\n",
			mockSrcCreateFriend: mockSrcCreateFriend{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "common@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
				expErr:   ErrRequesterAndTargetEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case requester email not exist": {
			input:   `{"friends":["andy123@example.com", "common@example.com"]}`,
			expBody: "{\"success\":false}\n",
			mockSrcCreateFriend: mockSrcCreateFriend{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy123@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
				expErr:   friends.ErrRequestEmailInvalid,
			},
			expCode: http.StatusCreated,
		},
		"case unable create friends": {
			input:   `{"friends":["andy@example.com", "common@example.com"]}`,
			expBody: "{\"success\":false}\n",
			mockSrcCreateFriend: mockSrcCreateFriend{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
			},
			expCode: http.StatusCreated,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/relations/createfriendrelation",
				strings.NewReader(tc.input))
			res := httptest.NewRecorder()

			// mock data RelationService
			mockFriendsSv := new(friends.IServiceMock)
			handlers := Handler{friendService: mockFriendsSv}
			if tc.mockSrcCreateFriend.wantCall {
				mockFriendsSv.On("CreateFriend", mock.Anything,
					tc.mockSrcCreateFriend.givenInput).Return(tc.mockSrcCreateFriend.expErr)
			}

			handler := http.HandlerFunc(handlers.CreateFriends)
			handler.ServeHTTP(res, req)

			// test cases
			if tc.mockSrcCreateFriend.expErr != nil {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, res.Body.String(), tc.expBody)
			} else {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, res.Body.String(), tc.expBody)
			}

		})
	}
}
func TestHandler_validateRelationInput(t *testing.T) {
	tcs := map[string]struct {
		input     FriendsRequest
		expResult friends.CreateRelationsInput
		expErr    error
	}{
		"success": {
			input: FriendsRequest{
				Friends: []string{"andy@example.com", "common@example.com"},
			},
			expResult: friends.CreateRelationsInput{
				RequesterEmail: "andy@example.com",
				TargetEmail:    "common@example.com",
			},
		},
		"email cannot be blank": {
			input: FriendsRequest{
				Friends: []string{"", "common@example.com"},
			},
			expErr: ErrInvalidEmail,
		},
		"invalidEmail": {
			input: FriendsRequest{
				Friends: []string{"andyexample.com", "common@example.com"},
			},
			expErr: ErrInvalidEmail,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res, err := validateRelationInput(tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}
}

func TestHandler_validateSubAndBlockRelationInput(t *testing.T) {
	tcs := map[string]struct {
		input     CreateFriendsRequest
		expResult friends.CreateRelationsInput
		expErr    error
	}{
		"success": {
			input: CreateFriendsRequest{
				Requester: "andy@example.com",
				Target:    "john@example.com",
			},
			expResult: friends.CreateRelationsInput{
				RequesterEmail: "andy@example.com",
				TargetEmail:    "john@example.com",
			},
		},
		"case email cannot be blank": {
			input: CreateFriendsRequest{
				Requester: "andy@example.com",
			},
			expErr: ErrInvalidEmail,
		},
		"case requester email invalidEmail": {
			input: CreateFriendsRequest{
				Requester: "andyexample.com",
				Target:    "john@example.com",
			},
			expErr: ErrInvalidEmail,
		},
		"case requester email is empty": {
			input: CreateFriendsRequest{
				Requester: "",
				Target:    "john@example.com",
			},
			expErr: ErrInvalidEmail,
		},
		"case addressee email is empty": {
			input: CreateFriendsRequest{
				Requester: "andy@example.com",
				Target:    "",
			},
			expErr: ErrInvalidEmail,
		},
		"case addressee email invalidEmail": {
			input: CreateFriendsRequest{
				Requester: "andy@example.com",
				Target:    "johnexample.com",
			},
			expErr: ErrInvalidEmail,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res, err := validateSubAndBlockRelationInput(tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}
}

func TestHandler_validateGetRelationInput(t *testing.T) {
	tcs := map[string]struct {
		input     GetFriendsRequest
		expResult friends.GetAllFriendsInput
		expErr    error
	}{
		"success": {
			input: GetFriendsRequest{
				Email: "andy@example.com",
			},
			expResult: friends.GetAllFriendsInput{
				Email: "andy@example.com",
			},
		},
		"case email cannot be blank": {
			input: GetFriendsRequest{
				Email: "",
			},
			expErr: ErrInvalidEmail,
		},
		"case requester email invalidEmail": {
			input: GetFriendsRequest{
				Email: "andyexample.com",
			},
			expErr: ErrInvalidEmail,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res, err := validateGetRelationInput(tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}
}
func TestHandler_validateRelationCommonInput(t *testing.T) {
	tcs := map[string]struct {
		input     FriendsRequest
		expResult friends.CommonFriendsInput
		expErr    error
	}{
		"success": {
			input: FriendsRequest{
				Friends: []string{"andy@example.com", "john@example.com"},
			},
			expResult: friends.CommonFriendsInput{
				RequesterEmail: "andy@example.com",
				TargetEmail:    "john@example.com",
			},
		},
		"case not enough param": {
			input: FriendsRequest{
				Friends: []string{"john@example.com"},
			},
			expErr: ErrInvalidBodyRequest,
		},
		"case requester email invalidEmail": {
			input: FriendsRequest{
				Friends: []string{"andyexample.com", "john@example.com"},
			},
			expErr: ErrInvalidEmail,
		},
		"case requester email is empty": {
			input: FriendsRequest{
				Friends: []string{"", "john@example.com"},
			},
			expErr: ErrInvalidEmail,
		},
		"case addressee email is empty": {
			input: FriendsRequest{
				Friends: []string{"andy@example.com", ""},
			},
			expErr: ErrInvalidEmail,
		},
		"case addressee email invalidEmail": {
			input: FriendsRequest{
				Friends: []string{"andy@example.com", "johnexample.com"},
			},
			expErr: ErrInvalidEmail,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res, err := validateRelationCommonInput(tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}
}

func TestHandler_validateEmailReceiveInput(t *testing.T) {
	tcs := map[string]struct {
		input     EmailReceiveRequest
		expResult friends.EmailReceiveInput
		expErr    error
	}{
		"success": {
			input: EmailReceiveRequest{
				Sender: "andy@example.com",
				Text:   "Hello World! kate@example.com",
			},
			expResult: friends.EmailReceiveInput{
				Sender: "andy@example.com",
				Text:   "Hello World! kate@example.com",
			},
		},

		"case sender email invalidEmail": {
			input: EmailReceiveRequest{
				Sender: "andyexample.com",
				Text:   "Hello World! kate@example.com",
			},
			expErr: ErrInvalidEmail,
		},
		"case sender email is empty": {
			input: EmailReceiveRequest{
				Sender: "",
				Text:   "Hello World! kate@example.com",
			},
			expErr: ErrInvalidEmail,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res, err := validateEmailReceiveInput(tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}
}

func TestHandler_GetAllFriendOfUser(t *testing.T) {
	type mockSrcGetFriend struct {
		givenInput friends.GetAllFriendsInput
		wantCall   bool
		expResult  friends.FriendListResponse
		expErr     error
	}
	tcs := map[string]struct {
		input            string
		expCode          int
		expBody          string
		mockSrcGetFriend mockSrcGetFriend
	}{
		"success": {
			input:   `{"email":"andy@example.com"}`,
			expBody: "{\"success\":true,\"friends\":[\"common@example.com\"],\"count\":1}\n",
			mockSrcGetFriend: mockSrcGetFriend{
				givenInput: friends.GetAllFriendsInput{
					Email: "andy@example.com",
				},
				wantCall: true,
				expResult: friends.FriendListResponse{
					Success: true,
					Friends: []string{"common@example.com"},
					Count:   1,
				},
			},
			expCode: http.StatusOK,
		},
		"case email input is empty": {
			input:   `{"email":""}`,
			expBody: "{\"message\":\"Email cannot be empty\"}\n",
			mockSrcGetFriend: mockSrcGetFriend{
				givenInput: friends.GetAllFriendsInput{
					Email: "",
				},
				wantCall: true,
				expErr:   ErrInvalidEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case email input is invalid": {
			input:   `{"email":"andyexample.com"}`,
			expBody: "{\"message\":\"Invalid email address\"}\n",
			mockSrcGetFriend: mockSrcGetFriend{
				givenInput: friends.GetAllFriendsInput{
					Email: "andyexample.com",
				},
				wantCall: true,
				expErr:   ErrInvalidEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case email is not exist": {
			input:   `{"email":"andy123@example.com"}`,
			expBody: "{\"success\":true,\"friends\":[],\"count\":0}\n",
			mockSrcGetFriend: mockSrcGetFriend{
				givenInput: friends.GetAllFriendsInput{
					Email: "andy123@example.com",
				},
				wantCall: true,
				expResult: friends.FriendListResponse{
					Success: true,
					Friends: []string{},
				},
			},
			expCode: http.StatusOK,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/relations/getfriendfromemail",
				strings.NewReader(tc.input))
			res := httptest.NewRecorder()

			// mock data RelationService
			mockRelationSv := new(friends.IServiceMock)
			relationHandler := Handler{friendService: mockRelationSv}

			if tc.mockSrcGetFriend.wantCall {
				mockRelationSv.On("GetAllFriendsOfUser", mock.Anything,
					tc.mockSrcGetFriend.givenInput).Return(
					tc.mockSrcGetFriend.expResult, tc.mockSrcGetFriend.expErr)
			}

			handler := http.HandlerFunc(relationHandler.GetFriends)
			handler.ServeHTTP(res, req)

			// test cases
			if tc.mockSrcGetFriend.expErr != nil {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, tc.expBody, res.Body.String())
			} else {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, res.Body.String(), tc.expBody)
			}

		})
	}
}

func TestHandler_GetCommonFriend(t *testing.T) {
	type mockSrcGetCommonFriend struct {
		givenInput friends.CommonFriendsInput
		wantCall   bool
		expResult  friends.FriendListResponse
		expErr     error
	}
	tcs := map[string]struct {
		input                  string
		expCode                int
		expBody                string
		mockSrcGetCommonFriend mockSrcGetCommonFriend
	}{
		"success": {
			input:   `{"friends":["andy@example.com", "john@example.com"]}`,
			expBody: "{\"success\":true,\"friends\":[\"common@example.com\"],\"count\":1}\n",
			mockSrcGetCommonFriend: mockSrcGetCommonFriend{
				givenInput: friends.CommonFriendsInput{
					RequesterEmail: "andy@example.com",
					TargetEmail:    "john@example.com",
				},
				wantCall: true,
				expResult: friends.FriendListResponse{
					Success: true,
					Friends: []string{"common@example.com"},
					Count:   1,
				},
			},
			expCode: http.StatusOK,
		},
		"case email input is empty": {
			input:   `{"friends":["", "john@example.com"]}`,
			expBody: "{\"message\":\"Email cannot be empty\"}\n",
			mockSrcGetCommonFriend: mockSrcGetCommonFriend{
				givenInput: friends.CommonFriendsInput{
					RequesterEmail: "",
					TargetEmail:    "john@example.com",
				},
				wantCall: true,
				expErr:   ErrInvalidEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case email input is invalid": {
			input:   `{"friends":["andyexample.com", "john@example.com"]}`,
			expBody: "{\"message\":\"Invalid email address\"}\n",
			mockSrcGetCommonFriend: mockSrcGetCommonFriend{
				givenInput: friends.CommonFriendsInput{
					RequesterEmail: "andyexample",
					TargetEmail:    "john@example.com",
				},
				wantCall: true,
				expErr:   ErrInvalidEmail,
			},
			expCode: http.StatusBadRequest,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/relations/getcommonfriends",
				strings.NewReader(tc.input))
			res := httptest.NewRecorder()

			// mock data RelationService
			mockRelationSv := new(friends.IServiceMock)
			relationHandler := Handler{friendService: mockRelationSv}

			if tc.mockSrcGetCommonFriend.wantCall {
				mockRelationSv.On("GetCommonFriendList", mock.Anything,
					tc.mockSrcGetCommonFriend.givenInput).Return(
					tc.mockSrcGetCommonFriend.expResult, tc.mockSrcGetCommonFriend.expErr)
			}

			handler := http.HandlerFunc(relationHandler.GetCommonFriends)
			handler.ServeHTTP(res, req)

			// test cases
			if tc.mockSrcGetCommonFriend.expErr != nil {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, tc.expBody, res.Body.String())
			} else {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, res.Body.String(), tc.expBody)
			}

		})
	}
}

func TestHandler_CreateSubscriptionRelation(t *testing.T) {

	type mockSrcCreateSub struct {
		givenInput friends.CreateRelationsInput
		wantCall   bool
		expResult  bool
		expErr     error
	}

	tcs := map[string]struct {
		input            string
		expCode          int
		expBody          string
		mockSrcCreateSub mockSrcCreateSub
	}{
		"success": {
			input:   `{"requestor":"andy@example.com", "target":"common@example.com"}`,
			expBody: "{\"success\":true}\n",
			mockSrcCreateSub: mockSrcCreateSub{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall:  true,
				expResult: true,
			},
			expCode: http.StatusCreated,
		},
		"case input empty requester": {
			input:   `{"requestor":"", "target":"common@example.com"}`,
			expBody: "{\"message\":\"Email cannot be empty\"}\n",
			mockSrcCreateSub: mockSrcCreateSub{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
				expErr:   ErrInvalidEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case input email invalid": {
			input:   `{"requestor":"andyexample.com", "target":"common@example.com"}`,
			expBody: "{\"message\":\"Invalid email address\"}\n",
			mockSrcCreateSub: mockSrcCreateSub{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
				expErr:   ErrInvalidEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case requester email and addressee email is same": {
			input:   `{"requestor":"common@example.com", "target":"common@example.com"}`,
			expBody: "{\"message\":\"Requester email and target email must not be the same\"}\n",
			mockSrcCreateSub: mockSrcCreateSub{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
				expErr:   ErrRequesterAndTargetEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case requester email not exist": {
			input:   `{"requestor":"andy123@example.com", "target":"common@example.com"}`,
			expBody: "{\"message\":\"Requester email not exist\"}\n",
			mockSrcCreateSub: mockSrcCreateSub{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy123@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
				expErr:   friends.ErrRequestEmailInvalid,
			},
			expCode: http.StatusInternalServerError,
		},
		"case unable create friends": {
			input:   `{"requestor":"andy123@example.com", "target":"common@example.com"}`,
			expBody: "{\"success\":false}\n",
			mockSrcCreateSub: mockSrcCreateSub{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy123@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall:  true,
				expResult: false,
			},
			expCode: http.StatusCreated,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/relations/createsubscriptionrelation",
				strings.NewReader(tc.input))
			res := httptest.NewRecorder()

			// mock
			mockRelationSv := new(friends.IServiceMock)
			friendsHandler := Handler{friendService: mockRelationSv}

			if tc.mockSrcCreateSub.wantCall {
				mockRelationSv.On("CreateSubscription", mock.Anything, tc.mockSrcCreateSub.givenInput).
					Return(
						tc.mockSrcCreateSub.expResult, tc.mockSrcCreateSub.expErr)
			}
			//when
			handler := http.HandlerFunc(friendsHandler.CreateSubscription)
			handler.ServeHTTP(res, req)

			// then
			if tc.mockSrcCreateSub.expErr != nil {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, res.Body.String(), tc.expBody)
			} else {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, res.Body.String(), tc.expBody)
			}

		})
	}
}

func TestHandler_CreateBlockRelation(t *testing.T) {
	type mockSrcCreateBlock struct {
		givenInput friends.CreateRelationsInput
		wantCall   bool
		expResult  bool
		expErr     error
	}
	tcs := map[string]struct {
		input              string
		expCode            int
		expBody            string
		mockSrcCreateBlock mockSrcCreateBlock
	}{
		"success": {
			input:   `{"requestor":"andy@example.com", "target":"common@example.com"}`,
			expBody: "{\"success\":true}\n",
			mockSrcCreateBlock: mockSrcCreateBlock{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall:  true,
				expResult: true,
			},
			expCode: http.StatusCreated,
		},
		"case input empty requester": {
			input:   `{"requestor":"", "target":"common@example.com"}`,
			expBody: "{\"message\":\"Email cannot be empty\"}\n",
			mockSrcCreateBlock: mockSrcCreateBlock{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
				expErr:   ErrInvalidEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case input email invalid": {
			input:   `{"requestor":"andyexample.com", "target":"common@example.com"}`,
			expBody: "{\"message\":\"Invalid email address\"}\n",
			mockSrcCreateBlock: mockSrcCreateBlock{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andyexample.com",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
				expErr:   ErrInvalidEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case requester email and addressee email is same": {
			input:   `{"requestor":"common@example.com", "target":"common@example.com"}`,
			expBody: "{\"message\":\"Requester email and target email must not be the same\"}\n",
			mockSrcCreateBlock: mockSrcCreateBlock{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "common@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
				expErr:   ErrRequesterAndTargetEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case requester email not exist": {
			input:   `{"requestor":"andy123@example.com", "target":"common@example.com"}`,
			expBody: "{\"message\":\"Requester email not exist\"}\n",
			mockSrcCreateBlock: mockSrcCreateBlock{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy123@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall: true,
				expErr:   friends.ErrRequestEmailInvalid,
			},
			expCode: http.StatusInternalServerError,
		},
		"case unable create friend friends": {
			input:   `{"requestor":"andy123@example.com", "target":"common@example.com"}`,
			expBody: "{\"success\":false}\n",
			mockSrcCreateBlock: mockSrcCreateBlock{
				givenInput: friends.CreateRelationsInput{
					RequesterEmail: "andy123@example.com",
					TargetEmail:    "common@example.com",
				},
				wantCall:  true,
				expResult: false,
			},
			expCode: http.StatusCreated,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/relations/createblockrelation",
				strings.NewReader(tc.input))
			res := httptest.NewRecorder()

			// mock data RelationService
			mockRelationSv := new(friends.IServiceMock)
			friendsHandler := Handler{friendService: mockRelationSv}
			if tc.mockSrcCreateBlock.wantCall {
				mockRelationSv.On("CreateBlock", mock.Anything,
					tc.mockSrcCreateBlock.givenInput).Return(
					tc.mockSrcCreateBlock.expResult, tc.mockSrcCreateBlock.expErr)
			}

			handler := http.HandlerFunc(friendsHandler.CreateBlock)
			handler.ServeHTTP(res, req)

			// test cases
			if tc.mockSrcCreateBlock.expErr != nil {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, res.Body.String(), tc.expBody)
			} else {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, res.Body.String(), tc.expBody)
			}

		})
	}
}

func TestHandler_GetEmailReceive(t *testing.T) {
	type mockSrcGetEmailReceive struct {
		givenInput friends.EmailReceiveInput
		wantCall   bool
		expResult  friends.EmailReceiveResponse
		expErr     error
	}
	tcs := map[string]struct {
		input                  string
		expCode                int
		expBody                string
		mockSrcGetEmailReceive mockSrcGetEmailReceive
	}{
		"success": {
			input:   `{"sender":"andy@example.com", "text":"Hello World! lisa@example.com"}`,
			expBody: "{\"success\":true,\"recipients\":[\"common@example.com\",\"lisa@example.com\"]}\n",
			mockSrcGetEmailReceive: mockSrcGetEmailReceive{
				givenInput: friends.EmailReceiveInput{
					Sender: "andy@example.com",
					Text:   "Hello World! lisa@example.com",
				},
				wantCall: true,
				expResult: friends.EmailReceiveResponse{
					Success: true,
					Recipients: []string{
						"common@example.com",
						"lisa@example.com",
					},
				},
			},
			expCode: http.StatusOK,
		},
		"case sender email input is empty": {
			input:   `{"sender":"", "text":"Hello World! lisa@example.com"}`,
			expBody: "{\"message\":\"Email cannot be empty\"}\n",
			mockSrcGetEmailReceive: mockSrcGetEmailReceive{
				givenInput: friends.EmailReceiveInput{
					Sender: "",
					Text:   "Hello World! lisa@example.com",
				},
				wantCall: true,
				expErr:   ErrInvalidEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case sender email input is invalid": {
			input:   `{"sender":"andyexample.com", "text":"Hello World! lisa@example.com"}`,
			expBody: "{\"message\":\"Invalid email address\"}\n",
			mockSrcGetEmailReceive: mockSrcGetEmailReceive{
				givenInput: friends.EmailReceiveInput{
					Sender: "andyexample.com",
					Text:   "Hello World! lisa@example.com",
				},
				wantCall: true,
				expErr:   ErrInvalidEmail,
			},
			expCode: http.StatusBadRequest,
		},
		"case email is not exist": {
			input:   `{"sender":"andy123@example.com", "text":"Hello World! lisa@example.com"}`,
			expBody: "{\"message\":\"Email not exist\"}\n",
			mockSrcGetEmailReceive: mockSrcGetEmailReceive{
				givenInput: friends.EmailReceiveInput{
					Sender: "andy123@example.com",
					Text:   "Hello World! lisa@example.com",
				},
				wantCall: true,
				expErr:   friends.ErrRequestEmailInvalid,
			},
			expCode: http.StatusInternalServerError,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/relations/getemailreceive",
				strings.NewReader(tc.input))
			res := httptest.NewRecorder()

			// mock data RelationService
			mockRelationSv := new(friends.IServiceMock)
			relationHandler := Handler{friendService: mockRelationSv}

			if tc.mockSrcGetEmailReceive.wantCall {
				mockRelationSv.On("GetEmailReceive", mock.Anything,
					tc.mockSrcGetEmailReceive.givenInput).Return(
					tc.mockSrcGetEmailReceive.expResult, tc.mockSrcGetEmailReceive.expErr)
			}

			handler := http.HandlerFunc(relationHandler.GetEmailReceive)
			handler.ServeHTTP(res, req)

			// test cases
			if tc.mockSrcGetEmailReceive.expErr != nil {
				require.Equal(t, tc.expCode, res.Code)
				require.Equal(t, tc.expBody, res.Body.String())
			} else {
				require.Equal(t, tc.expCode, res.Code)
				require.Equal(t, tc.expBody, res.Body.String())
			}

		})
	}
}
