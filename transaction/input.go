package transaction

type GetTransactionDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
