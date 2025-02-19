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
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

type FinancialAccountRepository interface {
	GetOneByUniqueField(ctx context.Context, field string, value any) (user *entities.FinancialAccount, err error)
	Credit(ctx context.Context, bankAccountNumber string, totalAmount float64) (err error)
	// Debit(ctx context.Context, tx *sql.Tx, bankAccountNumber string, amount int64) (err error)

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

func (r *FinancialAccountRepositoryImpl) Credit(ctx context.Context, bankAccountNumber string, totalAmount float64) (err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	updatedField := map[string]any{"balance": totalAmount}
	err = r.Update(ctx, tx, "bank_account_number", bankAccountNumber, updatedField)
	if err != nil {
		return
	}

	return nil

}

func (r *FinancialAccountRepositoryImpl) GetOneByUniqueField(ctx context.Context, field string, value any) (financialAccount *entities.FinancialAccount, err error) {

	q := buildFinancialAccountQuery()

	if field != "" && value != "" {
		q.Apply(sm.Where(psql.Raw(field).EQ(psql.Arg(value))))
	}

	query, args, err := q.Build(ctx)
	if err != nil {
		r.logger.WithFields(logrus.Fields{"log": ctx.Value(constants.LogContextKey), "query": query}).Error(err)
		return
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	financialAccount, err = scanFinancialAccount(row)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return
		}
		r.logger.WithFields(logrus.Fields{"log": ctx.Value(constants.LogContextKey), "query": query}).Error(err)
		return
	}
	return

}

func (r *FinancialAccountRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, column_name string, uniqueField any, updateFields map[string]any) (err error) {
	var (
		placeholders []string
		values       []interface{}
		paramCount   = 1
	)

	for field, value := range updateFields {
		placeholders = append(placeholders, fmt.Sprintf("%s = $%d", field, paramCount))
		values = append(values, value)
		paramCount++
	}

	placeholdersStr := strings.Join(placeholders, ", ")
	command := fmt.Sprintf("UPDATE %s SET %s WHERE %s = $%d", r.tableName, placeholdersStr, column_name, paramCount)
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
