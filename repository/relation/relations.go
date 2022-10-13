package relation

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	models "github.com/quangpham789/golang-assessment/models"
	"github.com/quangpham789/golang-assessment/repository"
	"github.com/quangpham789/golang-assessment/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
)

type relationsRepository struct {
	connection *sql.DB
}

func NewRelationsRepository(db *sql.DB) repository.RelationsRepo {
	return relationsRepository{
		connection: db,
	}
}

func (repo relationsRepository) IsRelationExist(ctx context.Context, requesterId int, addresseeId int, relationType int) (bool, error) {
	if requesterId == 0 {
		return false, errors.New("requesterId cannot be null")
	}
	if addresseeId == 0 {
		return false, errors.New("addresseeId cannot be null")
	}
	if relationType == 0 {
		return false, errors.New("relationType cannot be null")
	}
	//relation is two-way relationship
	if relationType == utils.FriendRelation {

		isExists, err := models.Relations(
			models.RelationWhere.RelationType.EQ(null.IntFrom(relationType)),
			models.RelationWhere.RequesterID.EQ(addresseeId),
			models.RelationWhere.AddresseeID.EQ(requesterId),
		).Exists(ctx, repo.connection)

		if err != nil {
			return false, err
		}

		if isExists == true {
			return true, nil
		}
	}
	isExists, err := models.Relations(
		models.RelationWhere.RelationType.EQ(null.IntFrom(relationType)),
		models.RelationWhere.RequesterID.EQ(requesterId),
		models.RelationWhere.AddresseeID.EQ(addresseeId),
	).Exists(ctx, repo.connection)

	if err != nil {
		return false, err
	}

	return isExists, nil
}

func (repo relationsRepository) CreateRelation(ctx context.Context, input models.Relation) (bool, error) {
	var relation = models.Relation{}
	relation.RequesterID = input.RequesterID
	relation.AddresseeID = input.AddresseeID
	relation.AddresseeEmail = input.AddresseeEmail
	relation.RequesterEmail = input.RequesterEmail
	relation.RelationType = input.RelationType
	if err := relation.Insert(ctx, repo.connection, boil.Infer()); err != nil {
		return false, err
	}
	return true, nil
}

func (repo relationsRepository) GetRelationIDsOfUser(ctx context.Context, requesterId int, relationType int) ([]int, error) {
	if requesterId == 0 {
		return nil, errors.New("requesterId cannot be null")
	}
	if relationType == 0 {
		return nil, errors.New("relationType cannot be null")
	}
	var qms []qm.QueryMod
	qms = append(qms, qm.Where(models.RelationColumns.RequesterID+" = ? OR "+models.RelationColumns.AddresseeID+" = ?", requesterId, requesterId))
	qms = append(qms, qm.Where(models.RelationColumns.RelationType+" = ?", relationType))

	friendRelations, err := models.Relations(qms...).All(ctx, repo.connection)
	if err != nil {
		log.Println("GetRelationIDsOfUser error when get all friends of user", err)
		return nil, err
	}

	userIDs := make([]int, len(friendRelations))

	for idx, o := range friendRelations {
		if o.RequesterID != requesterId {
			userIDs[idx] = o.RequesterID
		} else {
			userIDs[idx] = o.AddresseeID
		}
	}

	return userIDs, nil
}
func (repo relationsRepository) DeleteRelation(ctx context.Context, requesterId int, addresseeId int, relationType int) error {
	if requesterId == 0 {
		return errors.New("requesterId cannot be null")
	}
	if addresseeId == 0 {
		return errors.New("addresseeId cannot be null")
	}
	if relationType == 0 {
		return errors.New("relationType cannot be null")
	}
	_, err := models.Relations(
		models.RelationWhere.RelationType.EQ(null.IntFrom(relationType)),
		models.RelationWhere.RequesterID.EQ(requesterId),
		models.RelationWhere.AddresseeID.EQ(addresseeId),
	).DeleteAll(ctx, repo.connection)

	if err != nil {
		return err
	}

	return nil
}

func (repo relationsRepository) GetRequesterIDRelation(ctx context.Context, requesterId int, relationType int) ([]int, error) {
	if requesterId == 0 {
		return nil, errors.New("requesterId cannot be null")
	}
	if relationType == 0 {
		return nil, errors.New("relationType cannot be null")
	}
	var qms []qm.QueryMod
	qms = append(qms, qm.Where(models.RelationColumns.RequesterID+" = ? ", requesterId))
	qms = append(qms, qm.Where(models.RelationColumns.RelationType+" = ?", relationType))

	friendRelations, err := models.Relations(qms...).All(ctx, repo.connection)
	if err != nil {
		log.Println("GetRequesterSubRelation error when get all friends of user", err)
		return nil, err
	}

	userIDs := make([]int, len(friendRelations))

	for idx, o := range friendRelations {
		if o.RequesterID != requesterId {
			userIDs[idx] = o.RequesterID
		} else {
			userIDs[idx] = o.AddresseeID
		}
	}

	return userIDs, nil
}
