package campaign

import (
	"bwastartup/user"
	"time"
)

type Campaign struct {
	ID               uint64
	UserID           uint64
	Name             string
	ShortDescription string
	Description      string
	TargetAmount     string
	CurrentAmount    string
	Perks            string
	BackerCount      uint64
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImages
	User             user.User
}

type CampaignImages struct {
	ID         uint64
	CampaignID uint64
	Filename   string
	IsPrimary  int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
