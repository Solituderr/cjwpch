package Response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SenRes(c echo.Context, code int, msg string, data ...interface{}) error {
	return c.JSON(http.StatusOK, Res{
		code,
		msg,
		data,
	})
}
