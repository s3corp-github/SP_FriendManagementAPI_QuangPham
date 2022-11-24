package friends

import "errors"

var (
	ErrTwoEmailInvalidCreateSub   = errors.New("two emails are invalid create subscription")
	ErrTwoEmailInvalidCreateBlock = errors.New("two emails are invalid create block")
	ErrRelationIsExists           = errors.New("relation is exists")
)
