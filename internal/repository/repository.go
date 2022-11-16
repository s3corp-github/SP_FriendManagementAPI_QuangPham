package repository

import (
	"context"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/models"
)

type UsersRepo interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)

	GetUserByID(ctx context.Context, id int) (models.User, error)

	GetUserByEmail(ctx context.Context, email string) (models.User, error)

	GetUserIDsByEmail(ctx context.Context, email []string) ([]int, error)

	GetListEmailByIDs(ctx context.Context, ids []int) ([]string, error)

	GetAllUser(ctx context.Context) ([]string, error)
}

type FriendsRepo interface {
	CreateUserFriend(ctx context.Context, input models.UserFriend) error

	IsRelationExist(ctx context.Context, requesterId int, targetId int, relationType int) (bool, error)

	GetRelationIDsOfUser(ctx context.Context, requesterId int, relationType int) ([]int, error)

	GetRequesterIDRelation(ctx context.Context, requesterId int, relationType int) ([]int, error)

	DeleteRelation(ctx context.Context, requesterId int, addresseeId int, relationType int) error
}
