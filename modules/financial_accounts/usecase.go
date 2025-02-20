package financialaccounts

import (
	"context"
	"difaal21/ihsan-solusi-assessment/constants"
	"difaal21/ihsan-solusi-assessment/dto"
	"difaal21/ihsan-solusi-assessment/exceptions"
	"difaal21/ihsan-solusi-assessment/messages"
	"difaal21/ihsan-solusi-assessment/repositories"
	"difaal21/ihsan-solusi-assessment/responses"

	"github.com/sirupsen/logrus"
)

type Usecase interface {
	Credit(ctx context.Context, param *dto.Credit) responses.Responses
	Debit(ctx context.Context, param *dto.Debit) responses.Responses
}

type UsecaseImpl struct {
	logger     *logrus.Logger
	repository repositories.FinancialAccountRepository
}

func NewUseCase(logger *logrus.Logger, repo repositories.FinancialAccountRepository) Usecase {
	return &UsecaseImpl{
		logger:     logger,
		repository: repo,
	}
}

func (u *UsecaseImpl) Credit(ctx context.Context, param *dto.Credit) responses.Responses {

	err := u.repository.Credit(ctx, param)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return httpResponse.BadRequest("").SetData(nil).SetMessage(messages.Common["not_found"]).Send()
		}

		u.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
		return httpResponse.InternalServerError("").SetData(exceptions.LogError{ID: ctx.Value(constants.LogContextKey)}).SetMessage(messages.Common["internal_server_error"]).Send()
	}

	return httpResponse.Ok("").SetData(nil).SetMessage("").Send()
}

func (u *UsecaseImpl) Debit(ctx context.Context, param *dto.Debit) responses.Responses {

	err := u.repository.Debit(ctx, param)
	if err != nil {
		if err == exceptions.ErrNotFound {
			return httpResponse.BadRequest("").SetData(nil).SetMessage(messages.Common["not_found"]).Send()
		}

		if err == exceptions.ErrInsufficientBalance {
			return httpResponse.BadRequest("").SetData(nil).SetMessage(messages.Users["insuficient_balance"]).Send()
		}

		u.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
		return httpResponse.InternalServerError("").SetData(exceptions.LogError{ID: ctx.Value(constants.LogContextKey)}).SetMessage(messages.Common["internal_server_error"]).Send()
	}

	return httpResponse.Ok("").SetData(nil).SetMessage("").Send()
}
