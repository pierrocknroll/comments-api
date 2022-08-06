package domains

import (
	"time"
)

type Comment struct {
	Id        int
	AuthorID  Author
	Target    Target
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Target string
type Author int
