package friend

import (
	"context"
	"os"
	"testing"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/db"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
	"github.com/volatiletech/null/v8"

	"github.com/friendsofgo/errors"
	"github.com/stretchr/testify/require"
)

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
		"case missing target id": {
			input: models.UserFriend{
				RequesterID:  1,
				RelationType: null.IntFrom(utils.FriendRelation),
			},
			expResult: false,
			expErr:    errors.New(`models: unable to insert into user_friends: pq: insert or update on table "user_friends" violates foreign key constraint "user_friends_target_id_fkey"`),
		},
		"case missing requester id": {
			input: models.UserFriend{
				TargetID:     2,
				RelationType: null.IntFrom(utils.FriendRelation),
			},
			expResult: false,
			expErr:    errors.New(`models: unable to insert into user_friends: pq: insert or update on table "user_friends" violates foreign key constraint "user_friends_requester_id_fkey"`),
		},
		"case missing requester id and target id": {
			input: models.UserFriend{
				RelationType: null.IntFrom(utils.FriendRelation),
			},
			expResult: false,
			expErr:    errors.New(`models: unable to insert into user_friends: pq: insert or update on table "user_friends" violates foreign key constraint "user_friends_requester_id_fkey"`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(os.Getenv("DB_SOURCE"))
			require.NoError(t, err)
			defer dbConn.Close()

			friendshipRepo := NewFriendsRepository(dbConn)

			var res bool
			err = friendshipRepo.CreateUserFriend(ctx, tc.input)
			if err != nil {
				res = false
			}

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

func TestRepository_GetRelationIDs(t *testing.T) {
	tcs := map[string]struct {
		requesterID  int
		relationType int
		expResult    []int
		expErr       error
	}{
		"success": {
			requesterID:  1,
			relationType: 1,
			expResult:    []int{2, 3, 2},
		},
		"missing requesterID": {
			relationType: 1,
			expErr:       errors.New(`requesterID cannot be null`),
		},
		"missing relationType": {
			requesterID: 1,
			expErr:      errors.New(`relationType cannot be null`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(os.Getenv("DB_SOURCE"))
			require.NoError(t, err)
			defer dbConn.Close()

			friendshipRepo := NewFriendsRepository(dbConn)
			res, err := friendshipRepo.GetRelationIDs(ctx, tc.requesterID, tc.relationType)

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
		requesterID  int
		relationType int
		expResult    []int
		expErr       error
	}{
		"success": {
			requesterID:  4,
			relationType: 2,
			expResult:    []int{2},
		},
		"missing requesterID": {
			relationType: 1,
			expErr:       errors.New(`requesterID cannot be null`),
		},
		"missing relationType": {
			requesterID: 1,
			expErr:      errors.New(`relationType cannot be null`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(os.Getenv("DB_SOURCE"))
			require.NoError(t, err)
			defer dbConn.Close()

			friendshipRepo := NewFriendsRepository(dbConn)
			res, err := friendshipRepo.GetRequesterIDFriends(ctx, tc.requesterID, tc.relationType)

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
		requesterID  int
		relationType int
		targetID     int
		expResult    error
		expErr       error
	}{
		"success": {
			requesterID:  4,
			targetID:     2,
			relationType: 2,
			expResult:    nil,
		},
		"missing relationType": {
			requesterID: 4,
			targetID:    2,
			expErr:      errors.New(`relationType cannot be null`),
		},
		"missing requesterID": {
			targetID:     2,
			relationType: 1,
			expErr:       errors.New(`requesterID cannot be null`),
		},
		"missing targetID": {
			requesterID:  4,
			relationType: 1,
			expErr:       errors.New(`targetID cannot be null`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(os.Getenv("DB_SOURCE"))
			require.NoError(t, err)
			defer dbConn.Close()

			friendshipRepo := NewFriendsRepository(dbConn)
			err = friendshipRepo.DeleteRelation(ctx, tc.requesterID, tc.targetID, tc.relationType)

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
		requesterID  int
		relationType int
		targetID     int
		expResult    bool
		expErr       error
	}{
		"success friends": {
			requesterID:  2,
			targetID:     3,
			relationType: 2,
			expResult:    true,
		},
		"success sub friends": {
			requesterID:  1,
			targetID:     2,
			relationType: 2,
			expResult:    true,
		},
		"success block friends": {
			requesterID:  1,
			targetID:     2,
			relationType: 3,
			expResult:    true,
		},
		"missing relationType": {
			requesterID: 4,
			targetID:    2,
			expErr:      errors.New(`relationType cannot be null`),
		},
		"missing requesterID": {
			targetID:     2,
			relationType: 1,
			expErr:       errors.New(`requesterID cannot be null`),
		},
		"missing targetID": {
			requesterID:  4,
			relationType: 1,
			expErr:       errors.New(`targetID cannot be null`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			// Connect DB test
			dbConn, err := db.ConnectDB(os.Getenv("DB_SOURCE"))
			require.NoError(t, err)
			defer dbConn.Close()

			friendshipRepo := NewFriendsRepository(dbConn)
			result, err := friendshipRepo.IsFriendRelationExist(ctx, tc.requesterID, tc.targetID)

			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, result)
			}
		})
	}

}
