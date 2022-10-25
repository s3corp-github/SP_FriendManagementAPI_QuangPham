package user

import (
	"github.com/quangpham789/golang-assessment/api/internal/api/rest/errors"
	"github.com/quangpham789/golang-assessment/api/internal/controller/user"
	"github.com/quangpham789/golang-assessment/api/internal/controller/user/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
		"error email cannot blank": {
			input:   `{"email":"", "phone":"031544284", "is_active": false}`,
			expBody: "{\"message\":\"Email cannot be empty\"}\n",
			expCode: http.StatusBadRequest,
			expErr:  errors.ErrNameCannotBeBlank,
		},
	}

	tcsMockUserServ := map[string]struct {
		result user.UserResponse
		err    error
	}{
		"success": {
			result: user.UserResponse{
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

			// mock data UserService
			mockUserSv := new(mocks.UserServ)
			userHandler := UserHandler{mockUserSv}
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

func TestHandler_ValidateUserInput(t *testing.T) {
	tcs := map[string]struct {
		input     UserRequest
		expResult user.CreateUserInput
		expErr    error
	}{
		"success": {
			input: UserRequest{
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
		"email cannot be blank": {
			input: UserRequest{
				Phone:    "0343450044",
				IsActive: true,
			},
			expErr: errors.ErrEmailCannotBeBlank,
		},
		"invalidEmail": {
			input: UserRequest{
				Email:    "Thang12344email",
				Phone:    "0343450044",
				IsActive: true,
			},
			expErr: errors.ErrInvalidEmail,
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
