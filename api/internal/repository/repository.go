package repository

import (
	"context"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/api/internal/repository/orm/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user dbmodels.User) (dbmodels.User, error)

	GetUserByID(ctx context.Context, id int) (dbmodels.User, error)

	GetUserByEmail(ctx context.Context, email string) (dbmodels.User, error)

	GetUserIDsByEmail(ctx context.Context, email []string) ([]int, error)

	GetListEmailByIDs(ctx context.Context, ids []int) ([]string, error)

	GetAllUser(ctx context.Context) ([]string, error)
}

type RelationsRepo interface {
	CreateRelation(ctx context.Context, input dbmodels.Relation) (bool, error)

	IsRelationExist(ctx context.Context, requesterId int, addresseeId int, relationType int) (bool, error)

	GetRelationIDsOfUser(ctx context.Context, requesterId int, relationType int) ([]int, error)

	GetRequesterIDRelation(ctx context.Context, requesterId int, relationType int) ([]int, error)

	DeleteRelation(ctx context.Context, requesterId int, addresseeId int, relationType int) error
}
