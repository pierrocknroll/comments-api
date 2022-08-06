package models

import (
	"comments-api/internal/core/domains"
	"time"
)

// See domains/comment for the Domain

type Comment struct {
	Id        int            `db:"id"`
	AuthorID  domains.Author `db:"author_id"`
	Target    domains.Target `db:"target"`
	Message   string         `db:"message"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`

	// ParentId  int       `db:"parent_id"`
	// Children  []*Comment  `db:"children"`
}
