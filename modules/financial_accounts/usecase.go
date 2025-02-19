package financialaccounts

import (
	"context"
	"difaal21/ihsan-solusi-assessment/dto"
	"difaal21/ihsan-solusi-assessment/repositories"
	"difaal21/ihsan-solusi-assessment/responses"

	"github.com/sirupsen/logrus"
)

type Usecase interface {
	Credit(ctx context.Context, param *dto.Credit) responses.Responses
	Debit(ctx context.Context, param *dto.Credit) responses.Responses
}

type UsecaseImpl struct {
	logger     *logrus.Logger
	repository repositories.UserRepository
}

func NewUseCase(logger *logrus.Logger, repo repositories.UserRepository) Usecase {
	return &UsecaseImpl{
		logger:     logger,
		repository: repo,
	}
}

func (u *UsecaseImpl) Credit(ctx context.Context, param *dto.Credit) responses.Responses {

	return httpResponse.Ok("").SetData(nil).SetMessage("").Send()
}

func (u *UsecaseImpl) Debit(ctx context.Context, param *dto.Credit) responses.Responses {

	return httpResponse.Ok("").SetData(nil).SetMessage("").Send()
}
