package friend

import (
	"context"
	"database/sql"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type friendsRepo struct {
	connection *sql.DB
}

func NewRelationsRepository(db *sql.DB) repository.FriendsRepo {
	return friendsRepo{
		connection: db,
	}
}

func (repo friendsRepo) IsRelationExist(ctx context.Context, requesterId int, addresseeId int, relationType int) (bool, error) {
	if relationType == utils.FriendRelation {
		isExists, err := models.UserFriends(
			models.UserFriendWhere.RelationType.EQ(null.IntFrom(relationType)),
			models.UserFriendWhere.RequesterID.EQ(addresseeId),
			models.UserFriendWhere.TargetID.EQ(requesterId),
		).Exists(ctx, repo.connection)
		if err != nil {
			return false, err
		}

		if isExists == true {
			return true, nil
		}
	}

	isExists, err := models.UserFriends(
		models.UserFriendWhere.RelationType.EQ(null.IntFrom(relationType)),
		models.UserFriendWhere.RequesterID.EQ(requesterId),
		models.UserFriendWhere.TargetID.EQ(addresseeId),
	).Exists(ctx, repo.connection)
	if err != nil {
		return false, err
	}

	return isExists, nil
}

// CreateUserFriend create new user friends
func (repo friendsRepo) CreateUserFriend(ctx context.Context, input models.UserFriend) error {
	var relation = models.UserFriend{
		RequesterID:  input.RequesterID,
		TargetID:     input.TargetID,
		RelationType: input.RelationType,
	}

	return relation.Insert(ctx, repo.connection, boil.Infer())
}

func (repo friendsRepo) GetRelationIDsOfUser(ctx context.Context, requesterId int, relationType int) ([]int, error) {
	var qms []qm.QueryMod
	qms = append(qms, qm.Where(models.UserFriendColumns.RequesterID+" = ? OR "+models.UserFriendColumns.TargetID+" = ?", requesterId, requesterId))
	qms = append(qms, qm.Where(models.UserFriendColumns.RelationType+" = ?", relationType))

	friendRelations, err := models.UserFriends(qms...).All(ctx, repo.connection)
	if err != nil {
		return nil, err
	}

	userIDs := make([]int, len(friendRelations))

	for idx, o := range friendRelations {
		if o.RequesterID != requesterId {
			userIDs[idx] = o.RequesterID
		} else {
			userIDs[idx] = o.TargetID
		}
	}

	return userIDs, nil
}
func (repo friendsRepo) DeleteRelation(ctx context.Context, requesterId int, addresseeId int, relationType int) error {

	_, err := models.UserFriends(
		models.UserFriendWhere.RelationType.EQ(null.IntFrom(relationType)),
		models.UserFriendWhere.RequesterID.EQ(requesterId),
		models.UserFriendWhere.TargetID.EQ(addresseeId),
	).DeleteAll(ctx, repo.connection)

	return err
}

func (repo friendsRepo) GetRequesterIDRelation(ctx context.Context, requesterId int, relationType int) ([]int, error) {

	var qms []qm.QueryMod
	qms = append(qms, qm.Where(models.UserFriendColumns.RequesterID+" = ? ", requesterId))
	qms = append(qms, qm.Where(models.UserFriendColumns.RelationType+" = ?", relationType))

	friendRelations, err := models.UserFriends(qms...).All(ctx, repo.connection)
	if err != nil {
		return nil, err
	}

	userIDs := make([]int, len(friendRelations))

	for idx, o := range friendRelations {
		if o.RequesterID != requesterId {
			userIDs[idx] = o.RequesterID
		} else {
			userIDs[idx] = o.TargetID
		}
	}

	return userIDs, nil
}
