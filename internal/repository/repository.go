package repository

import (
	"context"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/models"
)

type UsersRepo interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)

	GetUserByID(ctx context.Context, id int) (models.User, error)

	GetUserByEmail(ctx context.Context, email string) (models.User, error)

	GetUserIDsByEmail(ctx context.Context, email []string) ([]int, error)

	GetEmailsByIDs(ctx context.Context, ids []int) ([]string, error)

	GetUsers(ctx context.Context) (models.UserSlice, error)

	CheckEmailIsExist(ctx context.Context, email string) (bool, error)
}

type FriendsRepo interface {
	CreateUserFriend(ctx context.Context, input models.UserFriend) error

	CheckFriendRelationExist(ctx context.Context, requesterID int, targetID int) (bool, error)

	CheckSubscriptionRelationExist(ctx context.Context, requesterID int, targetID int) (bool, error)

	CheckBlockRelationExist(ctx context.Context, requesterID int, targetID int) (bool, error)

	GetRelationIDs(ctx context.Context, requesterID int, relationType int) ([]int, error)

	GetRequesterIDFriends(ctx context.Context, requesterID int, relationType int) ([]int, error)

	DeleteRelation(ctx context.Context, requesterID int, targetID int, relationType int) error
}
