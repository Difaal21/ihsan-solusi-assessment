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
	"github.com/stephenafamo/bob/dialect/mysql"
	"github.com/stephenafamo/bob/dialect/mysql/sm"
)

type UserRepository interface {
	GetOneUserByUniqueField(ctx context.Context, field string, value any) (user entities.Users, err error)
	Update(ctx context.Context, tx *sql.Tx, tableName string, id int64, updateFields map[string]any) (err error)
	EnableMFATOTP(ctx context.Context, userId int64, data map[string]any) (err error)
}

type UserRepositoryImpl struct {
	logger *logrus.Logger
	db     *sql.DB
}

func NewUserRepository(logger *logrus.Logger, db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		logger: logger,
		db:     db,
	}
}

func (r *UserRepositoryImpl) EnableMFATOTP(ctx context.Context, userId int64, data map[string]any) (err error) {

	tx, err := r.db.Begin()
	if err != nil {
		r.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
		return exceptions.ErrInternalServerError
	}

	err = r.Update(ctx, tx, "users_account", userId, data)
	if err != nil {
		tx.Rollback()
		r.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
		return exceptions.ErrInternalServerError
	}

	err = tx.Commit()
	if err != nil {
		r.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
		return exceptions.ErrInternalServerError
	}

	return
}

func (r *UserRepositoryImpl) GetOneUserByUniqueField(ctx context.Context, field string, value any) (user entities.Users, err error) {

	q := buildUserQuery()

	if field != "" && value != "" {
		q.Apply(sm.Where(mysql.Raw(field).EQ(mysql.Arg(value))))
	}

	query, args, err := q.Build(ctx)
	if err != nil {
		r.logger.WithFields(logrus.Fields{"log": ctx.Value(constants.LogContextKey), "query": query}).Error(err)
		return
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	user, err = scanUser(row)
	if err != nil {
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
