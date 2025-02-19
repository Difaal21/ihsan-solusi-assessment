package users

import (
	"context"
	"difaal21/ihsan-solusi-assessment/constants"
	"difaal21/ihsan-solusi-assessment/dto"
	"difaal21/ihsan-solusi-assessment/entities"
	"difaal21/ihsan-solusi-assessment/exceptions"
	"difaal21/ihsan-solusi-assessment/helpers/date"
	"difaal21/ihsan-solusi-assessment/messages"
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

	uniqueNationalityID, err := u.repository.GetOneUserByUniqueField(ctx, "u.nationality_id", param.NationalityID)
	if err != nil {
		if err == exceptions.ErrInternalServerError {
			u.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
			return httpResponse.InternalServerError("").SetData(exceptions.LogError{ID: ctx.Value(constants.LogContextKey)}).SetMessage(messages.Common["internal_server_error"]).Send()
		}
	}

	if uniqueNationalityID != nil {
		return httpResponse.BadRequest("").SetData(nil).SetMessage(messages.Users["user_already_exist"]).Send()
	}

	uniquePhoneNumber, err := u.repository.GetOneUserByUniqueField(ctx, "u.phone_number", param.PhoneNumber)
	if err != nil {
		if err == exceptions.ErrInternalServerError {
			u.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
			return httpResponse.InternalServerError("").SetData(exceptions.LogError{ID: ctx.Value(constants.LogContextKey)}).SetMessage(messages.Common["internal_server_error"]).Send()
		}
	}

	if uniquePhoneNumber != nil {
		return httpResponse.BadRequest("").SetData(nil).SetMessage(messages.Users["user_already_exist"]).Send()
	}

	user := &entities.Users{
		Name:          param.Name,
		NationalityID: param.NationalityID,
		PhoneNumber:   param.PhoneNumber,
		CreatedAt:     *date.CurrentUTCTime(),
		Balance:       0,
	}

	err = u.repository.Registration(ctx, user)
	if err != nil {
		u.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
		return httpResponse.InternalServerError("").SetData(exceptions.LogError{ID: ctx.Value(constants.LogContextKey)}).SetMessage(messages.Common["internal_server_error"]).Send()
	}

	return httpResponse.Ok("").SetData(nil).SetMessage("").Send()
}
