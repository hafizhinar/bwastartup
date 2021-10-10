package campaign

type GetCampaignDetailInput struct {
	ID uint64 `uri:"id" binding:"required"`
}
