package presenter

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Announcement struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}
