package repository

import (
	"context"
	models "github.com/quangpham789/golang-assessment/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)

	GetUserByID(ctx context.Context, id int) (models.User, error)

	GetUserByEmail(ctx context.Context, email string) (models.User, error)

	GetUserIDsByEmail(ctx context.Context, email []string) ([]int, error)

	GetListEmailByIDs(ctx context.Context, ids []int) ([]string, error)
}

type RelationsRepo interface {
	CreateRelation(ctx context.Context, input models.Relation) (bool, error)

	IsRelationExist(ctx context.Context, requesterId int, addresseeId int, relationType int) (bool, error)

	GetRelationIDsOfUser(ctx context.Context, requesterId int, relationType int) ([]int, error)

	GetRequesterIDRelation(ctx context.Context, requesterId int, relationType int) ([]int, error)

	DeleteRelation(ctx context.Context, requesterId int, addresseeId int, relationType int) error
}
