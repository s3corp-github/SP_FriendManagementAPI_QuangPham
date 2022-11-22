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
		input      string
		givenInput users.UserEmail
		expCode    int
		expBody    string
		expErr     error
	}{
		"success": {
			input:      `{"email":"john@example.com", "name":"john"}`,
			givenInput: users.UserEmail{Name: "john", Email: "john@example.com"},
			expBody: `{"ID":15,"Email":"john@example.com"}
`,
			expCode: http.StatusCreated,
		},
		"case error email cannot blank": {
			input:      `{"email":"", "name":""}`,
			givenInput: users.UserEmail{},
			expBody:    "{\"message\":\"Invalid email address\"}\n",
			expCode:    http.StatusBadRequest,
			expErr:     ErrInvalidName,
		},
	}

	tcsMockUserServ := map[string]struct {
		result users.UserResponse
		err    error
	}{
		"success": {
			result: users.UserResponse{
				ID:    15,
				Email: "john@example.com",
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
			mockUserSv.On("CreateUser", mock.Anything, tc.givenInput).Return(
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
			expBody: "[{\"Email\":\"andy@example.com\",\"Name\":\"andy\"}]\n",
			expCode: http.StatusOK,
		},
	}

	tcsMockUserServ := map[string]struct {
		result []users.UserEmail
		err    error
	}{
		"success": {
			result: []users.UserEmail{
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

			handler := http.HandlerFunc(userHandler.GetUsers)
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
		expResult users.UserEmail
		expErr    error
	}{
		"success": {
			input: UsersRequest{
				Email: "john@example.com",
			},
			expResult: users.UserEmail{
				Email: "john@example.com",
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
