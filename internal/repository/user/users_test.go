package user

import (
	"context"
	"os"
	"testing"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/db"
	"github.com/stretchr/testify/require"
)

func TestRepository_CreateUser(t *testing.T) {
	tcs := map[string]struct {
		input     models.User
		expResult models.User
		expErr    error
	}{
		"success": {
			input: models.User{
				Name:  "john",
				Email: "john@example.com",
			},
			expResult: models.User{
				Name:  "john",
				Email: "john@example.com",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(os.Getenv("DB_SOURCE"))
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
			input: "john@example.com",
			expResult: models.User{
				Name:  "john",
				Email: "john@example.com",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()

			// Connect DB test
			dbConn, err := db.ConnectDB(os.Getenv("DB_SOURCE"))
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
			input: 1,
			expResult: models.User{
				Name:  "john",
				Email: "john@example.com",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(os.Getenv("DB_SOURCE"))
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
			input:     []string{"john@example.com"},
			expResult: []int{1},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(os.Getenv("DB_SOURCE"))
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
			dbConn, err := db.ConnectDB(os.Getenv("DB_SOURCE"))
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
