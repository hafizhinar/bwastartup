package transaction

import "bwastartup/user"

type GetCampaignTransactionsInput struct {
	ID   uint64 `uri:"id" binding:"required"`
	User user.User
}
