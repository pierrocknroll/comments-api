package services

import (
	"comments-api/internal/core/domains"
	"comments-api/internal/core/ports"
	"comments-api/internal/errors"
	log "github.com/sirupsen/logrus"
)

type service struct {
	repository ports.CommentsRepository
}

func NewCommentService(repository ports.CommentsRepository) *service {
	return &service{
		repository: repository,
	}
}

func (service *service) GetCommentsByTarget(target domains.Target) ([]domains.Comment, error) {
	comments, err := service.repository.GetCommentsByTarget(target)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "target": target}).
			Error("Cannot find comments")
		return nil, err
	}
	return comments, nil
}
func (service *service) StoreComment(target domains.Target, userId domains.Author, message string) (*domains.Comment, error) {
	comment, err := service.repository.StoreComment(target, userId, message)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "target": target}).
			Error("Cannot store comment")
		return nil, err
	}

	return comment, nil

}
func (service *service) UpdateCommentMessage(commentId int, userId domains.Author, message string) (*domains.Comment, error) {
	owns, err := service.repository.IsUserOwnsComment(userId, commentId)
	if err != nil || owns == false {
		return nil, errors.ErrUserIsNotOwnerOfComment
	}

	comment, err := service.repository.UpdateCommentMessage(commentId, message)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "commentId": commentId}).
			Error("Cannot update comment")
		return nil, err
	}

	return comment, nil
}

func (service *service) DeleteComment(commentId int, userId domains.Author) error {
	owns, err := service.repository.IsUserOwnsComment(userId, commentId)
	if err != nil || owns == false {
		return errors.ErrUserIsNotOwnerOfComment
	}

	err = service.repository.DeleteComment(commentId)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "commentId": commentId}).
			Error("Cannot delete comment")
	}
	return err
}
