package financialaccounts

import (
	"context"
	"difaal21/ihsan-solusi-assessment/constants"
	"difaal21/ihsan-solusi-assessment/dto"
	"difaal21/ihsan-solusi-assessment/helpers/validation"
	"difaal21/ihsan-solusi-assessment/messages"
	"difaal21/ihsan-solusi-assessment/responses"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var httpResponse = responses.NewResponse()

type HTTPHandler struct {
	logger   *logrus.Logger
	validate *validator.Validate
	usecase  Usecase
}

func NewHTTPHandler(router *echo.Echo, logger *logrus.Logger, validate *validator.Validate, usecase Usecase) {
	handler := &HTTPHandler{
		logger:   logger,
		validate: validate,
		usecase:  usecase,
	}

	router.POST("/ihsan-solusi-assessment/v1/tabung", handler.Credit)
	router.POST("/ihsan-solusi-assessment/v1/tarik", handler.Debit)
	// router.GET("/ihsan-solusi-assessment/v1/saldo", handler.CheckBalance)
}

func (handler *HTTPHandler) Credit(c echo.Context) error {
	var ctx = c.Request().Context()
	var payload *dto.Credit

	logId, _ := uuid.NewV7()
	ctx = context.WithValue(ctx, constants.LogContextKey, logId)

	defer func() {
		r := recover()
		if r != nil {
			handler.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(r)
			httpResponse.InternalServerError("").SetData(r).SetMessage(messages.Common["internal_server_error"]).Send()
			responses.REST(c, httpResponse)
			return
		}
	}()

	if err := c.Bind(&payload); err != nil {
		httpResponse.UnprocessableEntity("").SetData(nil).SetMessage(messages.Common["unprocessible_entity"]).Send()
		return responses.REST(c, httpResponse)
	}

	if err := validation.RequestBody(handler.validate, payload); err != nil {
		httpResponse.BadRequest("").SetData(err).SetMessage(messages.Common["invalid_request"]).Send()
		return responses.REST(c, httpResponse)
	}

	return responses.REST(c, handler.usecase.Credit(ctx, payload))
}

func (handler *HTTPHandler) Debit(c echo.Context) error {
	var ctx = c.Request().Context()
	var payload *dto.Credit

	logId, _ := uuid.NewV7()
	ctx = context.WithValue(ctx, constants.LogContextKey, logId)

	defer func() {
		r := recover()
		if r != nil {
			handler.logger.WithField("log", ctx.Value(constants.LogContextKey)).Error(r)
			httpResponse.InternalServerError("").SetData(r).SetMessage(messages.Common["internal_server_error"]).Send()
			responses.REST(c, httpResponse)
			return
		}
	}()

	if err := c.Bind(&payload); err != nil {
		httpResponse.UnprocessableEntity("").SetData(nil).SetMessage(messages.Common["unprocessible_entity"]).Send()
		return responses.REST(c, httpResponse)
	}

	if err := validation.RequestBody(handler.validate, payload); err != nil {
		httpResponse.BadRequest("").SetData(err).SetMessage(messages.Common["invalid_request"]).Send()
		return responses.REST(c, httpResponse)
	}

	return responses.REST(c, handler.usecase.Debit(ctx, payload))
}
