package user

import (
	"context"
	"testing"

	"github.com/friendsofgo/errors"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/db"
	"github.com/stretchr/testify/require"
)

var dbURL = "postgresql://root:secret@localhost:5432/friends_management?sslmode=disable"
var dbDriver = "postgres"

func TestRepository_CreateUser(t *testing.T) {
	tcs := map[string]struct {
		input     models.User
		expResult models.User
		expErr    error
	}{
		"success": {
			input: models.User{
				Email: "nhutquang23@gmail.com",
			},
			expResult: models.User{
				Email: "nhutquang23@gmail.com",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(dbURL)
			require.NoError(t, err)
			defer dbConn.Close()

			userRepo := NewUserRepository(dbConn)
			res, err := userRepo.CreateUser(ctx, tc.input)

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				tc.expResult.ID = res.ID
				tc.expResult.CreatedAt = res.CreatedAt
				tc.expResult.UpdatedAt = res.UpdatedAt
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}

func TestRepository_GetUserByEmail(t *testing.T) {
	tcs := map[string]struct {
		input     string
		expResult models.User
		expErr    error
	}{
		"success": {
			input: "andy@example.com",
			expResult: models.User{
				Email: "andy@example.com",
			},
		},
		"email is empty": {
			input:  "",
			expErr: errors.New("Email cannot be empty"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()

			// Connect DB test
			dbConn, err := db.ConnectDB(dbURL)
			require.NoError(t, err)
			defer dbConn.Close()
			//defer dbConn.Exec("DELETE FROM public.users;")

			userRepo := NewUserRepository(dbConn)
			res, err := userRepo.GetUserByEmail(ctx, tc.input)

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				tc.expResult.ID = res.ID
				tc.expResult.CreatedAt = res.CreatedAt
				tc.expResult.UpdatedAt = res.UpdatedAt
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}

func TestRepository_GetUserByID(t *testing.T) {
	tcs := map[string]struct {
		input     int
		expResult models.User
		expErr    error
	}{
		"success": {
			input: 3,
			expResult: models.User{
				Email: "common@example.com",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(dbURL)
			require.NoError(t, err)
			defer dbConn.Close()

			userRepo := NewUserRepository(dbConn)
			res, err := userRepo.GetUserByID(ctx, tc.input)

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				tc.expResult.ID = res.ID
				tc.expResult.CreatedAt = res.CreatedAt
				tc.expResult.UpdatedAt = res.UpdatedAt
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}

func TestRepository_GetUserIDsByEmail(t *testing.T) {
	tcs := map[string]struct {
		input     []string
		expResult []int
		expErr    error
	}{
		"success": {
			input:     []string{"andy@example.com"},
			expResult: []int{1},
		},
		"emails is null": {
			input:  []string{},
			expErr: errors.New("Slice of emails cannot empty"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(dbURL)
			require.NoError(t, err)
			defer dbConn.Close()

			userRepo := NewUserRepository(dbConn)
			res, err := userRepo.GetUserIDsByEmail(ctx, tc.input)

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				tc.expResult = res
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}

func TestRepository_GetListUser(t *testing.T) {
	tcs := map[string]struct {
		expResult models.UserSlice
		expErr    error
	}{
		"success": {
			expResult: models.UserSlice{
				{
					Email: "andy@example.com",
					Name:  "andy",
				},
				{
					Email: "john@example.com",
					Name:  "john",
				},
				{
					Email: "common@example.com",
					Name:  "common",
				},
				{
					Email: "lisa@example.com",
					Name:  "lisa",
				},
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()

			// Connect DB test
			dbConn, err := db.ConnectDB(dbURL)
			require.NoError(t, err)
			defer dbConn.Close()

			userRepo := NewUserRepository(dbConn)
			res, err := userRepo.GetUsers(ctx)

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				tc.expResult = res
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}
