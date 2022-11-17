package friends

import "errors"

var (
	ErrRequestEmailInvalid        = errors.New("request email from request is invalid")
	ErrTargetEmailInvalid         = errors.New("target email from request is invalid")
	ErrTwoEmailAlreadyFriend      = errors.New("two emails are friend")
	ErrTwoEmailInvalidCreateSub   = errors.New("two emails are invalid create subscription")
	ErrTwoEmailInvalidCreateBlock = errors.New("two emails are invalid create block")
	ErrRelationIsExists           = errors.New("relation is exists")
)
