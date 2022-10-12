package handler

import (
	"github.com/quangpham789/golang-assessment/service"
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
			input: `{"first_name":"Quang", "last_name":"Pham", "email":"nhutquang23@gmail.com", "phone":"032548784", "is_active": false}`,
			expBody: `{"ID":15,"FirstName":"Quang","LastName":"Pham","Email":"nhutquang23@gmail.com","Phone":"031544284","IsActive":false}
`,
			expCode: http.StatusCreated,
		},
		"error name cannot blank": {
			input:   `{"name": "", "email":"thang12@gmail.com", "phone":"031544284", "role":"ADMIN", "is_active": false}`,
			expCode: http.StatusBadRequest,
			expErr:  errNameCannotBeBlank,
		},
	}

	tcsMockUserServ := map[string]struct {
		result service.UserResponse
		err    error
	}{
		"success": {
			result: service.UserResponse{
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
			mockUserSv := new(MockUserService)
			userHandler := UserHandler{mockUserSv}
			mockUserSv.On("CreateUser", mock.Anything, mock.AnythingOfType("service.CreateUserInput")).Return(
				tcsMockUserServ[desc].result, tcsMockUserServ[desc].err)

			handler := http.HandlerFunc(userHandler.CreateUser)
			handler.ServeHTTP(res, req)

			// test cases
			if tc.expErr != nil {
				require.Equal(t, res.Code, tc.expCode)

				// TODO: compare response body with expErr?
				//require.Equal(t, res.Body.String(), tc.expErr)
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
		expResult service.CreateUserInput
		expErr    error
	}{
		"success": {
			input: UserRequest{
				Email:    "nhutquang23@gmail.com",
				Phone:    "02312545678",
				IsActive: true,
			},
			expResult: service.CreateUserInput{
				Email:    "nhutquang23@gmail.com",
				Phone:    "0343450044",
				IsActive: true,
			},
		},
		"nameCannotBeBlank": {
			input: UserRequest{
				Email:    "dcthang@gmail.com",
				Phone:    "0343450044",
				IsActive: true,
			},
			expErr: errNameCannotBeBlank,
		},
		"email cannot be blank": {
			input: UserRequest{
				Phone:    "0343450044",
				IsActive: true,
			},
			expErr: errEmailCannotBeBlank,
		},
		"invalidEmail": {
			input: UserRequest{
				Email:    "Thang12344email",
				Phone:    "0343450044",
				IsActive: true,
			},
			expErr: errInvalidEmail,
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
