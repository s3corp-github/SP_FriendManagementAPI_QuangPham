package rest

import (
	"database/sql"

	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/service/friends"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/service/users"
)

// Handler presents for handler
type Handler struct {
	userService   users.IService
	friendService friends.IService
}

// NewHandler create users handler contain UsersHandler
func NewHandler(db *sql.DB) Handler {
	return Handler{
		userService:   users.NewUserService(db),
		friendService: friends.NewService(db),
	}
}
