package entities

import "time"

type FinancialAccount struct {
	ID                int64     `json:"id"`
	UserID            int64     `json:"user_id"`
	Balance           float64   `json:"balance"`
	BankAccountNumber string    `json:"bank_account_number"`
	CreatedAt         time.Time `json:"created_at"`
}
