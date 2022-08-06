package ports

import (
	"comments-api/internal/core/domains"
)

type CommentsRepository interface {
	IsUserOwnsComment(userId domains.Author, commentId int) (bool, error)
	GetCommentsByTarget(target domains.Target) ([]domains.Comment, error)
	StoreComment(target domains.Target, userId domains.Author, message string) (*domains.Comment, error)
	UpdateCommentMessage(commentId int, message string) (*domains.Comment, error)
	DeleteComment(commentId int) error
}
