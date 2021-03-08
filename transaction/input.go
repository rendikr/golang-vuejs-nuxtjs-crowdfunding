package transaction

import "golang-crowdfunding-backend/user"

type GetTransactionDetailInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	CampaignID int `json:"campaign_id"`
	Amount     int `json:"amount" binding:"required"`
	User       user.User
}
