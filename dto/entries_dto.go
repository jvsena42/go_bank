package dto

type CreateEntryParamets struct {
	AccountId string `json:"account_id"`
	Amount    int64  `json:"amount"`
}
