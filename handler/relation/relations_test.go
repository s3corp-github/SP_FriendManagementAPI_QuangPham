package relation

import (
	"github.com/quangpham789/golang-assessment/service/user/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_CreateFriendsRelation(t *testing.T) {
	tcs := map[string]struct {
		params  map[string]string
		expCode int
		expBody []string
		expErr  error
	}{
		"success": {
			input:   `{"friends":["andy@example.com", "common@example.com"]}`,
			expBody: []string{"andy@example.com", "common@example.com"},
			expCode: http.StatusCreated,
		},
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

				// TODO: compare response body with expErr?
				//require.Equal(t, res.Body.String(), tc.expErr)
			} else {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, res.Body.String(), tc.expBody)
			}

		})
	}
}
