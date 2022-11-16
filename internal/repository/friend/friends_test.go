package friend

import (
	"context"
	"testing"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/db"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/models"
	"github.com/volatiletech/null/v8"

	"github.com/friendsofgo/errors"
	"github.com/stretchr/testify/require"
)

var dbURL = "postgresql://root:secret@localhost:5432/friends_management?sslmode=disable"
var dbDriver = "postgres"

func TestRepository_CreateFriendRelation(t *testing.T) {
	tcs := map[string]struct {
		input     models.UserFriend
		expResult bool
		expErr    error
	}{
		"success": {
			input: models.UserFriend{
				RequesterID:  1,
				TargetID:     2,
				RelationType: null.IntFrom(utils.FriendRelation),
			},
			expResult: true,
		},
		"case missing addressee id": {
			input: models.UserFriend{
				RequesterID:    1,
				RequesterEmail: "andy@example.com",
				AddresseeEmail: "john@example.com",
				RelationType:   utils.FriendRelation,
			},
			expResult: false,
			expErr:    errors.New(`dbmodels: unable to insert into relations: pq: insert or update on table "relations" violates foreign key constraint "relations_addressee_id_fkey"`),
		},
		"case missing requester id": {
			input: models.Relation{
				AddresseeID:    2,
				RequesterEmail: "andy@example.com",
				AddresseeEmail: "john@example.com",
				RelationType:   utils.FriendRelation,
			},
			expResult: false,
			expErr:    errors.New(`dbmodels: unable to insert into relations: pq: insert or update on table "relations" violates foreign key constraint "relations_requester_id_fkey"`),
		},
		"case missing requester id and addressee id": {
			input: models.Relation{
				RequesterEmail: "andy@example.com",
				AddresseeEmail: "john@example.com",
				RelationType:   utils.FriendRelation,
			},
			expResult: false,
			expErr:    errors.New(`dbmodels: unable to insert into relations: pq: insert or update on table "relations" violates foreign key constraint "relations_requester_id_fkey"`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(dbDriver, dbURL)
			require.NoError(t, err)
			defer dbConn.Close()

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

func TestRepository_GetRelationIDsOfUser(t *testing.T) {
	tcs := map[string]struct {
		requesterId  int
		relationType int
		expResult    []int
		expErr       error
	}{
		"success": {
			requesterId:  1,
			relationType: 1,
			expResult:    []int{2, 3, 2},
		},
		"missing requesterId": {
			relationType: 1,
			expErr:       errors.New(`requesterId cannot be null`),
		},
		"missing relationType": {
			requesterId: 1,
			expErr:      errors.New(`relationType cannot be null`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(dbDriver, dbURL)
			require.NoError(t, err)
			defer dbConn.Close()

			friendshipRepo := NewRelationsRepository(dbConn)
			res, err := friendshipRepo.GetRelationIDsOfUser(ctx, tc.requesterId, tc.relationType)

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}

func TestRepository_GetRequesterIDRelation(t *testing.T) {
	tcs := map[string]struct {
		requesterId  int
		relationType int
		expResult    []int
		expErr       error
	}{
		"success": {
			requesterId:  4,
			relationType: 2,
			expResult:    []int{2},
		},
		"missing requesterId": {
			relationType: 1,
			expErr:       errors.New(`requesterId cannot be null`),
		},
		"missing relationType": {
			requesterId: 1,
			expErr:      errors.New(`relationType cannot be null`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(dbDriver, dbURL)
			require.NoError(t, err)
			defer dbConn.Close()

			friendshipRepo := NewRelationsRepository(dbConn)
			res, err := friendshipRepo.GetRequesterIDRelation(ctx, tc.requesterId, tc.relationType)

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, res)
			}
		})
	}

}

func TestRepository_DeleteRelation(t *testing.T) {
	tcs := map[string]struct {
		requesterId  int
		relationType int
		addresseeId  int
		expResult    error
		expErr       error
	}{
		"success": {
			requesterId:  4,
			addresseeId:  2,
			relationType: 2,
			expResult:    nil,
		},
		"missing relationType": {
			requesterId: 4,
			addresseeId: 2,
			expErr:      errors.New(`relationType cannot be null`),
		},
		"missing requesterId": {
			addresseeId:  2,
			relationType: 1,
			expErr:       errors.New(`requesterId cannot be null`),
		},
		"missing addresseeId": {
			requesterId:  4,
			relationType: 1,
			expErr:       errors.New(`addresseeId cannot be null`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(dbDriver, dbURL)
			require.NoError(t, err)
			defer dbConn.Close()

			friendshipRepo := NewRelationsRepository(dbConn)
			err = friendshipRepo.DeleteRelation(ctx, tc.requesterId, tc.addresseeId, tc.relationType)

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, nil)
			}
		})
	}

}

func TestRepository_IsRelationExist(t *testing.T) {
	tcs := map[string]struct {
		requesterId  int
		relationType int
		addresseeId  int
		expResult    bool
		expErr       error
	}{
		"success friend friends": {
			requesterId:  2,
			addresseeId:  3,
			relationType: 2,
			expResult:    true,
		},
		"success sub friends": {
			requesterId:  1,
			addresseeId:  2,
			relationType: 2,
			expResult:    true,
		},
		"success block friends": {
			requesterId:  1,
			addresseeId:  2,
			relationType: 3,
			expResult:    true,
		},
		"missing relationType": {
			requesterId: 4,
			addresseeId: 2,
			expErr:      errors.New(`relationType cannot be null`),
		},
		"missing requesterId": {
			addresseeId:  2,
			relationType: 1,
			expErr:       errors.New(`requesterId cannot be null`),
		},
		"missing addresseeId": {
			requesterId:  4,
			relationType: 1,
			expErr:       errors.New(`addresseeId cannot be null`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(dbDriver, dbURL)
			require.NoError(t, err)
			defer dbConn.Close()

			friendshipRepo := NewRelationsRepository(dbConn)
			result, err := friendshipRepo.IsRelationExist(ctx, tc.requesterId, tc.addresseeId, tc.relationType)

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, result)
			}
		})
	}

}
