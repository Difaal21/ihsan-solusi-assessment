package users

import (
	"context"
	"difaal21/ihsan-solusi-assessment/dto"
	"difaal21/ihsan-solusi-assessment/repositories"
	"difaal21/ihsan-solusi-assessment/responses"

	"github.com/sirupsen/logrus"
)

type Usecase interface {
	Login(ctx context.Context, param *dto.UserLoginRequest) responses.Responses
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

func (u *UsecaseImpl) Login(ctx context.Context, param *dto.UserLoginRequest) responses.Responses {

	return httpResponse.Ok("").SetData(nil).SetMessage("").Send()
}

// func (u *UsecaseImpl) Login(ctx context.Context, param *dto.UserLoginRequest) responses.Responses {
// 	user, err := u.repository.GetOneUserByUniqueField(ctx, "ua.email", param.Email)
// 	if err != nil {
// 		if err == exceptions.ErrNotFound {
// 			return httpResponse.BadRequest("INVALID_CREDENTIAL").SetData(nil).SetMessage(messages.Users["invalid_credential"]).Send()
// 		}

// 		u.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
// 		return httpResponse.InternalServerError("").SetData(exceptions.LogError{ID: ctx.Value(constants.LogContextKey)}).SetMessage(messages.Common["internal_server_error"]).Send()
// 	}

// 	if !user.IsEmailVerified {
// 		return httpResponse.BadRequest("EMAIL_NOT_VERIFIED").SetData(nil).SetMessage(messages.Users["email_not_verified"]).Send()
// 	}

// 	passwordMatch := u.bcrypt.Verify(user.Password, []byte(param.Password))
// 	if !passwordMatch {
// 		return httpResponse.BadRequest("INVALID_CREDENTIAL").SetData(nil).SetMessage(messages.Users["invalid_credential"]).Send()
// 	}

// 	loginSessionValue := entities.UserLoginSessionCache{
// 		ID:        user.ID,
// 		Name:      user.Name,
// 		Email:     user.Email,
// 		CreatedAt: user.CreatedAt,
// 	}

// 	loginSessionBuff, err := json.Marshal(loginSessionValue)
// 	if err != nil {
// 		u.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
// 		return httpResponse.InternalServerError("").SetData(exceptions.LogError{ID: ctx.Value(constants.LogContextKey)}).SetMessage(messages.Common["marshaling_data"]).Send()
// 	}

// 	loginSession, _ := uuid.NewV7()
// 	loginSessionTTL := time.Hour
// 	key := fmt.Sprintf("%s:%s", constants.UserLoginSessionKey, loginSession)
// 	err = u.cache.Set(ctx, key, loginSessionBuff, loginSessionTTL)
// 	if err != nil {
// 		u.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(err)
// 		return httpResponse.InternalServerError("").SetData(exceptions.LogError{ID: ctx.Value(constants.LogContextKey)}).SetMessage(messages.Common["set_cache"]).Send()
// 	}

// 	response := dto.NewUserLoginResponse(loginSession, time.Duration(loginSessionTTL.Seconds()))
// 	return httpResponse.Ok("").SetData(response).SetMessage(messages.Users["login_success"]).Send()
// }
