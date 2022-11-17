package users

import "github.com/friendsofgo/errors"

var (
	ErrEmailExist = errors.New("Email already in use ")
	ErrCreateUser = errors.New("Error when create user please contact with admin")
)
