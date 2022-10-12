package repository

import (
	"context"
	models "github.com/quangpham789/golang-assessment/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)

	GetUserByID(ctx context.Context, id int) (models.User, error)

	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

type RelationsRepo interface {
	CreateRelation(ctx context.Context, input models.Relation) (bool, error)

	GetRelationByIdsAndType(ctx context.Context, requesterId int, addresseeId int, relationType int) (models.Relation, error)

	GetAllRelationFriendOfUser(ctx context.Context, requesterId int) (models.RelationSlice, error)

	GetCommonFriend(ctx context.Context, firstRequesterId int, secondRequesterId int) (models.RelationSlice, error)
}

type MockUserRepo interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)

	GetUserByID(ctx context.Context, id int) (*models.User, error)

	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}
