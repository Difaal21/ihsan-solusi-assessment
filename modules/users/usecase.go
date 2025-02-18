package users

import (
	"context"
	"difaal21/ihsan-solusi-assessment/dto"
	"difaal21/ihsan-solusi-assessment/repositories"
	"difaal21/ihsan-solusi-assessment/responses"

	"github.com/sirupsen/logrus"
)

type Usecase interface {
	Registration(ctx context.Context, param *dto.UserRegistration) responses.Responses
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

func (u *UsecaseImpl) Registration(ctx context.Context, param *dto.UserRegistration) responses.Responses {

	return httpResponse.Ok("").SetData(param).SetMessage("").Send()
}
