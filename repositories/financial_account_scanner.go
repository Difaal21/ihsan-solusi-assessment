package repositories

import (
	"database/sql"
	"difaal21/ihsan-solusi-assessment/database/postgresql"
	"difaal21/ihsan-solusi-assessment/entities"
)

func scanFinancialAccount(row *sql.Row) (*entities.FinancialAccount, error) {
	var financialAccount = &entities.FinancialAccount{}

	err := row.Scan(&financialAccount.ID, &financialAccount.UserID, &financialAccount.Balance, &financialAccount.BankAccountNumber, &financialAccount.CreatedAt)
	err = postgresql.WrapError(err)
	if err != nil {
		return nil, err
	}
	return financialAccount, nil
}
