package dto

type CreateEntryParamets struct {
	AccountId int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}
