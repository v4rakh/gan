package presenter

import (
	"time"
)

type Subscription struct {
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	State     string    `json:"state"`
	Token     string    `json:"token"`
}
