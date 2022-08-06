package ports

import (
	"comments-api/internal/core/domains"
)

type CommentsService interface {
	GetCommentsByTarget(target domains.Target) ([]domains.Comment, error)
	StoreComment(target domains.Target, userId domains.Author, message string) (*domains.Comment, error)
	UpdateCommentMessage(commentId int, userId domains.Author, message string) (*domains.Comment, error)
	DeleteComment(commentId int, userId domains.Author) error
}
