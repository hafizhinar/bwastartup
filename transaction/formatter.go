package transaction

import "time"

type CampaignTransactionsFormatter struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionsFormatter struct {
	ID        uint64            `json:"id"`
	Amount    float64           `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatCampaignTransaction(transaction Transactions) CampaignTransactionsFormatter {
	formatter := CampaignTransactionsFormatter{}

	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func FormatCampaignTransactions(transactions []Transactions) []CampaignTransactionsFormatter {

	if len(transactions) == 0 {
		return []CampaignTransactionsFormatter{}
	}

	var transactionsFormatter []CampaignTransactionsFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)

		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

func FormatUserTransaction(transaction Transactions) UserTransactionsFormatter {
	var status string
	formatter := UserTransactionsFormatter{}

	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount

	switch transaction.Status {
	case 1:
		status = "Pending"
	case 2:
		status = "Paid"
	}

	formatter.Status = status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {

		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].Filename
	}

	formatter.Campaign = campaignFormatter

	return formatter

}

func FormatUserTransactions(transactions []Transactions) []UserTransactionsFormatter {

	if len(transactions) == 0 {
		return []UserTransactionsFormatter{}
	}

	var transactionsFormatter []UserTransactionsFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)

		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}
