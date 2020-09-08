package announcement

import "github.com/v4rakh/gan/internal/gan/domain"

type Announcement struct {
	domain.Base
	Title   string
	Content string
}
