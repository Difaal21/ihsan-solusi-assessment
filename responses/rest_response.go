package responses

import (
	"github.com/labstack/echo/v4"
)

func REST(c echo.Context, r Responses) error {
	response := r.Send()

	return c.JSON(response.Code, response)
}
