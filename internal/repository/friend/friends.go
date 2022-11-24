package friend

import (
	"context"
	"database/sql"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type friendsRepo struct {
	db *sql.DB
}

func NewFriendsRepository(db *sql.DB) repository.FriendsRepo {
	return friendsRepo{
		db: db,
	}
}

// CheckFriendRelationExist func check if 2 emails are friend
func (repo friendsRepo) CheckFriendRelationExist(ctx context.Context, requesterID int, targetID int) (bool, error) {
	return models.UserFriends(
		qm.Where(
			models.UserFriendColumns.RequesterID+" = ? and "+
				models.UserFriendColumns.TargetID+" = ? and "+
				models.UserFriendColumns.RelationType+"= ? ", requesterID, targetID, utils.FriendRelation,
		),
		qm.Or2(
			qm.Where(
				models.UserFriendColumns.RequesterID+" = ? and "+
					models.UserFriendColumns.TargetID+" = ? and "+
					models.UserFriendColumns.RelationType+" = ? ", targetID, requesterID, utils.FriendRelation),
		),
	).Exists(ctx, repo.db)
}

// CheckSubscriptionRelationExist func check if 2 emails are subscript
func (repo friendsRepo) CheckSubscriptionRelationExist(ctx context.Context, requesterID int, targetID int) (bool, error) {
	return models.UserFriends(
		qm.Where(
			models.UserFriendColumns.RequesterID+" = ? and "+
				models.UserFriendColumns.TargetID+" = ? and "+
				models.UserFriendColumns.RelationType+" = ? ", requesterID, targetID, utils.Subscribe,
		),
	).Exists(ctx, repo.db)
}

// CheckBlockRelationExist func check if 2 emails are block
func (repo friendsRepo) CheckBlockRelationExist(ctx context.Context, requesterID int, targetID int) (bool, error) {
	return models.UserFriends(
		qm.Where(
			models.UserFriendColumns.RequesterID+" = ? and "+
				models.UserFriendColumns.TargetID+" = ? and "+
				models.UserFriendColumns.RelationType+"= ? ", requesterID, targetID, utils.Blocked,
		),
	).Exists(ctx, repo.db)
}

// CreateUserFriend create new user friends
func (repo friendsRepo) CreateUserFriend(ctx context.Context, input models.UserFriend) error {
	var relation = models.UserFriend{
		RequesterID:  input.RequesterID,
		TargetID:     input.TargetID,
		RelationType: input.RelationType,
	}

	return relation.Insert(ctx, repo.db, boil.Infer())
}

// GetRelationIDs function get IDs of relations
func (repo friendsRepo) GetRelationIDs(ctx context.Context, requesterID int, relationType int) ([]int, error) {
	var qms []qm.QueryMod
	qms = append(qms, qm.Where(models.UserFriendColumns.RequesterID+" = ? OR "+models.UserFriendColumns.TargetID+" = ?", requesterID, requesterID))
	qms = append(qms, qm.Where(models.UserFriendColumns.RelationType+" = ?", relationType))

	friendRelations, err := models.UserFriends(qms...).All(ctx, repo.db)
	if err != nil {
		return nil, err
	}

	userIDs := make([]int, len(friendRelations))
	for idx, o := range friendRelations {
		if o.RequesterID != requesterID {
			userIDs[idx] = o.RequesterID
		} else {
			userIDs[idx] = o.TargetID
		}
	}

	return userIDs, nil
}

// DeleteRelation function delete relation between 2 ids
func (repo friendsRepo) DeleteRelation(ctx context.Context, requesterID int, targetID int, relationType int) error {
	_, err := models.UserFriends(
		models.UserFriendWhere.RelationType.EQ(null.IntFrom(relationType)),
		models.UserFriendWhere.RequesterID.EQ(requesterID),
		models.UserFriendWhere.TargetID.EQ(targetID),
	).DeleteAll(ctx, repo.db)

	return err
}

// GetRequesterIDFriends function get requestID from relation
func (repo friendsRepo) GetRequesterIDFriends(ctx context.Context, requesterID int, relationType int) ([]int, error) {
	var qms []qm.QueryMod
	qms = append(qms, qm.Where(models.UserFriendColumns.RequesterID+" = ? ", requesterID))
	qms = append(qms, qm.Where(models.UserFriendColumns.RelationType+" = ?", relationType))

	friendRelations, err := models.UserFriends(qms...).All(ctx, repo.db)
	if err != nil {
		return nil, err
	}

	userIDs := make([]int, len(friendRelations))
	for idx, o := range friendRelations {
		if o.RequesterID != requesterID {
			userIDs[idx] = o.RequesterID
		} else {
			userIDs[idx] = o.TargetID
		}
	}

	return userIDs, nil
}
