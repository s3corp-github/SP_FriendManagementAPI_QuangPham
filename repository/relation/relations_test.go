package relation

import (
	"context"
	models "github.com/quangpham789/golang-assessment/models"
	"github.com/quangpham789/golang-assessment/utils"
	"github.com/quangpham789/golang-assessment/utils/db"
	"github.com/stretchr/testify/require"
	"testing"
)

var dbURL = "postgresql://root:secret@localhost:5432/friends_management?sslmode=disable"

func TestRepository_CreateFriendship(t *testing.T) {
	tcs := map[string]struct {
		input     models.Relation
		expResult bool
		expErr    error
	}{
		"success": {
			input: models.Relation{
				RequesterID:    1,
				AddresseeID:    2,
				RequesterEmail: "andy@gmail.com",
				AddresseeEmail: "john@gmail.com",
				RelationType:   utils.FriendRelation,
			},
			expResult: true,
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

			friendshipRepo := NewRelationsRepository(dbConn)
			res, err := friendshipRepo.CreateRelation(ctx, tc.input)

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

func TestRepository_GetAllFriendRelationOfUser(t *testing.T) {
	tcs := map[string]struct {
		input     int
		expResult models.RelationSlice
		expErr    error
	}{
		"success": {
			input: 1,
			expResult: models.RelationSlice{
				&models.Relation{
					RequesterID:    1,
					AddresseeID:    2,
					RequesterEmail: "andy@gmail.com",
					AddresseeEmail: "john@gmail.com",
					RelationType:   utils.FriendRelation,
				},
			},
		},
		// TODO: "error duplicate email"

		// TODO: "error duplicate primary_key"
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

			friendshipRepo := NewRelationsRepository(dbConn)
			res, err := friendshipRepo.GetAllRelationFriendOfUser(ctx, tc.input)

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

func TestRepository_GetCommonFriendRelation(t *testing.T) {
	tcs := map[string]struct {
		input1    int
		input2    int
		expResult models.RelationSlice
		expErr    error
	}{
		"success": {
			input1: 2,
			input2: 3,
			expResult: models.RelationSlice{
				&models.Relation{
					RequesterID:    1,
					AddresseeID:    2,
					RequesterEmail: "andy@gmail.com",
					AddresseeEmail: "john@gmail.com",
					RelationType:   utils.FriendRelation,
				},
			},
		},
		// TODO: "error duplicate email"

		// TODO: "error duplicate primary_key"
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

			friendshipRepo := NewRelationsRepository(dbConn)
			res, err := friendshipRepo.GetCommonFriend(ctx, tc.input1, tc.input2)

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
