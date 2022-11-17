package rest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/service/users"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateUser(t *testing.T) {
	tcs := map[string]struct {
		input   string
		expCode int
		expBody string
		expErr  error
	}{
		"success": {
			input: `{"email":"nhutquang23@gmail.com", "phone":"032548784", "is_active": false}`,
			expBody: `{"ID":15,"Email":"nhutquang23@gmail.com","Phone":"031544284","IsActive":false}
`,
			expCode: http.StatusCreated,
		},
		"case error email cannot blank": {
			input:   `{"email":"", "phone":"031544284", "is_active": false}`,
			expBody: "{\"message\":\"Email cannot be empty\"}\n",
			expCode: http.StatusBadRequest,
			expErr:  ErrInvalidName,
		},
	}

	tcsMockUserServ := map[string]struct {
		result users.UserResponse
		err    error
	}{
		"success": {
			result: users.UserResponse{
				ID:    15,
				Email: "nhutquang23@gmail.com",
			},
		},
		"error name cannot blank": {},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(tc.input))
			res := httptest.NewRecorder()

			// mock data UserService
			mockUserSv := new(users.IServiceMock)
			userHandler := Handler{userService: mockUserSv}
			mockUserSv.On("CreateUser", mock.Anything, mock.AnythingOfType("users.CreateUserInput")).Return(
				tcsMockUserServ[desc].result, tcsMockUserServ[desc].err)

			handler := http.HandlerFunc(userHandler.CreateUser)
			handler.ServeHTTP(res, req)

			// test cases
			if tc.expErr != nil {
				require.Equal(t, res.Code, tc.expCode)

				require.Equal(t, res.Body.String(), tc.expBody)
			} else {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, res.Body.String(), tc.expBody)
			}

		})
	}
}

func TestHandler_GetListUser(t *testing.T) {
	tcs := map[string]struct {
		expCode int
		expBody string
		expErr  error
	}{
		"success": {
			expBody: `{"Email":["andy@example.com","john@example.com","common@example.com","lisa@example.com"]}
`,
			expCode: http.StatusOK,
		},
	}

	tcsMockUserServ := map[string]struct {
		result []users.UserEmailResponse
		err    error
	}{
		"success": {
			result: []users.UserEmailResponse{
				{
					Email: "andy@example.com",
					Name:  "andy",
				},
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/users", nil)
			res := httptest.NewRecorder()

			// mock data UserService
			mockUserSv := new(users.IServiceMock)
			userHandler := Handler{userService: mockUserSv}
			mockUserSv.On("GetUsers", mock.Anything).Return(
				tcsMockUserServ[desc].result, tcsMockUserServ[desc].err)

			handler := http.HandlerFunc(userHandler.GetListUser)
			handler.ServeHTTP(res, req)

			// test cases
			if tc.expErr != nil {
				require.Equal(t, res.Code, tc.expCode)

				require.Equal(t, res.Body.String(), tc.expBody)
			} else {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, res.Body.String(), tc.expBody)
			}

		})
	}
}

func TestHandler_ValidateUserInput(t *testing.T) {
	tcs := map[string]struct {
		input     UsersRequest
		expResult users.CreateUserInput
		expErr    error
	}{
		"success": {
			input: UsersRequest{
				Email: "nhutquang23@gmail.com",
			},
			expResult: users.CreateUserInput{
				Email: "nhutquang23@gmail.com",
			},
		},
		"case email cannot be blank": {
			input:  UsersRequest{},
			expErr: ErrInvalidEmail,
		},
		"case invalidEmail": {
			input: UsersRequest{
				Email: "Thang12344email",
			},
			expErr: ErrInvalidEmail,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res, err := validateUserInput(tc.input)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}
}
