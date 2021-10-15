package transaction

import "time"

type CampaignTransactionsFormatter struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
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
