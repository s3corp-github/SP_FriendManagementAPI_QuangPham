package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/api/rest"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/controller/user"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/controller/user/mocks"
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
			expErr:  rest.ErrNameInvalid,
		},
	}

	tcsMockUserServ := map[string]struct {
		result user.UsersResponse
		err    error
	}{
		"success": {
			result: user.UsersResponse{
				ID:       15,
				Email:    "nhutquang23@gmail.com",
				Phone:    "031544284",
				IsActive: false,
			},
		},
		"error name cannot blank": {},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(tc.input))
			res := httptest.NewRecorder()

			// mock data UsersService
			mockUserSv := new(mocks.UserServ)
			userHandler := UsersHandler{mockUserSv}
			mockUserSv.On("CreateUser", mock.Anything, mock.AnythingOfType("user.CreateUserInput")).Return(
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
		result user.UsersEmailResponse
		err    error
	}{
		"success": {
			result: user.UsersEmailResponse{
				Email: []string{
					"andy@example.com",
					"john@example.com",
					"common@example.com",
					"lisa@example.com",
				},
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/users", nil)
			res := httptest.NewRecorder()

			// mock data UsersService
			mockUserSv := new(mocks.UserServ)
			userHandler := UsersHandler{mockUserSv}
			mockUserSv.On("GetListUser", mock.Anything).Return(
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
		expResult user.CreateUserInput
		expErr    error
	}{
		"success": {
			input: UsersRequest{
				Email:    "nhutquang23@gmail.com",
				Phone:    "02312545678",
				IsActive: true,
			},
			expResult: user.CreateUserInput{
				Email:    "nhutquang23@gmail.com",
				Phone:    "02312545678",
				IsActive: true,
			},
		},
		"case email cannot be blank": {
			input: UsersRequest{
				Phone:    "0343450044",
				IsActive: true,
			},
			expErr: rest.ErrEmailCannotBeBlank,
		},
		"case invalidEmail": {
			input: UsersRequest{
				Email:    "Thang12344email",
				Phone:    "0343450044",
				IsActive: true,
			},
			expErr: rest.ErrInvalidEmail,
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
