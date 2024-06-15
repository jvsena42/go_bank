package dto

type CreateAccountParameters struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type UpdateAccountParameters struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}
