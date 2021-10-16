package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

type Service interface {
	GetTransactionsByCampaignId(input GetCampaignTransactionsInput) ([]Transactions, error)
	GetTransactionsByUserId(userId uint64) ([]Transactions, error)
	CreateTransactions(input CreateTransactionsInput) (Transactions, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
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

func (s *service) GetTransactionsByUserId(userId uint64) ([]Transactions, error) {
	transactions, err := s.repository.GetUserId(userId)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransactions(input CreateTransactionsInput) (Transactions, error) {
	transaction := Transactions{}

	transaction.UserID = input.User.ID
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.Status = 1

	newTransaction, err := s.repository.Save(transaction)

	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: int64(newTransaction.Amount),
	}

	paymentUrl, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)

	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentUrl

	newTransaction, err = s.repository.Update(newTransaction)

	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
