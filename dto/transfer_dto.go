package dto

type CreateTransferParameters struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	//must be positive
	Amount int64 `json:"amount"`
}

func (parameters CreateTransferParameters) IsValid() bool {
	return parameters.Amount > 0
}
