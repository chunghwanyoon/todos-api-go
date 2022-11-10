package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Response(c echo.Context, statusCode int, data interface{}) error {
	return c.JSON(statusCode, data)
}

func ResponseMessage(c echo.Context, msgData ResponseData) error {
	return Response(c, msgData.StatusCode, msgData.ResponseMessage)
}

func ResponseError(c echo.Context, err error) error {
	// TODO: handle errors status codes
	// need customized errors
	return Response(c, http.StatusBadRequest, err)
}
