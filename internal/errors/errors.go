package errors

import "errors"

var (
	ErrCommentNotFound         = errors.New("comment not found")
	ErrUserIsNotOwnerOfComment = errors.New("user doesn't own the comment")
)
