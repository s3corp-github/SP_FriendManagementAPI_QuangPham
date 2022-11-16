package friends

import (
	"context"
	"database/sql"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository/friend"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/repository/user"
)

type IService interface {
	CreateFriend(ctx context.Context, input CreateRelationsInput) error
	CreateSubscription(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error)
	CreateBlock(ctx context.Context, input CreateRelationsInput) (CreateRelationsResponse, error)
	GetAllFriends(ctx context.Context, input GetAllFriendsInput) ([]string, error)
	GetCommonFriends(ctx context.Context, input CommonFriendsInput) ([]string, error)
	GetEmailReceive(ctx context.Context, input EmailReceiveInput) (EmailReceiveResponse, error)
}

// FriendService type contain repository needed
type FriendService struct {
	friendRepo repository.FriendsRepo
	userRepo   repository.UsersRepo
}

// NewService create new friends service
func NewService(db *sql.DB) IService {
	return FriendService{
		friendRepo: friend.NewRelationsRepository(db),
		userRepo:   user.NewUserRepository(db),
	}
}
