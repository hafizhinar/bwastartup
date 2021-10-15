package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionsByCampaignId(input GetCampaignTransactionsInput) ([]Transactions, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignId(input GetCampaignTransactionsInput) ([]Transactions, error) {

	campaign, err := s.campaignRepository.FindId(input.ID)

	if err != nil {
		return []Transactions{}, err
	}

	if campaign.ID != input.User.ID {
		return []Transactions{}, errors.New("not an owner of the campaign")
	}

	transactions, err := s.repository.GetByCampaignId(input.ID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
