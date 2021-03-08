package transaction

import "golang-crowdfunding-backend/user"

type GetTransactionDetailInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
