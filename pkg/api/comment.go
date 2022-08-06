package api

import (
	"comments-api/internal/core/domains"
	"time"
)

type Comment struct {
	Id        int            `json:"id"`
	AuthorID  domains.Author `json:"-"`
	Target    domains.Target `json:"-"`
	Message   string         `json:"message"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Comments struct {
	Comments []Comment `json:"comments"`
}
