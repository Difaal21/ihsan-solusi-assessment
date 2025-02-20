package repositories

import (
	"context"
	"database/sql"
	"difaal21/ihsan-solusi-assessment/constants"
	"difaal21/ihsan-solusi-assessment/entities"
	"difaal21/ihsan-solusi-assessment/exceptions"
	"fmt"
	"math/rand"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

type UserRepository interface {
	GetOneUserByUniqueField(ctx context.Context, field string, value any) (user *entities.Users, err error)
	Registration(ctx context.Context, users *entities.Users) (financialAccount *entities.FinancialAccount, err error)

	Update(ctx context.Context, tx *sql.Tx, tableName string, id int64, updateFields map[string]any) (err error)
	Insert(ctx context.Context, tx *sql.Tx, users *entities.Users) (id int64, err error)
}

type UserRepositoryImpl struct {
	logger                     *logrus.Logger
	db                         *sql.DB
	tableName                  string
	financialAccountRepository FinancialAccountRepository
}

func NewUserRepository(logger *logrus.Logger, db *sql.DB, financialAccountRepo FinancialAccountRepository) UserRepository {
	return &UserRepositoryImpl{
		logger:                     logger,
		db:                         db,
		tableName:                  "users",
		financialAccountRepository: financialAccountRepo,
	}
}

func generateBankAccountNumber() string {
	accountNumber := rand.Int63n(1e12) // Generate a random 12-digit number
	return fmt.Sprintf("%012d", accountNumber)
}

func (r *UserRepositoryImpl) Registration(ctx context.Context, user *entities.Users) (financialAccount *entities.FinancialAccount, err error) {
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

		err = tx.Commit()
		if err != nil {
			r.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
			return
		}
	}()

	userId, err := r.Insert(ctx, tx, user)
	if err != nil {
		r.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
		return
	}

	financialAccount = &entities.FinancialAccount{
		UserID:            userId,
		Balance:           0,
		BankAccountNumber: generateBankAccountNumber(),
		CreatedAt:         user.CreatedAt,
	}

	_, err = r.financialAccountRepository.Insert(ctx, tx, financialAccount)
	if err != nil {
		r.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
		return
	}

	return
}

func (r *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user *entities.Users) (id int64, err error) {

	command := fmt.Sprintf(`
    INSERT INTO %s (
        name,
        phone_number,
        nationality_id,
        created_at
    ) VALUES ($1, $2, $3, $4)
		RETURNING id
    `, r.tableName)

	result := tx.QueryRowContext(ctx, command, user.Name, user.PhoneNumber, user.NationalityID, user.CreatedAt)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"log":   ctx.Value(constants.LogContextKey),
			"query": command,
		}).Error(err)
		return
	}

	err = result.Scan(&user.ID)
	if err != nil {
		return
	}

	return user.ID, nil
}

func (r *UserRepositoryImpl) GetOneUserByUniqueField(ctx context.Context, field string, value any) (user *entities.Users, err error) {

	q := buildUserQuery()

	if field != "" && value != "" {
		q.Apply(sm.Where(psql.Raw(field).EQ(psql.Arg(value))))
	}

	query, args, err := q.Build(ctx)
	if err != nil {
		r.logger.WithFields(logrus.Fields{"log": ctx.Value(constants.LogContextKey), "query": query}).Error(err)
		return
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	user, err = scanUser(row)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return
		}
		r.logger.WithFields(logrus.Fields{"log": ctx.Value(constants.LogContextKey), "query": query}).Error(err)
		return
	}
	return
}

func (r *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, tableName string, id int64, updateFields map[string]any) (err error) {
	var (
		placeholders []string
		values       []interface{}
	)

	for field, value := range updateFields {
		placeholders = append(placeholders, field+" = ?")
		values = append(values, value)
	}

	placeholdersStr := strings.Join(placeholders, ", ")
	command := fmt.Sprintf("UPDATE %s SET %s WHERE user_id = ?", tableName, placeholdersStr)
	values = append(values, id)

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
