package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entery struct {
	ID        int64         `json:"id"`
	AccountID sql.NullInt64 `json:"account_id"`
	// can be negative or +ve
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID            int64         `json:"id"`
	FromAccountID sql.NullInt64 `json:"from_account_id"`
	ToAccountID   sql.NullInt64 `json:"to_account_id"`
	// should be +ve
	Ammount   int64     `json:"ammount"`
	CreatedAt time.Time `json:"created_at"`
}
