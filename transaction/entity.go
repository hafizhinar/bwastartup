package transaction

import (
	"bwastartup/campaign"
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
	PaymentURL string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
