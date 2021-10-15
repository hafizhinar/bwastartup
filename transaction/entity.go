package transaction

import (
	"bwastartup/user"
	"time"
)

type Transactions struct {
	ID         uint64
	UserID     uint64
	CampaignID uint64
	Code       string
	Amount     float64
	Status     int
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
