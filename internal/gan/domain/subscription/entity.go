package subscription

import "time"

type Subscription struct {
	Address   string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"time;autoCreateTime"`
	UpdatedAt time.Time `gorm:"time;autoUpdateTime"`
	State     string
	Token     string
}
