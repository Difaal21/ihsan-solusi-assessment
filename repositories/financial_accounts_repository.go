package repositories

import (
	"context"
	"database/sql"
	"difaal21/ihsan-solusi-assessment/constants"
	"difaal21/ihsan-solusi-assessment/entities"
	"difaal21/ihsan-solusi-assessment/exceptions"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type FinancialAccountRepository interface {
	Update(ctx context.Context, tx *sql.Tx, column_name string, uniqueField any, updateFields map[string]any) (err error)
	Insert(ctx context.Context, tx *sql.Tx, financialAccount *entities.FinancialAccount) (id int64, err error)
}

type FinancialAccountRepositoryImpl struct {
	logger    *logrus.Logger
	db        *sql.DB
	tableName string
}

func NewFinancialAccountRepository(logger *logrus.Logger, db *sql.DB) FinancialAccountRepository {
	return &FinancialAccountRepositoryImpl{
		logger:    logger,
		db:        db,
		tableName: "financial_accounts",
	}
}

func (r *FinancialAccountRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, column_name string, uniqueField any, updateFields map[string]any) (err error) {
	var (
		placeholders []string
		values       []interface{}
	)

	for field, value := range updateFields {
		placeholders = append(placeholders, field+" = ?")
		values = append(values, value)
	}

	placeholdersStr := strings.Join(placeholders, ", ")
	command := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", r.tableName, placeholdersStr, column_name)
	values = append(values, uniqueField)

	result, err := tx.ExecContext(ctx, command, values...)
	if err != nil {
		r.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
		return exceptions.ErrInternalServerError
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
		return exceptions.ErrInternalServerError
	}

	if rowAffected == 0 {
		return exceptions.ErrNotFound
	}

	return nil
}

func (r *FinancialAccountRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, financialAccount *entities.FinancialAccount) (id int64, err error) {

	command := fmt.Sprintf(`
    INSERT INTO %s (
        user_id,
        balance,
        bank_account_number,
        created_at
    ) VALUES ($1, $2, $3, $4)
		RETURNING id
    `, r.tableName)

	result := tx.QueryRowContext(ctx, command, financialAccount.UserID, financialAccount.Balance, financialAccount.BankAccountNumber, financialAccount.CreatedAt)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"log":   ctx.Value(constants.LogContextKey),
			"query": command,
		}).Error(err)
		return
	}

	err = result.Scan(&financialAccount.ID)
	if err != nil {
		return
	}

	return financialAccount.ID, nil
}
