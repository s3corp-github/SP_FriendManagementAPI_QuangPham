package user

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/quangpham789/golang-assessment/api/internal/config/db"
	models "github.com/quangpham789/golang-assessment/api/internal/repository/orm/models"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"testing"
)

var dbURL = "postgresql://root:secret@localhost:5432/friends_management?sslmode=disable"

func TestRepository_CreateUser(t *testing.T) {
	tcs := map[string]struct {
		input     models.User
		expResult models.User
		expErr    error
	}{
		"success": {
			input: models.User{
				Email:    "nhutquang23@gmail.com",
				Phone:    null.StringFrom("0343450044"),
				IsActive: null.BoolFrom(true),
			},
			expResult: models.User{
				Email:    "nhutquang23@gmail.com",
				Phone:    null.StringFrom("0343450044"),
				IsActive: null.BoolFrom(true),
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
			//defer dbConn.Exec("DELETE FROM public.users;")

			// TODO: Load DB user test sql

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
				Email:    "andy@example.com",
				Phone:    null.StringFrom("123456"),
				IsActive: null.BoolFrom(true),
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
				Email:    "common@example.com",
				Phone:    null.StringFrom("123456"),
				IsActive: null.BoolFrom(true),
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
			//defer dbConn.Exec("DELETE FROM public.users;")

			// TODO: Load DB user test sql

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
			//defer dbConn.Exec("DELETE FROM public.users;")

			// TODO: Load DB user test sql

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
