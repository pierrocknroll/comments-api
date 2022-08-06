package repositories

import (
	"comments-api/internal/core/domains"
	"comments-api/internal/repositories/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type DatabaseContext struct {
	Name          string
	SQLConnection *sqlx.DB
}

func NewDatabaseContext(name string, dsn string) (*DatabaseContext, error) {

	driverName := "postsgres"

	SQLConnection, err := sqlx.Connect(driverName, dsn)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "database": name, "dsn": dsn}).
			Error("can't create SQL database connection")
		return nil, err
	}

	dbContext := &DatabaseContext{
		Name:          name,
		SQLConnection: SQLConnection,
	}

	return dbContext, nil
}

func (dbCtxt *DatabaseContext) GetCommentsByTarget(target domains.Target) ([]models.Comment, error) {
	comments := make([]models.Comment, 0)

	query := `
		SELECT id, author_id, target, message, created_at, updated_at
		FROM comments
		WHERE target = :target
	LIMIT 1`

	rows, err := dbCtxt.SQLConnection.NamedQuery(query, map[string]interface{}{
		"target": target,
	})
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "target": target, "query": query}).
			Error("SQL query error")
		return nil, err
	}

	for rows.Next() {
		var comment models.Comment
		err = rows.StructScan(&comment)
		if err != nil {
			log.
				WithFields(log.Fields{"error": err, "target": target, "comment": comment}).
				Error("SQL query error")
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (dbCtxt *DatabaseContext) StoreComment(target domains.Target, userId domains.Author, message string) (*models.Comment, error) {
	var comment models.Comment

	query := `
		INSERT INTO comments(author_id, target, message, created_at, updated_at)
		VALUES (:author_id, :target, :message, NOW(), NOW())
		RETURNING id, author_id, target, message, created_at, updated_at
	`

	row, err := dbCtxt.SQLConnection.NamedQuery(query, map[string]interface{}{
		"author_id": userId,
		"target":    target,
		"message":   message,
	})
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "target": target, "author_id": userId, "query": query}).
			Error("SQL query error")
		return nil, err
	}
	if row.Next() {
		row.Scan(&comment)
	}

	return &comment, nil
}

func (dbCtxt *DatabaseContext) UpdateCommentMessage(comment *models.Comment, message string) (*models.Comment, error) {
	var updatedComment *models.Comment

	query := `
		UPDATE comments
		SET message = :message,
		    updated_at = NOW()
		WHERE id = :comment_id
		RETURNING id, author_id, target, message, created_at, updated_at
	`

	row, err := dbCtxt.SQLConnection.NamedQuery(query, map[string]interface{}{
		"message":    message,
		"comment_id": comment.Id,
	})
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "commentId": comment.Id, "query": query}).
			Error("SQL query error")
		return nil, err
	}
	if row.Next() {
		row.Scan(&updatedComment)
	}

	return updatedComment, nil
}

func (dbCtxt *DatabaseContext) DeleteComment(comment *models.Comment) error {
	query := `
		DELETE FROM comments
		WHERE id = :comment_id
	`

	_, err := dbCtxt.SQLConnection.NamedExec(query, map[string]interface{}{
		"comment_id": comment.Id,
	})
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "commentId": comment.Id, "query": query}).
			Error("SQL query error")
		return err
	}

	return nil
}

func (dbCtxt *DatabaseContext) GetComment(commentId int) (*models.Comment, error) {
	var comment models.Comment

	query := `
		SELECT id, author_id, target, message, created_at, updated_at
		FROM comments
		WHERE id = ?
		LIMIT 1
	`

	err := dbCtxt.SQLConnection.Get(&comment, query, commentId)
	if err != nil {
		log.
			WithFields(log.Fields{"error": err, "commentId": commentId, "query": query}).
			Error("SQL query error")
		return nil, err
	}

	return &comment, nil
}

func (dbCtxt *DatabaseContext) Close() error {

	err := dbCtxt.SQLConnection.Close()
	if err != nil {
		log.
			WithFields(log.Fields{"error": err}).
			Debug("Can't close database connection")
		return err
	}

	return nil
}
