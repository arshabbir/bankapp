package domain

import "time"

type Account struct {
	Id       int64     `json:"id"`
	Owner    string    `json:"owner_name"`
	Balance  int64     `json:"balance"`
	Currency string    `json:"currency"`
	Created  time.Time `json:"created_at"`
}

type Entries struct {
	AccountID int64 `json:"accountid"`
	Amount    int64 `json:"amount"`
}

type Transactions struct {
	FromAccount int64 `json:"from_account"`
	ToAccount   int64 `json:"to_account"`
	Amount      int64 `json:"amount"`
}
