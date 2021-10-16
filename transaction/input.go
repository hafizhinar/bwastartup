package transaction

import "bwastartup/user"

type GetCampaignTransactionsInput struct {
	ID   uint64 `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionsInput struct {
	Amount     float64 `json:"amount"`
	CampaignID uint64  `json:"campaign_id"`
	User       user.User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
