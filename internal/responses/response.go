package responses

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   data,
	})
}

func BadRequestResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, Response{
		Status: http.StatusBadRequest,
		Data:   err.Error(),
	})
}

func InternalServerErrorResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, Response{
		Status: http.StatusInternalServerError,
		Data:   err.Error(),
	})
}

func NotFoundResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusNotFound, Response{
		Status: http.StatusNotFound,
		Data:   err.Error(),
	})
}

func NoContentResponse(c echo.Context) error {
	return c.JSON(http.StatusNoContent, Response{
		Status: http.StatusNoContent,
		Data:   nil,
	})
}
