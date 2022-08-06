package repositories

import (
	"comments-api/internal/core/domains"
	"comments-api/internal/errors"
	log "github.com/sirupsen/logrus"
)

// Follows the CommentsRepository Port
type commentRepository struct {
	commentDbContext *DatabaseContext
}

func NewCommentRepository(commentDbDSN string) (*commentRepository, error) {

	repository := &commentRepository{}

	var err error

	// Initialize database connection
	repository.commentDbContext, err = NewDatabaseContext("comments", commentDbDSN)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err}).
			Error("Can't initialize comments database connection")
		return nil, err
	}

	return repository, nil
}

func (repository *commentRepository) Close() {

	err := repository.commentDbContext.Close()
	if err != nil {
		log.
			WithFields(log.Fields{"error": err}).
			Error("Can't close comments database connection")
	}
}

func (repository *commentRepository) GetCommentsByTarget(target domains.Target) ([]domains.Comment, error) {
	comments := make([]domains.Comment, 0)

	commentsDb, err := repository.commentDbContext.GetCommentsByTarget(target)
	if err != nil {
		return nil, err
	}

	for _, commentDb := range commentsDb {
		comment := domains.Comment(commentDb)
		comments = append(comments, comment)
	}

	return comments, nil
}

func (repository *commentRepository) StoreComment(target domains.Target, userId domains.Author, message string) (*domains.Comment, error) {
	commentModel, err := repository.commentDbContext.StoreComment(target, userId, message)
	if err != nil {
		return nil, err
	}
	return (*domains.Comment)(commentModel), nil
}

func (repository *commentRepository) UpdateCommentMessage(commentId int, message string) (*domains.Comment, error) {
	commentModel, err := repository.commentDbContext.GetComment(commentId)
	if err != nil {
		return nil, err
	}
	if commentModel == nil {
		return nil, errors.ErrCommentNotFound
	}

	commentModel, err = repository.commentDbContext.UpdateCommentMessage(commentModel, message)
	if err != nil {
		return nil, err
	}

	return (*domains.Comment)(commentModel), nil
}

func (repository *commentRepository) DeleteComment(commentId int) error {
	commentModel, err := repository.commentDbContext.GetComment(commentId)
	if err != nil {
		return err
	}
	if commentModel == nil {
		return errors.ErrCommentNotFound
	}

	return repository.commentDbContext.DeleteComment(commentModel)
}

func (repository *commentRepository) IsUserOwnsComment(userId domains.Author, commentId int) (bool, error) {
	commentModel, err := repository.commentDbContext.GetComment(commentId)
	if err != nil {
		return false, err
	}
	return commentModel.AuthorID == userId, nil
}
