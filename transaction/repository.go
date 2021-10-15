package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByCampaignId(campaignId uint64) ([]Transactions, error)
	GetUserId(userId uint64) ([]Transactions, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignId(campaignId uint64) ([]Transactions, error) {
	var transactions []Transactions

	err := r.db.Preload("User").Where("campaign_id = ?", campaignId).Order("id desc").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetUserId(userId uint64) ([]Transactions, error) {
	var transactions []Transactions

	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id", userId).Order("id desc").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
