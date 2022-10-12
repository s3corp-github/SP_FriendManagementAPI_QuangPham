package relation

import (
	"context"
	"database/sql"
	models "github.com/quangpham789/golang-assessment/models"
	"github.com/quangpham789/golang-assessment/repository"
	"github.com/quangpham789/golang-assessment/utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
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

func (repo relationsRepository) GetCommonFriend(ctx context.Context, firstRequesterId int, secondRequesterId int) (models.RelationSlice, error) {
	var friendRelations models.RelationSlice
	err := queries.Raw(
		`SELECT r.addressee_email AS addressee_email, r.requester_email AS requester_email
					FROM
					relations r
					INNER JOIN (
						SELECT
							r1.addressee_email AS addressee_email,
							r1.requester_email AS requester_email
						FROM
							relations r1
						WHERE
							r1.requester_id = $1 AND r1.relation_type = 1) AS t1 ON t1.addressee_email = r.addressee_email
				WHERE
					r.requester_id = $2 AND r.relation_type = 1`, secondRequesterId, firstRequesterId).Bind(ctx, repo.connection, &friendRelations)
	if err != nil {
		log.Println("GetAllRelationFriendOfTwoEmail error when get all friends of user", err)
		return nil, err
	}
	return friendRelations, nil

}

func (repo relationsRepository) GetAllRelationFriendOfUser(ctx context.Context, requesterId int) (models.RelationSlice, error) {
	var qms []qm.QueryMod
	//qms = append(qms, qm.Where(models.RelationColumns.RequesterID+" = ? OR "+models.RelationColumns.AddresseeID+" = ?", requesterId, requesterId))
	qms = append(qms, qm.Where(models.RelationColumns.RequesterID+" = ?", requesterId))
	qms = append(qms, qm.Where(models.RelationColumns.RelationType+" = ?", utils.FriendRelation))

	friendRelations, err := models.Relations(qms...).All(ctx, repo.connection)
	if err != nil {
		log.Println("GetAllRelationFriendOfUser error when get all friends of user", err)
		return nil, err
	}
	return friendRelations, nil
}

func (repo relationsRepository) GetRelationByIdsAndType(ctx context.Context, requesterId int, addresseeId int, relationType int) (models.Relation, error) {
	var relationResult models.Relation
	var qms []qm.QueryMod
	qms = append(qms, qm.Where(models.RelationColumns.RequesterID+" = ?", requesterId))
	qms = append(qms, qm.Where(models.RelationColumns.AddresseeID+" = ?", addresseeId))
	qms = append(qms, qm.Where(models.RelationColumns.RelationType+" = ?", relationType))
	if err := models.Relations(qms...).Bind(ctx, repo.connection, &relationResult); err != nil {
		return models.Relation{}, err
	}
	return relationResult, nil
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
